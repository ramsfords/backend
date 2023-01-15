package models

type CapacityProviderAccountGroup struct {
	Code     string    `json:"code"`
	Accounts []Account `json:"accounts"`
}
