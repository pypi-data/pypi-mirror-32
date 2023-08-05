import copy
import json
import warnings
from urllib.parse import urlparse

from django.db import DEFAULT_DB_ALIAS, models
from django.db.models.sql import Query
from django.db.models.sql.constants import MULTI
from django.utils.crypto import get_random_string
from elasticsearch import Elasticsearch, NotFoundError
from elasticsearch.helpers import bulk

from wagtail.search.backends.base import (
    BaseSearchBackend, BaseSearchQueryCompiler, BaseSearchResults)
from wagtail.search.index import FilterField, Indexed, RelatedFields, SearchField, class_is_indexed
from wagtail.search.query import (
    And, Boost, Filter, Fuzzy, MatchAll, Not, Or, PlainText, Prefix, Term)
from wagtail.utils.deprecation import RemovedInWagtail22Warning
from wagtail.utils.utils import deep_update


def get_model_root(model):
    """
    This function finds the root model for any given model. The root model is
    the highest concrete model that it descends from. If the model doesn't
    descend from another concrete model then the model is it's own root model so
    it is returned.

    Examples:
    >>> get_model_root(wagtailcore.Page)
    wagtailcore.Page

    >>> get_model_root(myapp.HomePage)
    wagtailcore.Page

    >>> get_model_root(wagtailimages.Image)
    wagtailimages.Image
    """
    if model._meta.parents:
        parent_model = list(model._meta.parents.items())[0][0]
        return get_model_root(parent_model)

    return model


