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
        "/one-to-one/create": {
            "post": {
                "description": "Create a new weekly report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Create a new weekly report",
                "parameters": [
                    {
                        "description": "Weekly report object to be created",
                        "name": "report",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/one_to_one.CreateWeeklyReportRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Weekly report created successfully",
                        "schema": {
                            "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/one-to-one/report-to": {
            "get": {
                "description": "Get a weekly report by week and year for a reportTo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Get a weekly report by week and year for a reportTo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Week number",
                        "name": "week",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Weekly report",
                        "schema": {
                            "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/one-to-one/report-to/all": {
            "get": {
                "description": "Get all weekly reports",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Get all weekly reports for a reportTo",
                "responses": {
                    "200": {
                        "description": "List of weekly reports",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                            }
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
        "/one-to-one/report-to/update": {
            "put": {
                "description": "Update a weekly report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Update a weekly report for a reportTo",
                "parameters": [
                    {
                        "description": "Weekly report object to be updated",
                        "name": "report",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/one_to_one.UpdateWeeklyReportRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Weekly report updated successfully",
                        "schema": {
                            "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/one-to-one/reportee": {
            "get": {
                "description": "Get a weekly report by week and year for a reportee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Get a weekly report by week and year for a reportee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Week number",
                        "name": "week",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Weekly report",
                        "schema": {
                            "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/one-to-one/reportee/all": {
            "get": {
                "description": "Get all weekly reports",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Get all weekly reports for a reportee",
                "responses": {
                    "200": {
                        "description": "List of weekly reports",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                            }
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
        "/one-to-one/reportee/update": {
            "put": {
                "description": "Update a weekly report",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "one-to-one"
                ],
                "summary": "Update a weekly report for a reportee",
                "parameters": [
                    {
                        "description": "Weekly report object to be updated",
                        "name": "report",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/one_to_one.UpdateWeeklyReportRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Weekly report updated successfully",
                        "schema": {
                            "$ref": "#/definitions/one_to_one.WeeklyReportResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/user/all": {
            "get": {
                "description": "Get all users",
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
                        "description": "Users retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.UserResponse"
                            }
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
        "/user/create": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User object to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
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
        "/user/current": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get current user",
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
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
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
        "/user/email/{email}": {
            "get": {
                "description": "Get user by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
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
        "/user/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
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
        "/user/reportee/add": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add reportee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add reportee",
                "parameters": [
                    {
                        "description": "Reportee object to be added",
                        "name": "reportee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AddReporteeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reportee added successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Reportee not found",
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
        "/user/reportee/remove": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Remove reportee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Remove reportee",
                "parameters": [
                    {
                        "description": "Reportee object to be removed",
                        "name": "reportee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RemoveReporteeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reportee removed successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Reportee not found",
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
        "/user/reports-to/add": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add reports to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add reports to user",
                "parameters": [
                    {
                        "description": "Report object to be added",
                        "name": "report",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AddReportsToRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Report added successfully",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or parameters",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
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
        }
    },
    "definitions": {
        "one_to_one.Agenda": {
            "type": "object",
            "required": [
                "label"
            ],
            "properties": {
                "label": {
                    "type": "string"
                }
            }
        },
        "one_to_one.Challenges": {
            "type": "object",
            "required": [
                "label",
                "theme"
            ],
            "properties": {
                "label": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "one_to_one.CreateWeeklyReportRequest": {
            "type": "object",
            "required": [
                "agendas",
                "challenges",
                "goneWell",
                "week",
                "wellbeingScores",
                "year"
            ],
            "properties": {
                "agendas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Agenda"
                    }
                },
                "challenges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Challenges"
                    }
                },
                "goneWell": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.GoneWell"
                    }
                },
                "week": {
                    "type": "integer"
                },
                "wellbeingScores": {
                    "$ref": "#/definitions/one_to_one.WellbeingScores"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "one_to_one.GoneWell": {
            "type": "object",
            "required": [
                "label",
                "theme"
            ],
            "properties": {
                "label": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "one_to_one.UpdateWeeklyReportRequest": {
            "type": "object",
            "required": [
                "agendas",
                "challenges",
                "goneWell",
                "id",
                "week",
                "wellbeingScores",
                "year"
            ],
            "properties": {
                "agendas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Agenda"
                    }
                },
                "challenges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Challenges"
                    }
                },
                "goneWell": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.GoneWell"
                    }
                },
                "id": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                },
                "wellbeingScores": {
                    "$ref": "#/definitions/one_to_one.WellbeingScores"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "one_to_one.WeeklyReportResponse": {
            "type": "object",
            "properties": {
                "agendas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Agenda"
                    }
                },
                "challenges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.Challenges"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "goneWell": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/one_to_one.GoneWell"
                    }
                },
                "id": {
                    "type": "string"
                },
                "reportee": {
                    "type": "string"
                },
                "reportingTo": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                },
                "wellbeingScores": {
                    "$ref": "#/definitions/one_to_one.WellbeingScores"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "one_to_one.WellbeingScores": {
            "type": "object",
            "required": [
                "growth",
                "impactAndProductivity",
                "wellbeing",
                "workOverall",
                "workRelationships"
            ],
            "properties": {
                "growth": {
                    "type": "integer"
                },
                "impactAndProductivity": {
                    "type": "integer"
                },
                "wellbeing": {
                    "type": "integer"
                },
                "workOverall": {
                    "type": "integer"
                },
                "workRelationships": {
                    "type": "integer"
                }
            }
        },
        "user.AddReporteeRequest": {
            "type": "object",
            "required": [
                "reporteeEmail"
            ],
            "properties": {
                "reporteeEmail": {
                    "type": "string"
                }
            }
        },
        "user.AddReportsToRequest": {
            "type": "object",
            "required": [
                "reportsToEmail"
            ],
            "properties": {
                "reportsToEmail": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.UserResponse"
                }
            }
        },
        "user.RemoveReporteeRequest": {
            "type": "object",
            "required": [
                "reporteeEmail"
            ],
            "properties": {
                "reporteeEmail": {
                    "type": "string"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "reportees": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "reportsTo": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "one-to-one.backend.vercel.app",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "OneToOne API",
	Description:      "This is the REST API for OneToOne.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
