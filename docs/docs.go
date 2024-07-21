// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/decode/join/{roomId}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "join the room with roomId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "Join Room",
                "operationId": "join-room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "roomId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-string"
                        }
                    }
                }
            }
        },
        "/decode/task/{taskId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "get task details by taskId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get Task Details",
                "operationId": "get-task-details-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Response-types_TaskModel"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate a user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "loginInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "types.Response-string": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "types.Response-types_TaskModel": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/types.TaskModel"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "types.TaskModel": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "handler": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "user.LoginReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}