class Elasticsearch2Mapping:
    all_field_name = '_all'

    # Was originally named '_partials' but renamed '_edgengrams' when we added Elasticsearch 6 support
    # The ES 2 and 5 backends still use the old name for backwards compatibility
    edgengrams_field_name = '_partials'

    type_map = {
        'AutoField': 'integer',
        'BinaryField': 'binary',
        'BooleanField': 'boolean',
        'CharField': 'string',
        'CommaSeparatedIntegerField': 'string',
        'DateField': 'date',
        'DateTimeField': 'date',
        'DecimalField': 'double',
        'FileField': 'string',
        'FilePathField': 'string',
        'FloatField': 'double',
        'IntegerField': 'integer',
        'BigIntegerField': 'long',
        'IPAddressField': 'string',
        'GenericIPAddressField': 'string',
        'NullBooleanField': 'boolean',
        'OneToOneField': 'integer',
        'PositiveIntegerField': 'integer',
        'PositiveSmallIntegerField': 'integer',
        'SlugField': 'string',
        'SmallIntegerField': 'integer',
        'TextField': 'string',
        'TimeField': 'date',
    }

    keyword_type = 'string'
    text_type = 'string'
    set_index_not_analyzed_on_filter_fields = True
    edgengram_analyzer_config = {
        'analyzer': 'edgengram_analyzer',
        'search_analyzer': 'standard',
    }

    def __init__(self, model):
        self.model = model

    def get_parent(self):
        for base in self.model.__bases__:
            if issubclass(base, Indexed) and issubclass(base, models.Model):
                return type(self)(base)

    def get_document_type(self):
        return self.model.indexed_get_content_type()

    def get_field_column_name(self, field):
        # Fields in derived models get prefixed with their model name, fields
        # in the root model don't get prefixed at all
        # This is to prevent mapping clashes in cases where two page types have
        # a field with the same name but a different type.
        root_model = get_model_root(self.model)
        definition_model = field.get_definition_model(self.model)

        if definition_model != root_model:
            prefix = definition_model._meta.app_label.lower() + '_' + definition_model.__name__.lower() + '__'
        else:
            prefix = ''

        if isinstance(field, FilterField):
            return prefix + field.get_attname(self.model) + '_filter'
        elif isinstance(field, SearchField):
            return prefix + field.get_attname(self.model)
        elif isinstance(field, RelatedFields):
            return prefix + field.field_name

    def get_content_type(self):
        """
        Returns the content type as a string for the model.

        For example: "wagtailcore.Page"
                     "myapp.MyModel"
        """
        return self.model._meta.app_label + '.' + self.model.__name__

    def get_all_content_types(self):
        """
        Returns all the content type strings that apply to this model.
        This includes the models' content type and all concrete ancestor
        models that inherit from Indexed.

        For example: ["myapp.MyPageModel", "wagtailcore.Page"]
                     ["myapp.MyModel"]
        """
        # Add our content type
        content_types = [self.get_content_type()]

        # Add all ancestor classes content types as well
        ancestor = self.get_parent()
        while ancestor:
            content_types.append(ancestor.get_content_type())
            ancestor = ancestor.get_parent()

        return content_types

    def get_field_mapping(self, field):
        if isinstance(field, RelatedFields):
            mapping = {'type': 'nested', 'properties': {}}
            nested_model = field.get_field(self.model).related_model
            nested_mapping = type(self)(nested_model)

            for sub_field in field.fields:
                sub_field_name, sub_field_mapping = nested_mapping.get_field_mapping(sub_field)
                mapping['properties'][sub_field_name] = sub_field_mapping

            return self.get_field_column_name(field), mapping
        else:
            mapping = {'type': self.type_map.get(field.get_type(self.model), 'string')}

            if isinstance(field, SearchField):
                if mapping['type'] == 'string':
                    mapping['type'] = self.text_type

                if field.boost:
                    mapping['boost'] = field.boost

                if field.partial_match:
                    mapping.update(self.edgengram_analyzer_config)

                mapping['include_in_all'] = True

            elif isinstance(field, FilterField):
                if mapping['type'] == 'string':
                    mapping['type'] = self.keyword_type

                if self.set_index_not_analyzed_on_filter_fields:
                    # Not required on ES5 as that uses the "keyword" type for
                    # filtered string fields
                    mapping['index'] = 'not_analyzed'

                mapping['include_in_all'] = False

            if 'es_extra' in field.kwargs:
                for key, value in field.kwargs['es_extra'].items():
                    mapping[key] = value

            return self.get_field_column_name(field), mapping

    def get_mapping(self):
        # Make field list
        fields = {
            'pk': dict(type=self.keyword_type, store=True, include_in_all=False),
            'content_type': dict(type=self.keyword_type, include_in_all=False),
            self.edgengrams_field_name: dict(type=self.text_type, include_in_all=False),
        }
        fields[self.edgengrams_field_name].update(self.edgengram_analyzer_config)

        if self.set_index_not_analyzed_on_filter_fields:
            # Not required on ES5 as that uses the "keyword" type for
            # filtered string fields
            fields['pk']['index'] = 'not_analyzed'
            fields['content_type']['index'] = 'not_analyzed'

        fields.update(dict(
            self.get_field_mapping(field) for field in self.model.get_search_fields()
        ))

        return {
            self.get_document_type(): {
                'properties': fields,
            }
        }

    def get_document_id(self, obj):
        return obj.indexed_get_toplevel_content_type() + ':' + str(obj.pk)

    def _get_nested_document(self, fields, obj):
        doc = {}
        partials = []
        model = type(obj)
        mapping = type(self)(model)

        for field in fields:
            value = field.get_value(obj)
            doc[mapping.get_field_column_name(field)] = value

            # Check if this field should be added into _edgengrams
            if isinstance(field, SearchField) and field.partial_match:
                partials.append(value)

        return doc, partials

    def get_document(self, obj):
        # Build document
        doc = dict(pk=str(obj.pk), content_type=self.get_all_content_types())
        partials = []
        for field in self.model.get_search_fields():
            value = field.get_value(obj)

            if isinstance(field, RelatedFields):
                if isinstance(value, models.Manager):
                    nested_docs = []

                    for nested_obj in value.all():
                        nested_doc, extra_edgengrams = self._get_nested_document(field.fields, nested_obj)
                        nested_docs.append(nested_doc)
                        partials.extend(extra_edgengrams)

                    value = nested_docs
                elif isinstance(value, models.Model):
                    value, extra_edgengrams = self._get_nested_document(field.fields, value)
                    partials.extend(extra_edgengrams)

            doc[self.get_field_column_name(field)] = value

            # Check if this field should be added into _edgengrams
            if isinstance(field, SearchField) and field.partial_match:
                partials.append(value)

        # Add partials to document
        doc[self.edgengrams_field_name] = partials

        return doc

    def __repr__(self):
        return '<ElasticsearchMapping: %s>' % (self.model.__name__, )


