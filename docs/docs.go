// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/anime": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Anime"
                ],
                "summary": "Call All Data Anime",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Anime"
                ],
                "summary": "Add New Anime",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Anime"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/family": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Family"
                ],
                "summary": "Call All Data Family",
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Family"
                ],
                "summary": "Add New Family",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Family"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/family/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Family"
                ],
                "summary": "Call Single Data Family",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search family by id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {}
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Family"
                ],
                "summary": "Delete Family",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete family",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Anime": {
            "type": "object",
            "properties": {
                "condition": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "episode": {
                    "type": "integer"
                },
                "family_id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "rate": {
                    "type": "integer"
                },
                "season": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.Family": {
            "type": "object",
            "properties": {
                "altName": {
                    "type": "string"
                },
                "avgRate": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "totalDuration": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
