{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Basic",
                "None",
                "OAuth"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": false,
            "default_value": "None",
            "deprecated": null,
            "description": "AuthType refers to the type of HTTP authentication used to query endpoints",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "authType",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Endpoint is the API end point of the service",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "endpoint",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Password is the password of the user to be used in the HTTP Authorization header",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "password",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Registry",
                "VulnerabilityScanner"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Registry",
            "deprecated": null,
            "description": "Type refers to type of the server",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "type",
            "orderable": true,
            "primary_key": false,
            "read_only": null,
            "required": true,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Username refers to the username to be used in the HTTP Authorization header",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "userName",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": "Allows to declare a file, a folder or volume that can be used in File Access Policies to allow a processing unit to read, write or execute the represented item.",
        "entity_name": "Integration",
        "extends": [
            "@base",
            "@identifiable-pk-stored"
        ],
        "get": true,
        "package": "Infrastructure",
        "resource_name": "integrations",
        "rest_name": "integration",
        "root": null,
        "update": true
    }
}