class Elasticsearch2SearchQueryCompiler(BaseSearchQueryCompiler):
    mapping_class = Elasticsearch2Mapping
    DEFAULT_OPERATOR = 'or'

    def __init__(self, *args, **kwargs):
        super(Elasticsearch2SearchQueryCompiler, self).__init__(*args, **kwargs)
        self.mapping = self.mapping_class(self.queryset.model)

        # Convert field names into index column names
        if self.fields:
            fields = []
            searchable_fields = {f.field_name: f for f in self.queryset.model.get_searchable_search_fields()}
            for field_name in self.fields:
                if field_name in searchable_fields:
                    field_name = self.mapping.get_field_column_name(searchable_fields[field_name])

                fields.append(field_name)

            self.remapped_fields = fields
        else:
            self.remapped_fields = None

    def _process_lookup(self, field, lookup, value):
        column_name = self.mapping.get_field_column_name(field)

        if lookup == 'exact':
            if value is None:
                return {
                    'missing': {
                        'field': column_name,
                    }
                }
            else:
                return {
                    'term': {
                        column_name: value,
                    }
                }

        if lookup == 'isnull':
            if value:
                return {
                    'missing': {
                        'field': column_name,
                    }
                }
            else:
                return {
                    'exists': {
                        'field': column_name,
                    }
                }

        if lookup in ['startswith', 'prefix']:
            return {
                'prefix': {
                    column_name: value,
                }
            }

        if lookup in ['gt', 'gte', 'lt', 'lte']:
            return {
                'range': {
                    column_name: {
                        lookup: value,
                    }
                }
            }

        if lookup == 'range':
            lower, upper = value

            return {
                'range': {
                    column_name: {
                        'gte': lower,
                        'lte': upper,
                    }
                }
            }

        if lookup == 'in':
            if isinstance(value, Query):
                db_alias = self.queryset._db or DEFAULT_DB_ALIAS
                resultset = value.get_compiler(db_alias).execute_sql(result_type=MULTI)
                value = [row[0] for chunk in resultset for row in chunk]

            elif not isinstance(value, list):
                value = list(value)
            return {
                'terms': {
                    column_name: value,
                }
            }

    def _connect_filters(self, filters, connector, negated):
        if filters:
            if len(filters) == 1:
                filter_out = filters[0]
            else:
                filter_out = {
                    connector.lower(): [
                        fil for fil in filters if fil is not None
                    ]
                }

            if negated:
                filter_out = {
                    'not': filter_out
                }

            return filter_out

    def _compile_term_query(self, query_type, value, field, boost=1.0, **extra):
        term_query = {
            'value': value,
        }

        if boost != 1.0:
            term_query['boost'] = boost

        return {
            query_type: {
                field: term_query,
            }
        }

    def _compile_plaintext_query(self, query, fields, boost=1.0):
        match_query = {
            'query': query.query_string
        }

        if query.operator != 'or':
            match_query['operator'] = query.operator

        if boost != 1.0:
            match_query['boost'] = boost

        if len(fields) == 1:
            return {
                'match': {
                    fields[0]: match_query
                }
            }
        else:
            match_query['fields'] = fields

            return {
                'multi_match': match_query
            }

    def _compile_query(self, query, field, boost=1.0):
        if isinstance(query, MatchAll):
            match_all_query = {}

            if boost != 1.0:
                match_all_query['boost'] = boost

            return {'match_all': match_all_query}

        elif isinstance(query, Term):
            return self._compile_term_query('term', query.term, field, query.boost * boost)

        elif isinstance(query, Prefix):
            return self._compile_term_query('prefix', query.prefix, field, query.boost * boost)

        elif isinstance(query, Fuzzy):
            return self._compile_term_query('fuzzy', query.term, field, query.boost * boost, fuzziness=query.max_distance)

        elif isinstance(query, And):
            return {
                'bool': {
                    'must': [
                        self._compile_query(child_query, field, boost)
                        for child_query in query.get_children()
                    ]
                }
            }

        elif isinstance(query, Or):
            return {
                'bool': {
                    'should': [
                        self._compile_query(child_query, field, boost)
                        for child_query in query.get_children()
                    ]
                }
            }

        elif isinstance(query, Not):
            return {
                'bool': {
                    'mustNot': self._compile_query(query.subquery, field, boost)
                }
            }

        elif isinstance(query, PlainText):
            return self._compile_plaintext_query(self.query, [field], boost)

        elif isinstance(query, Filter):
            bool_query = {
                'must': self._compile_query(query.query, field, boost),
            }

            if query.include:
                bool_query['filter'] = self._compile_query(query.include, field, 0.0)

            if query.exclude:
                bool_query['mustNot'] = self._compile_query(query.exclude, field, 0.0)

            return {
                'bool': bool_query,
            }

        elif isinstance(query, Boost):
            return self._compile_query(query.subquery, field, boost * query.boost)

        else:
            raise NotImplementedError(
                '`%s` is not supported by the Elasticsearch search backend.'
                % query.__class__.__name__)

    def get_inner_query(self):
        fields = self.remapped_fields or [self.mapping.all_field_name, self.mapping.edgengrams_field_name]

        if len(fields) == 0:
            # No fields. Return a query that'll match nothing
            return {
                'bool': {
                    'mustNot': {'match_all': {}}
                }
            }

        # Handle MatchAll and PlainText separately as they were supported
        # before "search query classes" was implemented and we'd like to
        # keep the query the same as before
        if isinstance(self.query, MatchAll):
            return {'match_all': {}}

        elif isinstance(self.query, PlainText):
            return self._compile_plaintext_query(self.query, fields)

        else:
            if len(fields) == 1:
                return self._compile_query(self.query, fields[0])
            else:
                # Compile a query for each field then combine with disjunction
                # max (or operator which takes the max score out of each of the
                # field queries)
                field_queries = []
                for field in fields:
                    field_queries.append(self._compile_query(self.query, field))

                return {
                    'dis_max': {
                        'queries': field_queries
                    }
                }

    def get_content_type_filter(self):
        # Query content_type using a "match" query. See comment in
        # Elasticsearch2Mapping.get_document for more details
        content_type = self.mapping_class(self.queryset.model).get_content_type()

        return {
            'match': {
                'content_type': content_type
            }
        }

    def get_filters(self):
        filters = []

        # Filter by content type
        filters.append(self.get_content_type_filter())

        # Apply filters from queryset
        queryset_filters = self._get_filters_from_queryset()
        if queryset_filters:
            filters.append(queryset_filters)

        return filters

    def get_query(self):
        inner_query = self.get_inner_query()
        filters = self.get_filters()

        if len(filters) == 1:
            return {
                'filtered': {
                    'query': inner_query,
                    'filter': filters[0],
                }
            }
        elif len(filters) > 1:
            return {
                'filtered': {
                    'query': inner_query,
                    'filter': {
                        'and': filters,
                    }
                }
            }
        else:
            return inner_query

    def get_sort(self):
        # Ordering by relevance is the default in Elasticsearch
        if self.order_by_relevance:
            return

        # Get queryset and make sure its ordered
        if self.queryset.ordered:
            sort = []

            for reverse, field in self._get_order_by():
                column_name = self.mapping.get_field_column_name(field)

                sort.append({
                    column_name: 'desc' if reverse else 'asc'
                })

            return sort

        else:
            # Order by pk field
            return ['pk']

    def __repr__(self):
        return json.dumps(self.get_query())


