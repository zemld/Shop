// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/users/balance": {
            "get": {
                "description": "Returns user's balance.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User whose balance you want to get",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/change-balance": {
            "post": {
                "description": "Changes the balance of a registered user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User whose balance you want to change",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "New balance of the user",
                        "name": "balance",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/register": {
            "post": {
                "description": "Registers a new user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User which you want to register",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Balance of the user",
                        "name": "balance",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/{user}": {
            "get": {
                "description": "Checks if user is registered.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User which you want to check",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegistered"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.StatusResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.StatusResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "dto.UserRegistered": {
            "type": "object",
            "properties": {
                "is_registered": {
                    "type": "boolean"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "User Service",
	Description:      "API for interacting with users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
