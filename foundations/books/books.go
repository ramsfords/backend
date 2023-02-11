package books

import (
	"github.com/ramsfords/backend/configs"
	zoho "github.com/ramsfords/backend/foundations/auth"
)

// API is used for interacting with the Zoho Books API
type API struct {
	*configs.Config
	*zoho.Zoho
	id string
}

// New returns a *books.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho, conf *configs.Config) *API {
	z.OrganizationID = conf.Zoho.ZohoOrgId
	return &API{
		Config: conf,
		Zoho:   z,
		id:     conf.Zoho.ZohoOrgId,
	}
}
