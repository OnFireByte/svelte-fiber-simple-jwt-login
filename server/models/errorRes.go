package models

const (
	ErrorInvalidEmail     string = "invalid_email"
	ErrorUsernameExists   string = "username_already_exists"
	ErrorEmailExists      string = "email_already_exists"
	ErrorPasswordNotMatch string = "password_not_match"
	ErrorUsernameNotFound string = "username_not_found"
)

type ErrorResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}