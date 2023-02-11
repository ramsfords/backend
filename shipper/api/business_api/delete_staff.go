package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (business Business) DeleteStaff(ctx echo.Context) error {
	data := &struct {
		Token            string `json:"token"`
		RemoveStaffEmail string `json:"remove_staff_email"`
		BusinessID       string `json:"business_id"`
	}{}
	err := ctx.Bind(data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.PathParam("businessId")
	err = business.services.Db.DeleteStaff(ctx.Request().Context(), businessId, data.RemoveStaffEmail)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusAccepted)
}
