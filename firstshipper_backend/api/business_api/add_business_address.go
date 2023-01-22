package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) AddBusinessAddress(ctx echo.Context) error {
	address := &v1.Address{}
	//err := server.unMarshall(ctx, signUpReq)
	err := ctx.Bind(address)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.PathParam("businessId")
	if businessId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	address.BusinessId = businessId
	res, err := business.AddAddress(ctx.Request().Context(), address)
	if err != nil {
		ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}
func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {
	err := address.Validate()
	if err != nil {
		business.services.Logger.Errorf("req data validation failed: %s", err)
		return nil, errs.ErrInvalidInputs
	}
	err = business.services.AddLocationAddress(ctx, address.BusinessId, address)
	if err != nil {
		business.services.Logger.Errorf("adding address to database failded: %s", err)

	}
	err = business.services.UpdateBusinessAddressUpdateNeeded(ctx, address.BusinessId)
	if err != nil {
		business.services.Logger.Errorf("updating businesss address to database failded: %s", err)

	}
	if err != nil {
		business.services.Logger.Errorf("updating AddAdress to database failded: %s", err)
		return nil, errs.ErrInvalidInputs
	}
	return &v1.Ok{
		Success: true,
		Message: "user is created please confirm your email",
		Code:    200,
	}, nil
}
