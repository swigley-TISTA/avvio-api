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
  "consumes": [
    "text/plain"
  ],
  "produces": [
    "text/plain"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The purpose of this application is to test go-swagger in a simple GET request.",
    "title": "Testing go-swagger generation",
    "contact": {
      "name": "Daniel",
      "email": "danielfs.ti@gmail.com"
    },
    "license": {
      "name": "MIT",
      "url": "http://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "paths": {
    "/hello/{name}": {
      "get": {
        "description": "Returns a simple Hello message",
        "consumes": [
          "text/plain"
        ],
        "produces": [
          "text/plain"
        ],
        "tags": [
          "hello"
        ],
        "operationId": "Hello",
        "parameters": [
          {
            "type": "string",
            "description": "Name to be returned.",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The hello message"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "text/plain"
  ],
  "produces": [
    "text/plain"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The purpose of this application is to test go-swagger in a simple GET request.",
    "title": "Testing go-swagger generation",
    "contact": {
      "name": "Daniel",
      "email": "danielfs.ti@gmail.com"
    },
    "license": {
      "name": "MIT",
      "url": "http://opensource.org/licenses/MIT"
    },
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "paths": {
    "/hello/{name}": {
      "get": {
        "description": "Returns a simple Hello message",
        "consumes": [
          "text/plain"
        ],
        "produces": [
          "text/plain"
        ],
        "tags": [
          "hello"
        ],
        "operationId": "Hello",
        "parameters": [
          {
            "type": "string",
            "description": "Name to be returned.",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The hello message"
          }
        }
      }
    }
  }
}`))
}