package authapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

type RedirectLoginData struct {
	Token          string `json:"token"`
	UserId         string `json:"userId"`
	Email          string `json:"email"`
	ValidUntil     string `json:"validUntil"`
	OrganizationId string `json:"organizationId"`
	Password       string `json:"password"`
}

func (auth AuthApi) ConfirmEmail(ctx echo.Context) error {
	token := ctx.QueryParam("token")
	if len(token) < 50 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	subStr := strings.Split(token, "next")
	userData := subStr[0]
	next := subStr[1]
	data, err := auth.services.Crypto.Decrypt(userData)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	validDate, err := auth.services.Crypto.Decrypt(next)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	dateData := &time.Time{}
	// json unmarshal data to user
	err = json.Unmarshal(validDate, dateData)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if time.Now().After(*dateData) {
		return ctx.NoContent(http.StatusBadRequest)
	}
	user := &v1.User{}
	// json unmarshal data to user
	err = json.Unmarshal(data, user)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// err = auth.services.CognitoClient.ConfirmEmail(ctx.Request().Context(), user.Email)
	// if err != nil {
	// 	return ctx.String(http.StatusBadRequest, err.Error())
	// }
	createdAt := time.Now().Format(time.RFC3339)
	business := &v1.Business{
		BusinessId:                        user.OrganizationId,
		AccountingEmail:                   user.Email,
		BusinessEmail:                     user.Email,
		CustomerServiceEmail:              user.Email,
		HighPriorityEmail:                 user.Email,
		AdminEmail:                        user.Email,
		CreatedAt:                         createdAt,
		UpdatedAt:                         createdAt,
		NeedsAddressUpdate:                true,
		NeedsDefaultPickupAddressUpdate:   true,
		NeedsDefaultDeliveryAddressUpdate: true,
		ReferredBy:                        "referral",
		AdminUser:                         user,
	}
	go func() {
		err = auth.services.Db.SaveBusiness(context.Background(), business, user.OrganizationId)
		if err != nil {
			logger.Error(err, "could not save business")
		}
		err = auth.services.Db.SaveUser(context.Background(), user, user.OrganizationId)
		if err != nil {
			logger.Error(err, "could not save user")
		}
	}()

	return ctx.String(http.StatusOK, "Email confirmed")
}
