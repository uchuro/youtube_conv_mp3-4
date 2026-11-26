package types

type User struct {
	UserEmail    string `json:"email"`
	UserPassword string `json:"password"`
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
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	EffectedData string `json:"effectedData"`
}

type SuccessResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	EffectedData string `json:"effectedData"`
}
