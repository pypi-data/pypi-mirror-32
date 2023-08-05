# -*- coding: utf-8 -*-
import json

import typing
import marshmallow
import yaml

from apispec import APISpec
from apispec import Path
from apispec.ext.marshmallow.swagger import schema2jsonschema

from hapic.context import ContextInterface
from hapic.context import RouteRepresentation
from hapic.decorator import DecoratedController
from hapic.description import ControllerDescription


def generate_schema_ref(spec:APISpec, schema: marshmallow.Schema) -> str:
    schema_class = spec.schema_class_resolver(
        spec,
        schema
    )
    ref = {
        '$ref': '#/definitions/{}'.format(
            spec.schema_name_resolver(schema_class)
        )
    }
    if schema.many:
        ref = {
            'type': 'array',
            'items': ref
        }
    return ref


def bottle_generate_operations(
    spec,
    route: RouteRepresentation,
    description: ControllerDescription,
):
    method_operations = dict()
    if description.input_body:
        method_operations.setdefault('parameters', []).append({
            'in': 'body',
            'name': 'body',
            'schema': generate_schema_ref(
                spec,
                description.input_body.wrapper.processor.schema,
            )
        })

    if description.output_body:
        method_operations.setdefault('responses', {})\
            [int(description.output_body.wrapper.default_http_code)] = {
                'description': str(int(description.output_body.wrapper.default_http_code)),  # nopep8
                'schema': generate_schema_ref(
                    spec,
                    description.output_body.wrapper.processor.schema,
                )
            }

    if description.output_file:
        method_operations.setdefault('produces', []).extend(
            description.output_file.wrapper.output_types
        )
        method_operations.setdefault('responses', {})\
            [int(description.output_file.wrapper.default_http_code)] = {
            'description': str(int(description.output_file.wrapper.default_http_code)),  # nopep8
        }

    if description.errors:
        for error in description.errors:
            schema_class = type(error.wrapper.error_builder)
            method_operations.setdefault('responses', {})\
                [int(error.wrapper.http_code)] = {
                    'description': str(int(error.wrapper.http_code)),
                    'schema': {
                        '$ref': '#/definitions/{}'.format(
                            spec.schema_name_resolver(schema_class)
                        )
                    }
                }

    # jsonschema based
    if description.input_path:
        schema_class = spec.schema_class_resolver(
            spec,
            description.input_path.wrapper.processor.schema
        )
        # TODO: look schema2parameters ?
        jsonschema = schema2jsonschema(schema_class, spec=spec)
        for name, schema in jsonschema.get('properties', {}).items():
            method_operations.setdefault('parameters', []).append({
                'in': 'path',
                'name': name,
                'required': name in jsonschema.get('required', []),
                'type': schema['type']
            })

    if description.input_query:
        schema_class = spec.schema_class_resolver(
            spec,
            description.input_query.wrapper.processor.schema
        )
        jsonschema = schema2jsonschema(schema_class, spec=spec)
        for name, schema in jsonschema.get('properties', {}).items():
            method_operations.setdefault('parameters', []).append({
                'in': 'query',
                'name': name,
                'required': name in jsonschema.get('required', []),
                'type': schema['type']
            })

    if description.input_files:
        method_operations.setdefault('consumes', []).append('multipart/form-data')
        for field_name, field in description.input_files.wrapper.processor.schema.fields.items():
            method_operations.setdefault('parameters', []).append({
                'in': 'formData',
                'name': field_name,
                'required': field.required,
                'type': 'file',
            })

    if description.tags:
        method_operations['tags'] = description.tags

    operations = {
        route.method.lower(): method_operations,
    }

    return operations


class DocGenerator(object):
    def get_doc(
        self,
        controllers: typing.List[DecoratedController],
        context: ContextInterface,
        title: str='',
        description: str='',
    ) -> dict:
        """
        Generate an OpenApi 2.0 documentation. Th given context will be used
        to found controllers matching with given DecoratedController.
        :param controllers: List of DecoratedController to match with context
        controllers
        :param context: a context instance
        :param title: The generated doc title
        :param description: The generated doc description
        :return: a apispec documentation dict
        """
        spec = APISpec(
            title=title,
            info=dict(description=description),
            version='1.0.0',
            plugins=(
                'apispec.ext.marshmallow',
            ),
            auto_referencing=True,
            schema_name_resolver=generate_schema_name
        )

        schemas = []
        # parse schemas
        for controller in controllers:
            description = controller.description

            if description.input_body:
                schemas.append(spec.schema_class_resolver(
                    spec,
                    description.input_body.wrapper.processor.schema
                ))

            if description.input_forms:
                schemas.append(spec.schema_class_resolver(
                    spec,
                    description.input_forms.wrapper.processor.schema
                ))

            if description.output_body:
                schemas.append(spec.schema_class_resolver(
                    spec,
                    description.output_body.wrapper.processor.schema
                ))

            if description.errors:
                for error in description.errors:
                    schemas.append(type(error.wrapper.error_builder))

        for schema in set(schemas):
            spec.definition(
                spec.schema_name_resolver(schema),
                schema=schema
            )

        # add views
        # with app.test_request_context():
        paths = {}
        for controller in controllers:
            route = context.find_route(controller)
            swagger_path = context.get_swagger_path(route.rule)

            operations = bottle_generate_operations(
                spec,
                route,
                controller.description,
            )

            # TODO BS 20171114: TMP code waiting refact of doc
            doc_string = controller.reference.get_doc_string()
            if doc_string:
                for method in operations.keys():
                    operations[method]['description'] = doc_string

            path = Path(path=swagger_path, operations=operations)

            if swagger_path in paths:
                paths[swagger_path].update(path)
            else:
                paths[swagger_path] = path

            spec.add_path(path)

        return spec.to_dict()

    def get_doc_yaml(
        self,
        controllers: typing.List[DecoratedController],
        context: ContextInterface,
        title: str = '',
        description: str = '',
    ) -> str:
        dict_doc = self.get_doc(
            controllers=controllers,
            context=context,
            title=title,
            description=description,
        )
        json_doc = json.dumps(dict_doc)

        # We dump then load with json to use real scalar dict.
        # If not, yaml dump dict-like objects
        clean_dict_doc = json.loads(json_doc)
        return yaml.dump(clean_dict_doc, default_flow_style=False)

    def save_in_file(
        self,
        doc_file_path: str,
        controllers: typing.List[DecoratedController],
        context: ContextInterface,
        title: str='',
        description: str='',
    ) -> None:
        doc_yaml = self.get_doc_yaml(
            controllers=controllers,
            context=context,
            title=title,
            description=description,
        )
        with open(doc_file_path, 'w+') as doc_file:
            doc_file.write(doc_yaml)


# TODO BS 20171109: Must take care of already existing definition names
def generate_schema_name(schema: marshmallow.Schema):
    """
    Return best candidate name for one schema cls or instance.
    :param schema: instance or cls schema
    :return: best schema name
    """
    if not isinstance(schema, type):
        schema = type(schema)

    if getattr(schema, '_schema_name', None):
        if schema.opts.exclude:
            schema_name = "{}_without".format(schema.__name__)
            for elem in sorted(schema.opts.exclude):
                schema_name="{}_{}".format(schema_name, elem)
        else:
            schema_name = schema._schema_name
    else:
        schema_name = schema.__name__

    return schema_name