class Elasticsearch2SearchResults(BaseSearchResults):
    fields_param_name = 'fields'

    def _get_es_body(self, for_count=False):
        body = {
            'query': self.query_compiler.get_query()
        }

        if not for_count:
            sort = self.query_compiler.get_sort()

            if sort is not None:
                body['sort'] = sort

        return body

    def _get_results_from_hits(self, hits):
        """
        Yields Django model instances from a page of hits returned by Elasticsearch
        """
        # Get pks from results
        pks = [hit['fields']['pk'][0] for hit in hits]
        scores = {str(hit['fields']['pk'][0]): hit['_score'] for hit in hits}

        # Initialise results dictionary
        results = {str(pk): None for pk in pks}

        # Find objects in database and add them to dict
        for obj in self.query_compiler.queryset.filter(pk__in=pks):
            results[str(obj.pk)] = obj

            if self._score_field:
                setattr(obj, self._score_field, scores.get(str(obj.pk)))

        # Yield results in order given by Elasticsearch
        for pk in pks:
            result = results[str(pk)]
            if result:
                yield result

    def _do_search(self):
        PAGE_SIZE = 100

        if self.stop is not None:
            limit = self.stop - self.start
        else:
            limit = None

        use_scroll = limit is None or limit > PAGE_SIZE

        params = {
            'index': self.backend.get_index_for_model(self.query_compiler.queryset.model).name,
            'body': self._get_es_body(),
            '_source': False,
            self.fields_param_name: 'pk',
        }

        if use_scroll:
            params.update({
                'scroll': '2m',
                'size': PAGE_SIZE,
            })

            # The scroll API doesn't support offset, manually skip the first results
            skip = self.start

            # Send to Elasticsearch
            page = self.backend.es.search(**params)

            while True:
                hits = page['hits']['hits']

                if len(hits) == 0:
                    break

                # Get results
                if skip < len(hits):
                    for result in self._get_results_from_hits(hits):
                        if limit is not None and limit == 0:
                            break

                        if skip == 0:
                            yield result

                            if limit is not None:
                                limit -= 1
                        else:
                            skip -= 1

                    if limit is not None and limit == 0:
                        break
                else:
                    # Skip whole page
                    skip -= len(hits)

                # Fetch next page of results
                if '_scroll_id' not in page:
                    break

                page = self.backend.es.scroll(scroll_id=page['_scroll_id'], scroll='2m')

            # Clear the scroll
            if '_scroll_id' in page:
                self.backend.es.clear_scroll(scroll_id=page['_scroll_id'])
        else:
            params.update({
                'from_': self.start,
                'size': limit or PAGE_SIZE,
            })

            # Send to Elasticsearch
            hits = self.backend.es.search(**params)['hits']['hits']

            # Get results
            for result in self._get_results_from_hits(hits):
                yield result

    def _do_count(self):
        # Get count
        hit_count = self.backend.es.count(
            index=self.backend.get_index_for_model(self.query_compiler.queryset.model).name,
            body=self._get_es_body(for_count=True),
        )['count']

        # Add limits
        hit_count -= self.start
        if self.stop is not None:
            hit_count = min(hit_count, self.stop - self.start)

        return max(hit_count, 0)


