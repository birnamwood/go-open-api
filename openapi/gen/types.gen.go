// Package openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package openapi

// Error defines model for Error.
type Error struct {

	// Error code
	Code int `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewUser defines model for NewUser.
type NewUser struct {

	// email
	Email *string `json:"email,omitempty"`

	// pass
	Name string `json:"name"`

	// pass
	Password *string `json:"password,omitempty"`
}

// User defines model for User.
type User struct {

	// email
	Email *string `json:"email,omitempty"`

	// Unique id of the user
	Id int `json:"id"`

	// pass
	Name *string `json:"name,omitempty"`

	// pass
	Password *string `json:"password,omitempty"`
}

// Users defines model for Users.
type Users []User

// FindUsersParams defines parameters for FindUsers.
type FindUsersParams struct {

	// 件数制限
	Limit *int32 `json:"limit,omitempty"`
}

// AddUserJSONBody defines parameters for AddUser.
type AddUserJSONBody NewUser

// AddUserRequestBody defines body for AddUser for application/json ContentType.
type AddUserJSONRequestBody AddUserJSONBody

