package business_api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

type BasicInfo struct {
	Business v1.Business `json:"business"`
	Token    string      `json:"token"`
}

func (business Business) GetBasicInfo(ctx echo.Context) error {
	businessID := ctx.PathParam("businessId")
	if businessID == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	user, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)

	if user == nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	email := user.GetString("email")
	fmt.Println(email)
	users, err := business.services.GetStaffsForABusiness(ctx.Request().Context(), businessID)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	bis, err := business.services.GetBusiness(ctx.Request().Context(), businessID)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	bookings, err := business.services.GetAllBookingsByBusinessId(ctx.Request().Context(), businessID)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	quotes, err := business.services.GetAllQuotesByBusinessId(ctx.Request().Context(), businessID)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	homeReq := v1.BasicInfo{
		Business: bis,
		Users:    sanitizeUserToFrontEnd(users),
		Bookings: bookings,
		Quotes:   quotes,
		User:     getUserFromUsers(email, users),
		Valid:    true,
	}
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, homeReq)
}
func getUserFromUsers(email string, users []*v1.User) *v1.User {
	for _, j := range users {
		if j.Email == email {
			return j
		}
	}
	return &v1.User{}
}
func sanitizeUserToFrontEnd(user []*v1.User) []*v1.User {
	userData := []*v1.User{}
	for _, usr := range user {
		userData = append(userData, &v1.User{
			Email:    usr.Email,
			Name:     usr.Name,
			UserName: usr.UserName,
			TokenKey: usr.TokenKey,
			Avatar:   usr.Avatar,
		})
	}
	return userData
}
func getCurrentUser(users []v1.User, email string) *v1.User {
	for _, j := range users {
		if j.Email == email {
			return &j
		}
	}
	return nil
}
