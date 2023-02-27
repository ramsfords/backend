package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) UpdateBusinessPhoneNumber(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	phoneNumber := &v1.PhoneNumber{}
	err = ctx.Bind(phoneNumber)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = phoneNumber.Validate()
	if err != nil {
		logger.Error(err, "req data validation failed")
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	phoneNumber, err = business.services.Db.AddPhoneNumber(newContext, authContext.UserMetadata.OrganizationId, phoneNumber)
	if err != nil {
		logger.Error(err, "adding address to database failded")
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, phoneNumber)
}

// func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {

// }
