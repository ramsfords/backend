package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
)

func (business Business) AllowBooking(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	allowBooking := false
	err = ctx.Bind(&allowBooking)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	_, err = business.services.Db.UpdateAllowBooking(newContext, authContext.UserMetadata.OrganizationId, allowBooking)
	if err != nil {
		logger.Error(err, "adding address to database failded")
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, allowBooking)
}
