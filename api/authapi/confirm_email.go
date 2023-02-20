package authapi

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
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
type LoginData struct {
	Token          string `json:"token"`
	UserId         string `json:"userId"`
	Email          string `json:"email"`
	ValidUntil     string `json:"validUntil"`
	OrganizationId string `json:"organizationId"`
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

	err = auth.services.CognitoClient.ConfirmEmail(ctx.Request().Context(), user.Email)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	createdAt := time.Now().Format(time.RFC3339)
	business := &v1.Business{
		BusinessId:                        user.OrgId,
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
		err = auth.services.Db.SaveBusiness(context.Background(), business, user.OrgId)
		if err != nil {
			auth.services.Logger.Errorf("could not save business: %v", err)
		}
		err = auth.services.Db.SaveUser(context.Background(), user, user.OrgId)
		if err != nil {
			auth.services.Logger.Errorf("could not save user: %v", err)
		}
	}()

	loginData := &RedirectLoginData{
		UserId:         user.Id,
		Email:          user.Email,
		Password:       user.Password,
		OrganizationId: user.OrgId,
		ValidUntil:     time.Now().Add(time.Hour * 1).Format(time.RFC3339),
	}
	redirectLoginDataInc, err := auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		return ctx.NoContent(http.StatusOK)
	}
	return ctx.String(http.StatusOK, redirectLoginDataInc)
}
