package business_api

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) EchoAddStaff(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	addStaffReq := &v1.AddStaff{}
	err = ctx.Bind(addStaffReq)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	addStaffReq.BusinessId = authContext.UserMetadata.OrganizationId
	res, err := business.AddStaff(ctx.Request().Context(), addStaffReq)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusCreated, res)
}
func (business Business) AddStaff(ctx context.Context, req *v1.AddStaff) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		logger.Error(err, "req data validation failed")
		return nil, errs.ErrInvalidInputs
	}
	isAdmin := false

	if !isAdmin {
		logger.Error(errors.New("user does not have admin role"), "user does not have admin role to add staff for email")
		return nil, errs.ErrNotAllowed
	}
	dbUser := utils.SanitizeUser(req)

	err = business.services.Db.SaveUser(ctx, dbUser, req.BusinessId)
	if err != nil {
		logger.Error(err, "adding staff to database failded")
		return nil, errs.ErrInvalidInputs
	}

	return &v1.Ok{
		Success: true,
		Message: "user is created please confirm your email",
		Code:    200,
	}, nil
}
