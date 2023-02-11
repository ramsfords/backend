package models

type AuthRequestPayload struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
