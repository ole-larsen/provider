// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Service Provider API",
    "version": "1.0.0"
  },
  "host": "provider-service:3335",
  "basePath": "/api/v1",
  "paths": {
    "/authorize": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "access type",
            "name": "access_type",
            "in": "query"
          },
          {
            "type": "string",
            "description": "code challenge",
            "name": "code_challenge",
            "in": "query"
          },
          {
            "type": "string",
            "description": "code challenge method",
            "name": "code_challenge_method",
            "in": "query"
          },
          {
            "type": "string",
            "description": "redirect uri",
            "name": "redirect_uri",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "code",
              "token"
            ],
            "type": "string",
            "description": "response type",
            "name": "response_type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "scope",
            "name": "scope",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "authorize",
            "schema": {
              "$ref": "#/definitions/authorize"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/credentials": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "domain for app",
            "name": "domain",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "get client id and client secret",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "instruments"
        ],
        "summary": "Prometheus metrics",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/metrics"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint always responds with \"pong\". When used over HTTP this API endpoint could also be used to measure network delays between your software and the API metrics.",
        "responses": {
          "200": {
            "description": "ping response",
            "schema": {
              "$ref": "#/definitions/ping"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/token": {
      "post": {
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "client secret",
            "name": "client_secret",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "get credentials",
            "name": "scope",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "get domain for credentials",
            "name": "domain",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "grant type credentials",
            "name": "grant_type",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "get client id and client secret",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/validate": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "access token",
            "name": "access_token",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "authorize",
            "schema": {
              "$ref": "#/definitions/validate"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "authorize": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string"
        }
      }
    },
    "credentials": {
      "type": "object",
      "required": [
        "client_id",
        "client_secret",
        "domain"
      ],
      "properties": {
        "client_id": {
          "type": "string"
        },
        "client_secret": {
          "type": "string"
        },
        "domain": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message",
        "code",
        "error"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "error": {
          "type": "boolean"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "metrics": {
      "type": "object",
      "additionalProperties": {
        "type": "string",
        "format": "number"
      }
    },
    "ok": {
      "type": "object",
      "required": [
        "result"
      ],
      "properties": {
        "result": {
          "description": "Result of method execution. ` + "`" + `ok` + "`" + ` in case of success",
          "type": "string",
          "enum": [
            "ok"
          ]
        }
      }
    },
    "ping": {
      "type": "object",
      "required": [
        "ping"
      ],
      "properties": {
        "ping": {
          "description": "Result of method execution. ` + "`" + `pong` + "`" + ` in case of success",
          "type": "string",
          "enum": [
            "pong"
          ]
        }
      }
    },
    "token": {
      "type": "object",
      "required": [
        "access_token",
        "expires_in",
        "scope",
        "token_type"
      ],
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "number"
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        }
      }
    },
    "validate": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API V1 request",
      "name": "v1"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Service Provider API",
    "version": "1.0.0"
  },
  "host": "provider-service:3335",
  "basePath": "/api/v1",
  "paths": {
    "/authorize": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "access type",
            "name": "access_type",
            "in": "query"
          },
          {
            "type": "string",
            "description": "code challenge",
            "name": "code_challenge",
            "in": "query"
          },
          {
            "type": "string",
            "description": "code challenge method",
            "name": "code_challenge_method",
            "in": "query"
          },
          {
            "type": "string",
            "description": "redirect uri",
            "name": "redirect_uri",
            "in": "query",
            "required": true
          },
          {
            "enum": [
              "code",
              "token"
            ],
            "type": "string",
            "description": "response type",
            "name": "response_type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "scope",
            "name": "scope",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "state",
            "name": "state",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "authorize",
            "schema": {
              "$ref": "#/definitions/authorize"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/credentials": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "domain for app",
            "name": "domain",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "get client id and client secret",
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "instruments"
        ],
        "summary": "Prometheus metrics",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/metrics"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint always responds with \"pong\". When used over HTTP this API endpoint could also be used to measure network delays between your software and the API metrics.",
        "responses": {
          "200": {
            "description": "ping response",
            "schema": {
              "$ref": "#/definitions/ping"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/token": {
      "post": {
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "client id",
            "name": "client_id",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "client secret",
            "name": "client_secret",
            "in": "formData",
            "required": true
          },
          {
            "type": "string",
            "description": "get credentials",
            "name": "scope",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "get domain for credentials",
            "name": "domain",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "grant type credentials",
            "name": "grant_type",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "get client id and client secret",
            "schema": {
              "$ref": "#/definitions/token"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/validate": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "public"
        ],
        "summary": "This API endpoint create, store and returns credentials for new user",
        "parameters": [
          {
            "type": "string",
            "description": "access token",
            "name": "access_token",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "authorize",
            "schema": {
              "$ref": "#/definitions/validate"
            }
          },
          "500": {
            "description": "When some error occurs",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "authorize": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string"
        }
      }
    },
    "credentials": {
      "type": "object",
      "required": [
        "client_id",
        "client_secret",
        "domain"
      ],
      "properties": {
        "client_id": {
          "type": "string"
        },
        "client_secret": {
          "type": "string"
        },
        "domain": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message",
        "code",
        "error"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "error": {
          "type": "boolean"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "metrics": {
      "type": "object",
      "additionalProperties": {
        "type": "string",
        "format": "number"
      }
    },
    "ok": {
      "type": "object",
      "required": [
        "result"
      ],
      "properties": {
        "result": {
          "description": "Result of method execution. ` + "`" + `ok` + "`" + ` in case of success",
          "type": "string",
          "enum": [
            "ok"
          ]
        }
      }
    },
    "ping": {
      "type": "object",
      "required": [
        "ping"
      ],
      "properties": {
        "ping": {
          "description": "Result of method execution. ` + "`" + `pong` + "`" + ` in case of success",
          "type": "string",
          "enum": [
            "pong"
          ]
        }
      }
    },
    "token": {
      "type": "object",
      "required": [
        "access_token",
        "expires_in",
        "scope",
        "token_type"
      ],
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "number"
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        }
      }
    },
    "validate": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API V1 request",
      "name": "v1"
    }
  ]
}`))
}
