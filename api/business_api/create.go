package business_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) EchoCreateBusiness(ctx echo.Context) error {
	req := model.Session{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	usr := &v1.User{}
	usrBytes, err := json.Marshal(req.UserMetadata)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = json.Unmarshal(usrBytes, usr)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	id := business.services.Db.GetBusinessCount()
	bis := &v1.Business{
		BusinessId:                        fmt.Sprintf("%d", id),
		AccountingEmail:                   req.Email,
		BusinessEmail:                     req.Email,
		CustomerServiceEmail:              req.Email,
		HighPriorityEmail:                 req.Email,
		AdminEmail:                        req.Email,
		CreatedAt:                         req.UserMetadata.Created,
		UpdatedAt:                         req.UserMetadata.Created,
		NeedsAddressUpdate:                true,
		NeedsDefaultPickupAddressUpdate:   true,
		NeedsDefaultDeliveryAddressUpdate: true,
		ReferredBy:                        "referral",
		AdminUser:                         usr,
	}
	res, err := business.CreateBusiness(ctx.Request().Context(), bis)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (business Business) CreateBusiness(ctx context.Context, req *v1.Business) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		logger.Error(err, "CreateBusiness Validate : req data validation failed")
		return nil, errs.ErrInputDataNotValid
	}
	err = business.services.Db.SaveBusiness(ctx, req, req.BusinessId)
	if err != nil {
		logger.Error(err, "CreateBusiness InsertBusiness : error in inserting business into the database")
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
