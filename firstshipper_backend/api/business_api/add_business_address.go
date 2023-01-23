package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) AddBusinessAddress(ctx echo.Context) error {
	address := &v1.Address{}
	err := ctx.Bind(address)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.PathParam("businessId")
	if businessId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	address.BusinessId = businessId
	err = address.Validate()
	if err != nil {
		business.services.Logger.Errorf("req data validation failed: %s", err)
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	address, err = business.services.AddLocationAddress(newContext, address.BusinessId, address)
	if err != nil {
		business.services.Logger.Errorf("adding address to database failded: %s", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, address)
}

// func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {

// }
