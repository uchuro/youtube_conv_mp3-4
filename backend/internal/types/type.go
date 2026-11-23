package types

type User struct {
	UserEmail    string
	UserPassword string
}

/*
1. What should the backend send?
Generally, your API response can contain these fields:
Field	Purpose
success	Whether the operation succeeded
message	Human-readable result
data	Requested user or application data
token	Authentication token, when needed
error	Error information, when appropriate
*/

type ErrorResponse struct {
	Success      bool
	Message      string
	EffectedData string
}

type SuccessResponse struct {
	Message      string
	EffectedData string
}