class Elasticsearch2Index:
    def __init__(self, backend, name):
        self.backend = backend
        self.es = backend.es
        self.mapping_class = backend.mapping_class
        self.name = name

    def put(self):
        self.es.indices.create(self.name, self.backend.settings)

    def delete(self):
        try:
            self.es.indices.delete(self.name)
        except NotFoundError:
            pass

    def exists(self):
        return self.es.indices.exists(self.name)

    def is_alias(self):
        return self.es.indices.exists_alias(name=self.name)

    def aliased_indices(self):
        """
        If this index object represents an alias (which appear the same in the
        Elasticsearch API), this method can be used to fetch the list of indices
        the alias points to.

        Use the is_alias method if you need to find out if this an alias. This
        returns an empty list if called on an index.
        """
        return [
            self.backend.index_class(self.backend, index_name)
            for index_name in self.es.indices.get_alias(name=self.name).keys()
        ]

    def put_alias(self, name):
        """
        Creates a new alias to this index. If the alias already exists it will
        be repointed to this index.
        """
        self.es.indices.put_alias(name=name, index=self.name)

    def add_model(self, model):
        # Get mapping
        mapping = self.mapping_class(model)

        # Put mapping
        self.es.indices.put_mapping(
            # pass update_all_types=True as a workaround to avoid "Can't redefine search field" errors -
            # see https://github.com/wagtail/wagtail/issues/2968
            index=self.name, doc_type=mapping.get_document_type(), body=mapping.get_mapping(),
            update_all_types=True
        )

    def add_item(self, item):
        # Make sure the object can be indexed
        if not class_is_indexed(item.__class__):
            return

        # Get mapping
        mapping = self.mapping_class(item.__class__)

        # Add document to index
        self.es.index(
            self.name, mapping.get_document_type(), mapping.get_document(item), id=mapping.get_document_id(item)
        )

    def add_items(self, model, items):
        if not class_is_indexed(model):
            return

        # Get mapping
        mapping = self.mapping_class(model)
        doc_type = mapping.get_document_type()

        # Create list of actions
        actions = []
        for item in items:
            # Create the action
            action = {
                '_type': doc_type,
                '_id': mapping.get_document_id(item),
            }
            action.update(mapping.get_document(item))
            actions.append(action)

        # Run the actions
        bulk(self.es, actions, index=self.name)

    def delete_item(self, item):
        # Make sure the object can be indexed
        if not class_is_indexed(item.__class__):
            return

        # Get mapping
        mapping = self.mapping_class(item.__class__)

        # Delete document
        try:
            self.es.delete(
                self.name,
                mapping.get_document_type(),
                mapping.get_document_id(item),
            )
        except NotFoundError:
            pass  # Document doesn't exist, ignore this exception

    def refresh(self):
        self.es.indices.refresh(self.name)

    def reset(self):
        # Delete old index
        self.delete()

        # Create new index
        self.put()


