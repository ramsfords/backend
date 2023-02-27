package model

import (
	"time"

	v1 "github.com/ramsfords/types_gen/v1"
)

type LoginData struct {
	Token          string `json:"token"`
	UserId         string `json:"userId"`
	Email          string `json:"email"`
	ValidUntil     string `json:"validUntil"`
	OrganizationId string `json:"organizationId"`
}
type AppMetadata struct {
	Provider  string   `json:"provider"`
	Providers []string `json:"providers"`
}
type UserMetaData struct {
	Avatar                 string      `json:"avatar"`
	Created                string      `json:"created"`
	Email                  string      `json:"email"`
	EmailVisibility        bool        `json:"emailVisibility"`
	ID                     string      `json:"id"`
	LastResetSentAt        string      `json:"lastResetSentAt"`
	LastVerificationSentAt string      `json:"lastVerificationSentAt"`
	Name                   string      `json:"name"`
	Origin                 string      `json:"origin"`
	PasswordHash           string      `json:"passwordHash"`
	Token                  string      `json:"token"`
	TokenKey               string      `json:"tokenKey"`
	Type                   string      `json:"type"`
	Updated                string      `json:"updated"`
	UserName               string      `json:"userName"`
	Verified               bool        `json:"verified"`
	OrganizationId         string      `json:"organizationId"`
	Business               v1.Business `json:"business"`
}
type Identitity struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	IdentityData struct {
		Email string `json:"email"`
		Sub   string `json:"sub"`
	} `json:"identity_data"`
	Provider     string    `json:"provider"`
	LastSignInAt time.Time `json:"last_sign_in_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type User struct {
	ID                 string       `json:"id"`
	Aud                string       `json:"aud"`
	Role               string       `json:"role"`
	Email              string       `json:"email"`
	EmailConfirmedAt   time.Time    `json:"email_confirmed_at"`
	Phone              string       `json:"phone"`
	ConfirmationSentAt time.Time    `json:"confirmation_sent_at"`
	ConfirmedAt        time.Time    `json:"confirmed_at"`
	LastSignInAt       time.Time    `json:"last_sign_in_at"`
	AppMetadata        AppMetadata  `json:"app_metadata"`
	UserMetadata       UserMetaData `json:"user_metadata"`
	Identities         []Identitity `json:"identities"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}
type Session struct {
	AccessToken  string  `json:"access_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    float32 `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	User         `json:"user"`
	ExpiresAt    int64 `json:"expires_at"`
}
