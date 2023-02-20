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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Confidence James",
            "url": "https://github.com/jamesconfy",
            "email": "bobdence@gmail.com"
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
        "/": {
            "get": {
                "description": "Question route",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get questions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "category",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "difficulty",
                        "name": "difficulty",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Question"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Question": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "correctAnswer": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "incorrectAnswers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "isNiche": {
                    "type": "boolean"
                },
                "question": {
                    "type": "string"
                },
                "regions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "confy-quiz-api.fly.dev",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Quiz-API",
	Description:      "This is an api that brings out random quiz",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
