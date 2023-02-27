package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/utils"
)

func (business Business) DeleteStaff(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	data := &struct {
		Token            string `json:"token"`
		RemoveStaffEmail string `json:"remove_staff_email"`
		BusinessID       string `json:"business_id"`
	}{}
	err = ctx.Bind(data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = business.services.Db.DeleteStaff(ctx.Request().Context(), authContext.UserMetadata.OrganizationId, data.RemoveStaffEmail)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusAccepted)
}
