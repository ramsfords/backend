package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type AddBusinessName struct {
	BusinessName string `json:"businessName"`
}

func (business Business) UpdateBusinessName(ctx echo.Context) error {
	businessName := &AddBusinessName{}
	err := ctx.Bind(businessName)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.QueryParam("businessId")
	if businessId == "" || businessName.BusinessName == "" || len(businessName.BusinessName) < 3 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	err = business.services.UpdateBusinessName(newContext, businessId, businessName.BusinessName)
	if err != nil {
		business.services.Logger.Errorf("adding address to database failded: %s", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusCreated)
}

// func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {

// }
