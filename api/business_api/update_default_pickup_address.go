package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) UpdateDefaultPickupAddress(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	address := &v1.Address{}
	//err := server.unMarshall(ctx, signUpReq)
	err = ctx.Bind(address)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = business.services.Db.UpdateLocation(ctx.Request().Context(), authContext.UserMetadata.OrganizationId, address)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusAccepted)
}