class ElasticsearchIndexRebuilder:
    def __init__(self, index):
        self.index = index

    def reset_index(self):
        self.index.reset()

    def start(self):
        # Reset the index
        self.reset_index()

        return self.index

    def finish(self):
        self.index.refresh()


class ElasticsearchAtomicIndexRebuilder(ElasticsearchIndexRebuilder):
    def __init__(self, index):
        self.alias = index
        self.index = index.backend.index_class(
            index.backend,
            self.alias.name + '_' + get_random_string(7).lower()
        )

    def reset_index(self):
        # Delete old index using the alias
        # This should delete both the alias and the index
        self.alias.delete()

        # Create new index
        self.index.put()

        # Create a new alias
        self.index.put_alias(self.alias.name)

    def start(self):
        # Create the new index
        self.index.put()

        return self.index

    def finish(self):
        self.index.refresh()

        if self.alias.is_alias():
            # Update existing alias, then delete the old index

            # Find index that alias currently points to, we'll delete it after
            # updating the alias
            old_index = self.alias.aliased_indices()

            # Update alias to point to new index
            self.index.put_alias(self.alias.name)

            # Delete old index
            # aliased_indices() can return multiple indices. Delete them all
            for index in old_index:
                if index.name != self.index.name:
                    index.delete()

        else:
            # self.alias doesn't currently refer to an alias in Elasticsearch.
            # This means that either nothing exists in ES with that name or
            # there is currently an index with the that name

            # Run delete on the alias, just in case it is currently an index.
            # This happens on the first rebuild after switching ATOMIC_REBUILD on
            self.alias.delete()

            # Create the alias
            self.index.put_alias(self.alias.name)


