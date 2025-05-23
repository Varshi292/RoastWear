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
        "/cart/items": {
            "get": {
                "description": "Fetches all cart items associated with a username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cart"
                ],
                "summary": "Get cart items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CartItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageResponse"
                        }
                    }
                }
            }
        },
        "/cart/modify": {
            "post": {
                "description": "Adds an item to the user's cart, updates quantity, or deletes it if quantity is 0.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cart"
                ],
                "summary": "Modify user's cart",
                "parameters": [
                    {
                        "description": "Cart item details",
                        "name": "cartItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ModifyCartRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.MessageResponse"
                        }
                    }
                }
            }
        },
        "/checkout": {
            "post": {
                "description": "Processes the items in the user's cart, creates a purchase record, and clears the cart.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "purchase"
                ],
                "summary": "Checkout and complete purchase",
                "parameters": [
                    {
                        "description": "Checkout request with username",
                        "name": "checkoutRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CheckoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Purchase completed and cart cleared",
                        "schema": {
                            "$ref": "#/definitions/handlers.PurchaseResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input or empty cart",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error during purchase processing",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Destroys the user session and logs the user out",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logs out a user",
                "responses": {
                    "200": {
                        "description": "Logged out successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user by creating an account with the provided username, email, and password. If the user already exists (either by username or email), an error is returned. If there are missing fields, a validation error is returned. Includes a detailed response for successful registration or error scenarios.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "User registration details (username, email, and password)",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request, invalid request format",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "409": {
                        "description": "Conflict error, username or email already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "422": {
                        "description": "Validation error, missing required fields (username, email, or password)",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/session/create": {
            "post": {
                "description": "Stores a session in the database (used for login/session tracking)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Create a new user session",
                "parameters": [
                    {
                        "description": "Session data",
                        "name": "session",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionDoc"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Session successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid session data",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create session",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/session/delete": {
            "delete": {
                "description": "Removes a session from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Delete session",
                "parameters": [
                    {
                        "description": "Session details",
                        "name": "session",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionDoc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/session/verify": {
            "post": {
                "description": "Validates a session by username and session ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Verify session",
                "parameters": [
                    {
                        "description": "Session details",
                        "name": "session",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SessionDoc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.GenericResponse"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "Allows a user to upload a PNG image with a username. The image is stored on the server and the metadata is saved to the database.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Images"
                ],
                "summary": "Upload a PNG image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user uploading the image",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "PNG image file",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UploadSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CheckoutRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "handlers.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "handlers.ModifyCartRequest": {
            "type": "object",
            "properties": {
                "productid": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "number"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.PurchaseResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.CartItem": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "productid": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.GenericResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.SessionDoc": {
            "description": "This is the session model used for Swagger documentation",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "session_key": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UploadSuccessResponse": {
            "type": "object",
            "properties": {
                "filepath": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UserRegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7777",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "RoastWear API",
	Description:      "This is the backend API for the RoastWear application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
