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
        "/offers": {
            "post": {
                "description": "Create a new offer with media files (uploaded as form-data)",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Create a new offer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Offer Title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Offer Description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Offer Type",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "date",
                        "description": "Offer Expiry Date",
                        "name": "expiry_date",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Commission Agent",
                        "name": "commission_agent",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Commission Agency",
                        "name": "commission_agency",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Discount",
                        "name": "discount",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Price Installment",
                        "name": "price_installment",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "Price Cash",
                        "name": "price_cash",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Media files",
                        "name": "media",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Offer created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/offers/active": {
            "get": {
                "description": "Retrieve a list of active offers that are approved and not expired",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Get active offers",
                "responses": {
                    "200": {
                        "description": "List of active offers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/teklif_internal_api_models.OfferResp"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "No offers found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/offers/attach-properties": {
            "post": {
                "description": "Attach properties to a specific offer by their IDs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Attach properties to an offer",
                "parameters": [
                    {
                        "description": "Request body for attaching properties",
                        "name": "offer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/teklif_internal_api_models.AttachOfferPropertiesReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Properties successfully attached",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/offers/expire-properties": {
            "post": {
                "description": "Expire all offers associated with the provided property IDs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Expire offers by property IDs",
                "parameters": [
                    {
                        "description": "Request body containing property IDs",
                        "name": "property_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/teklif_internal_api_models.PropertyIDsReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/offers/properties": {
            "post": {
                "description": "Retrieve a list of active offers based on property IDs that are approved and not expired",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Get active offers by property IDs",
                "parameters": [
                    {
                        "description": "List of property IDs",
                        "name": "property_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/teklif_internal_api_models.PropertyIDsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of active offers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/teklif_internal_api_models.OfferResp"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "No offers found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/offers/week-expired": {
            "get": {
                "description": "Retrieve a list of offers that are marked as \"Expired\" and expired in the last 7 days.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Get list of offers with \"Expired\" status that have expired in the last 7 days.",
                "responses": {
                    "200": {
                        "description": "List of expired offers",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/teklif_internal_api_models.OfferResp"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "No offers found",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/offers/{offerID}/approve": {
            "post": {
                "description": "Approve the offer with the specified ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Approve an offer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Offer ID",
                        "name": "offerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/offers/{offerID}/delete": {
            "delete": {
                "description": "Delete an offer by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offers"
                ],
                "summary": "Delete an offer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the offer to be deleted",
                        "name": "offer_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "teklif_internal_api_models.AttachOfferPropertiesReq": {
            "type": "object",
            "required": [
                "offer_id",
                "property_ids"
            ],
            "properties": {
                "offer_id": {
                    "description": "OfferID UUID\n@example \"123e4567-e89b-12d3-a456-426614174000\"",
                    "type": "string"
                },
                "property_ids": {
                    "description": "PropertyIDs list of UUIDs\n@example [\"123e4567-e89b-12d3-a456-426614174001\", \"123e4567-e89b-12d3-a456-426614174002\"]",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "teklif_internal_api_models.OfferResp": {
            "type": "object",
            "properties": {
                "commission_agency": {
                    "type": "number"
                },
                "commission_agent": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "discount": {
                    "type": "number"
                },
                "expiry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "media": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "price_cash": {
                    "type": "number"
                },
                "price_installment": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "teklif_internal_api_models.PropertyIDsReq": {
            "type": "object",
            "required": [
                "property_ids"
            ],
            "properties": {
                "property_ids": {
                    "description": "PropertyIDs list of UUIDs\n@example [\"123e4567-e89b-12d3-a456-426614174001\", \"123e4567-e89b-12d3-a456-426614174002\"]",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
