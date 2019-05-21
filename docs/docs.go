// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-20 18:05:16.710162 +0600 +06 m=+0.024332514

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Serves changing in diagram and upload/update it in GitHub",
        "title": "Camunda Modeller GitHub",
        "termsOfService": "http://halykbank.kz",
        "contact": {
            "name": "Nartay Dembayev",
            "url": "http://instagram.com/nartaymalikov",
            "email": "mastaok02@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
        "/deployment/create": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Serves changing in diagram and upload/update it in GitHub",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Camunda Modeller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Repository name",
                        "name": "deployment-name",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Camunda Modeller",
                        "name": "diagram_1.bpmn",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "answer",
                        "schema": {
                            "type": "string"
                        },
                        "headers": {
                            "string": {
                                "type": "string",
                                "description": "Header"
                            }
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}