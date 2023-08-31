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
        "/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {}
            }
        },
        "/api/conference/health-check": {
            "post": {
                "description": "CONFERENCE SERVICE HEALTH CHECK",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR CONFERENCE SERVICE HEALTH CHECK",
                "operationId": "CONFERENCE-HEALTH-CHECK",
                "parameters": [
                    {
                        "description": "Request body for health check",
                        "name": "HEALTH-CHECK",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.HealthCheck"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.Response"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.Response"
                        }
                    }
                }
            }
        },
        "/api/conference/join-group-confernce": {
            "patch": {
                "description": "JOIN GROUP CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR JOIN GROUP CONFERENCE",
                "operationId": "JOIN-GROUP-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for join group conference",
                        "name": "JOIN-GROUP-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.JoinGroupConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinGroupConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinGroupConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinGroupConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/conference/join-private-conference": {
            "patch": {
                "description": "JOIN PRIVATE CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR JOIN PRIVATE CONFERENCE",
                "operationId": "JOIN-PRIVATE-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for join private conference",
                        "name": "JOIN-PRIVATE-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.JoinPrivateConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPrivateConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPrivateConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPrivateConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/conference/join-public-conference": {
            "patch": {
                "description": "JOIN PUBLIC CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR JOIN PUBLIC CONFERENCE",
                "operationId": "JOIN-PUBLIC-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for sjoin public conference",
                        "name": "JOIN-PUBLIC-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.JoinPublicConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPublicConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPublicConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.JoinPublicConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/conference/schedule-conference": {
            "post": {
                "description": "CONFERENCE SCHEDULE CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR SCHEDULE CONFERENCE",
                "operationId": "SCHEDULE-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for schedule conference",
                        "name": "SCHEDULE-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ScheduleConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.ScheduleConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.ScheduleConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.ScheduleConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/conference/start-private-conference": {
            "post": {
                "description": "START PRIVATE CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR START PRIVATE CONFERENCE",
                "operationId": "START-PRIVATE-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for start private conference",
                        "name": "START-PRIVATE-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.StartPrivateConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPrivateConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPrivateConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPrivateConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/conference/start-public-conference": {
            "post": {
                "description": "START PUBLIC CONFERENCE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CONFERENCE"
                ],
                "summary": "API FOR START PUBLIC CONFERENCE",
                "operationId": "START-PUBLIC-CONFERENCE",
                "parameters": [
                    {
                        "description": "Request body for start public conference",
                        "name": "START-PUBLIC-CONFERENCE",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.StartPublicConferenceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPublicConferenceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPublicConferenceResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/conference.StartPublicConferenceResponse"
                        }
                    }
                }
            }
        },
        "/api/user/login": {
            "post": {
                "description": "VERIFY THE EMAIL,PASSWORD AND GENERATE A JWT TOKEN AND SET IT TO A COOKIE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR USER LOGIN",
                "operationId": "USER-LOGIN",
                "parameters": [
                    {
                        "description": "Enter the email and password",
                        "name": "login_details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/user/logout": {
            "post": {
                "description": "LOGOUT USER AND ALSO CLEAR COOKIES",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR USER LOGOUT",
                "operationId": "USER-LOGOUT",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/otp": {
            "post": {
                "description": "CREATE A NEW USER WITH REQUIRED DETAILS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR NEW USER SIGN UP",
                "operationId": "SIGNUP-USER",
                "parameters": [
                    {
                        "description": "New user Details",
                        "name": "user_details",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.OtpValidation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.OtpSignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.OtpSignUpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pb.OtpSignUpResponse"
                        }
                    }
                }
            }
        },
        "/api/user/resend-otp": {
            "post": {
                "description": "RESEND OTP REQUEST",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR RESEND OTP REQUEST",
                "operationId": "RESEND-OTP",
                "parameters": [
                    {
                        "description": "enter phone number",
                        "name": "RESENDOTP",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.ResendOtp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ResendOtpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.ResendOtpResponse"
                        }
                    }
                }
            }
        },
        "/api/user/signup": {
            "post": {
                "description": "VERIFY THE OTP AND DATA INSERT TO THE DATABASE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR NEW USER SIGN UP OTP VERIFICATION",
                "operationId": "SIGNUP-USER-OTP-VERIFY",
                "parameters": [
                    {
                        "description": "otp",
                        "name": "user_details",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/pb.SignupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.SignupResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/pb.SignupResponse"
                        }
                    }
                }
            }
        },
        "/api/user/valid-name": {
            "post": {
                "description": "VERIFY USERNAME IF IT UNIQUE OR NOT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "USER"
                ],
                "summary": "API FOR USERNAME VALIDATION",
                "operationId": "USERNAME-VALIDATION",
                "parameters": [
                    {
                        "description": "Enter User Name",
                        "name": "USERNAME",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ValidName"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pb.ValidNameResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pb.ValidNameResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "conference.JoinGroupConferenceResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "conference.JoinPrivateConferenceResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "conference.JoinPublicConferenceResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "conference.Response": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "string"
                }
            }
        },
        "conference.ScheduleConferenceResponse": {
            "type": "object",
            "properties": {
                "ConferenceID": {
                    "type": "string"
                },
                "Result": {
                    "type": "string"
                }
            }
        },
        "conference.StartPrivateConferenceResponse": {
            "type": "object",
            "properties": {
                "Result": {
                    "type": "string"
                },
                "conferenceID": {
                    "type": "string"
                }
            }
        },
        "conference.StartPublicConferenceResponse": {
            "type": "object",
            "properties": {
                "ConferenceID": {
                    "type": "string"
                },
                "Result": {
                    "type": "string"
                }
            }
        },
        "models.HealthCheck": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "models.JoinGroupConferenceRequest": {
            "type": "object",
            "properties": {
                "conferenceID": {
                    "type": "string"
                },
                "groupID": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.JoinPrivateConferenceRequest": {
            "type": "object",
            "properties": {
                "conferenceID": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.JoinPublicConferenceRequest": {
            "type": "object",
            "properties": {
                "conferenceID": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequestBody": {
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
        "models.OtpValidation": {
            "type": "object",
            "properties": {
                "cPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.RegisterRequestBody": {
            "type": "object",
            "properties": {
                "cPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "referral": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "models.ResendOtp": {
            "type": "object",
            "properties": {
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "models.ScheduleConferenceRequest": {
            "type": "object",
            "properties": {
                "broadcast": {
                    "type": "boolean"
                },
                "chat": {
                    "type": "boolean"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "interest": {
                    "type": "string"
                },
                "participantlimit": {
                    "type": "integer"
                },
                "recording": {
                    "type": "boolean"
                },
                "time_nanos": {
                    "type": "integer"
                },
                "time_seconds": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.StartPrivateConferenceRequest": {
            "type": "object",
            "properties": {
                "broadcast": {
                    "type": "boolean"
                },
                "chat": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "interest": {
                    "type": "string"
                },
                "participantlimit": {
                    "type": "integer"
                },
                "recording": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.StartPublicConferenceRequest": {
            "type": "object",
            "properties": {
                "broadcast": {
                    "type": "boolean"
                },
                "chat": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "interest": {
                    "type": "string"
                },
                "joinType": {
                    "type": "string"
                },
                "participantlimit": {
                    "type": "integer"
                },
                "recording": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "models.ValidName": {
            "type": "object",
            "properties": {
                "userName": {
                    "type": "string"
                }
            }
        },
        "pb.LoginResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "pb.OtpSignUpResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.ResendOtpResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "pb.SignupResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "pb.ValidNameResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}