class Elasticsearch2SearchBackend(BaseSearchBackend):
    index_class = Elasticsearch2Index
    query_compiler_class = Elasticsearch2SearchQueryCompiler
    results_class = Elasticsearch2SearchResults
    mapping_class = Elasticsearch2Mapping
    basic_rebuilder_class = ElasticsearchIndexRebuilder
    atomic_rebuilder_class = ElasticsearchAtomicIndexRebuilder

    settings = {
        'settings': {
            'analysis': {
                'analyzer': {
                    'ngram_analyzer': {
                        'type': 'custom',
                        'tokenizer': 'lowercase',
                        'filter': ['asciifolding', 'ngram']
                    },
                    'edgengram_analyzer': {
                        'type': 'custom',
                        'tokenizer': 'lowercase',
                        'filter': ['asciifolding', 'edgengram']
                    }
                },
                'tokenizer': {
                    'ngram_tokenizer': {
                        'type': 'nGram',
                        'min_gram': 3,
                        'max_gram': 15,
                    },
                    'edgengram_tokenizer': {
                        'type': 'edgeNGram',
                        'min_gram': 2,
                        'max_gram': 15,
                        'side': 'front'
                    }
                },
                'filter': {
                    'ngram': {
                        'type': 'nGram',
                        'min_gram': 3,
                        'max_gram': 15
                    },
                    'edgengram': {
                        'type': 'edgeNGram',
                        'min_gram': 1,
                        'max_gram': 15
                    }
                }
            }
        }
    }

    def __init__(self, params):
        super(Elasticsearch2SearchBackend, self).__init__(params)

        # Get settings
        self.hosts = params.pop('HOSTS', None)
        self.index_name = params.pop('INDEX', 'wagtail')
        self.timeout = params.pop('TIMEOUT', 10)

        if params.pop('ATOMIC_REBUILD', False):
            self.rebuilder_class = self.atomic_rebuilder_class
        else:
            self.rebuilder_class = self.basic_rebuilder_class

        # If HOSTS is not set, convert URLS setting to HOSTS
        es_urls = params.pop('URLS', ['http://localhost:9200'])
        if self.hosts is None:
            self.hosts = []

            for url in es_urls:
                parsed_url = urlparse(url)

                use_ssl = parsed_url.scheme == 'https'
                port = parsed_url.port or (443 if use_ssl else 80)

                http_auth = None
                if parsed_url.username is not None and parsed_url.password is not None:
                    http_auth = (parsed_url.username, parsed_url.password)

                self.hosts.append({
                    'host': parsed_url.hostname,
                    'port': port,
                    'url_prefix': parsed_url.path,
                    'use_ssl': use_ssl,
                    'verify_certs': use_ssl,
                    'http_auth': http_auth,
                })

        self.settings = copy.deepcopy(self.settings)  # Make the class settings attribute as instance settings attribute
        self.settings = deep_update(self.settings, params.pop("INDEX_SETTINGS", {}))

        # Get Elasticsearch interface
        # Any remaining params are passed into the Elasticsearch constructor
        options = params.pop('OPTIONS', {})

        self.es = Elasticsearch(
            hosts=self.hosts,
            timeout=self.timeout,
            **options)

    def get_index_for_model(self, model):
        # Split models up into separate indices based on their root model.
        # For example, all page-derived models get put together in one index,
        # while images and documents each have their own index.
        root_model = get_model_root(model)
        index_suffix = '__' + root_model._meta.app_label.lower() + '_' + root_model.__name__.lower()

        return self.index_class(self, self.index_name + index_suffix)

    def get_index(self):
        return self.index_class(self, self.index_name)

    def get_rebuilder(self):
        return self.rebuilder_class(self.get_index())

    def reset_index(self):
        # Use the rebuilder to reset the index
        self.get_rebuilder().reset_index()

    def add_type(self, model):
        warnings.warn(
            "The `backend.add_type(model)` method is deprecated. "
            "Please use `backend.get_index_for_model(model).add_model(model)` instead.",
            category=RemovedInWagtail22Warning
        )

        self.get_index_for_model(model).add_model(model)

    def refresh_index(self):
        warnings.warn(
            "The `backend.refresh_index()` method is deprecated. "
            "Please use `backend.get_index_for_model(model).refresh()` for each model instead.",
            category=RemovedInWagtail22Warning
        )

        self.get_index().refresh()

    def add(self, obj):
        warnings.warn(
            "The `backend.add(obj)` method is deprecated. "
            "Please use `backend.get_index_for_model(type(obj)).add_item(obj)` instead.",
            category=RemovedInWagtail22Warning
        )

        self.get_index_for_model(type(obj)).add_item(obj)

    def add_bulk(self, model, obj_list):
        warnings.warn(
            "The `backend.add_bulk(model, obj_list)` method is deprecated. "
            "Please use `self.get_index_for_model(model).add_items(model, obj_list)` instead.",
            category=RemovedInWagtail22Warning
        )

        self.get_index_for_model(model).add_items(model, obj_list)

    def delete(self, obj):
        warnings.warn(
            "The `backend.delete(obj)` method is deprecated. "
            "Please use `backend.get_index_for_model(type(obj)).delete_item(obj)` instead.",
            category=RemovedInWagtail22Warning
        )

        self.get_index_for_model(type(obj)).delete_item(obj)


SearchBackend = Elasticsearch2SearchBackend
