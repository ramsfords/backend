package models

type AuthResponsePayload struct {
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	UserName     string `json:"userName,omitempty"`
}
