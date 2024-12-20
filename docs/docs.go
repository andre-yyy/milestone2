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
        "/categories": {
            "get": {
                "description": "Get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Category"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/coins": {
            "get": {
                "description": "Get all coins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all coins",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Coin"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/coins/{coinID}": {
            "get": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Get coin by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get coin by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "coinID",
                        "name": "coinID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.TopupInvoice-entity_Coin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Get all orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Order"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Create order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Create order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateOrderDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-entity_Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/orders/overdue": {
            "get": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Get all overdue orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all overdue orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Order"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/orders/{orderID}": {
            "patch": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Update order status by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Update order status by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "orderID",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateOrderStatusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-entity_Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Product"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/products/{productID}": {
            "get": {
                "description": "Get product by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "productID",
                        "name": "productID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.TopupInvoice-entity_Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/topups": {
            "get": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Get all topups",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all topups",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Topup"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/users/activities": {
            "get": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Get all user activities",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Get all user activities",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-array_entity_Activity"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-entity_User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "security": [
                    {
                        "Auth": []
                    }
                ],
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Milestone 2"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponseData-entity_User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.ServerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateOrderDto": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "rent_days": {
                    "type": "integer"
                }
            }
        },
        "dto.LoginUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateOrderStatusDto": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "entity.Activity": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.Coin": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "entity.Invoice": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "invoice_url": {
                    "type": "string"
                }
            }
        },
        "entity.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "destination_id": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "rent_days": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "total_price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "daily": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rental_cost": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "entity.ServerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Activity": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Activity"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Category": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Category"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Coin": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Coin"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Order": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Order"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Product": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Product"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-array_entity_Topup": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Topup"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-entity_Order": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/entity.Order"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.ServerResponseData-entity_User": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/entity.User"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Topup": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "external_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.TopupInvoice-entity_Coin": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/entity.Coin"
                },
                "invoice": {
                    "$ref": "#/definitions/entity.Invoice"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.TopupInvoice-entity_Product": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/entity.Product"
                },
                "invoice": {
                    "$ref": "#/definitions/entity.Invoice"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deposit": {
                    "type": "number"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Auth": {
            "description": "Authentication for Milestone 2 API",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Milestone 2 API",
	Description:      "Milestone 2 API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}