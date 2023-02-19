package model

type UserSession struct {
	UserId         string   `json:"userId"`
	OrganizationId string   `json:"organizationId"`
	ExpiresIn      string   `json:"expiresIn"`
	Token          string   `json:"token"`
	Role           []string `json:"role"`
}
