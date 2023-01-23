package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) GinUpdateStaffRole(ctx echo.Context) error {
	updateReq := &v1.UpdateUserRole{}
	//err := server.unMarshall(ctx, signUpReq)
	err := ctx.Bind(updateReq)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.PathParam("businessId")
	if businessId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := business.UpdateStaffRole(ctx.Request().Context(), updateReq, businessId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusAccepted, res)
}
func (business Business) UpdateStaffRole(ctx context.Context, req *v1.UpdateUserRole, businessId string) (*v1.Ok, error) {
	err := req.Validate()

	isAdmin := false

	if !isAdmin {
		business.services.Logger.Errorf("user trying to update role without valid role", req.Token)
		return &v1.Ok{}, errs.ErrNotAllowed
	}
	err = business.services.UpdateStaffRole(ctx, businessId, req.StaffEmail, req.NewRole)
	if err != nil {
		business.services.Logger.Info("could not update staff role ", err)
		return &v1.Ok{
			Success: false,
			Message: "please try again leter",
			Code:    403,
		}, errs.ErrInternal
	}
	return &v1.Ok{
		Success: true,
		Message: "user's role updated",
		Code:    200,
	}, nil
}