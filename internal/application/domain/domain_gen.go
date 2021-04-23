// Package types provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package domain

// UserAuth defines model for UserAuth.
type UserAuth struct {
	UserName   string `json:"UserName"`
	UserSecret string `json:"UserSecret"`
}

// UserAuthJSONBody defines parameters for UserAuth.
type UserAuthJSONBody UserAuth

// UserAuthJSONRequestBody defines body for UserAuth for application/json ContentType.
type UserAuthJSONRequestBody UserAuthJSONBody