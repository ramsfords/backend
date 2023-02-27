package authapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (authApi AuthApi) EchoCreateBusiness(ctx echo.Context) error {
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
	id := authApi.services.Db.GetBusinessCount()
	idStr := fmt.Sprintf("%d", id)
	bis := &v1.Business{
		BusinessId:                        idStr,
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
	err = authApi.services.Db.SaveBusiness(ctx.Request().Context(), bis, idStr)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	err = authApi.services.Db.SaveUser(ctx.Request().Context(), usr, idStr)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, bis)
}
