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
            "name": "Lawrence Matsuyama",
            "email": "law.matsuyama@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/get": {
            "post": {
                "description": "List transactions by giving filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "API to get transactions in the application.",
                "parameters": [
                    {
                        "description": "Transactions Get Request",
                        "name": "transactions_get_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimanager.TransactionsGetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-domain_TransactionsPaging"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-string"
                        }
                    }
                }
            }
        },
        "/v1/save": {
            "post": {
                "description": "Receives transactions data, registed it in application and finish notifying other applications.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "API to save transactions in the application.",
                "parameters": [
                    {
                        "description": "Transactions Save Request",
                        "name": "transactions_save_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimanager.TransactionsSaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-array_apimanager_TransactionSaveResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/apimanager.GenericResponse-array_apimanager_TransactionSaveResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apimanager.GenericResponse-array_apimanager_TransactionSaveResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimanager.TransactionSaveResponse"
                    }
                }
            }
        },
        "apimanager.GenericResponse-domain_TransactionsPaging": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {
                    "$ref": "#/definitions/domain.TransactionsPaging"
                }
            }
        },
        "apimanager.GenericResponse-string": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                }
            }
        },
        "apimanager.Paging": {
            "type": "object",
            "properties": {
                "next_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "apimanager.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "operation": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "apimanager.TransactionSaveRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "operation": {
                    "type": "string"
                }
            }
        },
        "apimanager.TransactionSaveResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "transaction": {
                    "$ref": "#/definitions/apimanager.Transaction"
                }
            }
        },
        "apimanager.TransactionsGetRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "amount_greater": {
                    "type": "number"
                },
                "amount_less": {
                    "type": "number"
                },
                "date_from": {
                    "type": "string"
                },
                "date_to": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "operation_type": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "paging": {
                    "$ref": "#/definitions/apimanager.Paging"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "apimanager.TransactionsSaveRequest": {
            "type": "object",
            "properties": {
                "origin_channel": {
                    "type": "string"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apimanager.TransactionSaveRequest"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.OperationType": {
            "type": "string",
            "enum": [
                "debit",
                "credit"
            ],
            "x-enum-varnames": [
                "DebitOperation",
                "CreditOperation"
            ]
        },
        "domain.OriginChannel": {
            "type": "string",
            "enum": [
                "desktop-web",
                "mobile-android",
                "mobile-ios"
            ],
            "x-enum-varnames": [
                "DesktopWeb",
                "MobileAndroid",
                "MobileIos"
            ]
        },
        "domain.Paging": {
            "type": "object",
            "properties": {
                "next_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "domain.Transaction": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "operation_type": {
                    "$ref": "#/definitions/domain.OperationType"
                },
                "origin": {
                    "$ref": "#/definitions/domain.OriginChannel"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.TransactionsPaging": {
            "type": "object",
            "properties": {
                "paging": {
                    "$ref": "#/definitions/domain.Paging"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Transaction"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/transactions",
	Schemes:          []string{},
	Title:            "Swagger Transactions API",
	Description:      "API to save and list user transactions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
