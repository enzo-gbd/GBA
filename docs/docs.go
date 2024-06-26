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
        "/login": {
            "post": {
                "description": "Logs in a user and returns an access token and a refresh token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "User Login Data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Returns the access token and refresh token on successful login",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Returns error message for bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Returns error message for invalid email or password",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Returns error message for internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Logs out a user by invalidating the tokens.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Logout a user",
                "responses": {
                    "200": {
                        "description": "Returns status success on successful logout",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Refreshes an access token using the refresh token provided in cookie.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Refresh access token",
                "responses": {
                    "200": {
                        "description": "Returns new access token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Returns error message for unauthorized or invalid token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Returns error message when the user does not exist",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Returns error message for internal server error",
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
                "description": "Registers a new user with the necessary details provided in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User Registration Data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Returns status success on successful registration",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Returns error message for bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "409": {
                        "description": "Returns error message for email already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Returns error message for internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Fetches all users from the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves information about the current logged in user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get current user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Fetches a user from the database by UUID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
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
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates details of an existing user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
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
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user by UUID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.SignInInput": {
            "description": "Fields required to login a user.",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Email address of the user",
                    "type": "string"
                },
                "password": {
                    "description": "Password for the user account",
                    "type": "string"
                }
            }
        },
        "models.SignUpInput": {
            "description": "Fields required to register a new user.",
            "type": "object",
            "required": [
                "birthday",
                "email",
                "first_name",
                "gender",
                "name",
                "password"
            ],
            "properties": {
                "birthday": {
                    "description": "Birthday of the user",
                    "type": "string"
                },
                "email": {
                    "description": "Email address of the user",
                    "type": "string"
                },
                "first_name": {
                    "description": "First name of the user",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender of the user",
                    "type": "string"
                },
                "name": {
                    "description": "Last name of the user",
                    "type": "string"
                },
                "password": {
                    "description": "Password for the user account",
                    "type": "string"
                }
            }
        },
        "models.User": {
            "description": "User holds all details about a user.",
            "type": "object",
            "properties": {
                "address": {
                    "description": "Optional address of the user",
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.NullString"
                        }
                    ]
                },
                "birthday": {
                    "description": "Birthday of the user",
                    "type": "string"
                },
                "createdAt": {
                    "description": "Timestamp when the user was created",
                    "type": "string"
                },
                "deletedAt": {
                    "description": "Optional timestamp when the user was deleted",
                    "type": "string"
                },
                "email": {
                    "description": "Email address of the user, must be unique",
                    "type": "string"
                },
                "firstName": {
                    "description": "First name of the user",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender of the user",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier for the user",
                    "type": "string"
                },
                "isActive": {
                    "description": "Flag indicating if the user account is active",
                    "type": "boolean"
                },
                "name": {
                    "description": "Last name of the user",
                    "type": "string"
                },
                "password": {
                    "description": "Password for the user account",
                    "type": "string"
                },
                "role": {
                    "description": "Role of the user in the system",
                    "type": "string"
                },
                "subscriptionCode": {
                    "description": "Optional subscription code",
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.NullString"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "Timestamp when the user was last updated",
                    "type": "string"
                },
                "verificationCode": {
                    "description": "Optional verification code",
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.NullString"
                        }
                    ]
                },
                "verified": {
                    "description": "Flag indicating if the user has verified their email",
                    "type": "boolean"
                }
            }
        },
        "models.UserResponse": {
            "description": "UserResponse holds the data that is exposed to the client after API requests.",
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address of the user (present even if empty as a blank string)",
                    "type": "string"
                },
                "birthday": {
                    "description": "Birthday of the user, omitted if empty",
                    "type": "string"
                },
                "created_at": {
                    "description": "Timestamp when the user was created",
                    "type": "string"
                },
                "email": {
                    "description": "Email address of the user, omitted if empty",
                    "type": "string"
                },
                "first_name": {
                    "description": "First name of the user, omitted if empty",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender of the user, omitted if empty",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier for the user, omitted if empty",
                    "type": "string"
                },
                "name": {
                    "description": "Last name of the user, omitted if empty",
                    "type": "string"
                },
                "role": {
                    "description": "Role of the user within the system, omitted if empty",
                    "type": "string"
                },
                "subscription_code": {
                    "description": "Subscription code related to the user's account (present even if empty as a blank string)",
                    "type": "string"
                },
                "updated_at": {
                    "description": "Timestamp when the user profile was last updated",
                    "type": "string"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
