package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) UpdateBusinessPhoneNumber(ctx echo.Context) error {
	phoneNumber := &v1.PhoneNumber{}
	err := ctx.Bind(phoneNumber)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.PathParam("businessId")
	if businessId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = phoneNumber.Validate()
	if err != nil {
		business.services.Logger.Errorf("req data validation failed: %s", err)
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	phoneNumber, err = business.services.AddPhoneNumber(newContext, businessId, phoneNumber)
	if err != nil {
		business.services.Logger.Errorf("adding address to database failded: %s", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, phoneNumber)
}

// func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {

// }
