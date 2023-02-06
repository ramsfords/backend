package business_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
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
	email := ctx.QueryParam("email")

	// user, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)

	// if user == nil {
	// 	return ctx.NoContent(http.StatusBadRequest)
	// }
	// email := user.GetString("email")
	fmt.Println(email)
	data, err := business.services.GetAllDataByBusinessId(ctx.Request().Context(), businessID)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	shipments, err := business.services.GetAllBookingsByBusinessId(ctx.Request().Context(), businessID)
	if err != nil {
		business.services.Logger.Errorf("error getting shipments: %v", err)
	}
	removeSaltedPassword(data.Users)
	resdata := model.FrontEndBusinessData{
		Business:      data.Business,
		Users:         sanitizeUserToFrontEnd(data.Users),
		QuoteRequests: data.QuoteRequests,
		User:          &v1.FrontEndUser{},
		Shipments:     shipments,
	}
	resdata.User = getCurrentSanitizedUser(resdata.Users, email)
	return ctx.JSON(http.StatusOK, resdata)
}
func getUserFromUsers(email string, users []*v1.User) *v1.User {
	for _, j := range users {
		if j.Email == email {
			return j
		}
	}
	return &v1.User{}
}

func getCurrentUser(users []*v1.User, email string) *v1.User {
	for _, j := range users {
		if j.Email == email {
			return j
		}
	}
	return nil
}
func removeSaltedPassword(users []*v1.User) []*v1.User {
	for i, j := range users {
		j.Password = ""
		j.PasswordHash = ""
		users[i] = j
	}
	return users
}
func sanitizeUserToFrontEnd(users []*v1.User) []*v1.FrontEndUser {
	userData := []*v1.FrontEndUser{}
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(jsonBytes, &userData)
	if err != nil {
		return nil
	}
	return userData
}
func getCurrentSanitizedUser(users []*v1.FrontEndUser, email string) *v1.FrontEndUser {
	for _, j := range users {
		if j.Email == email {
			userData := &v1.FrontEndUser{}
			jsonBytes, err := json.Marshal(j)
			if err != nil {
				return nil
			}
			err = json.Unmarshal(jsonBytes, &userData)
			if err != nil {
				return nil
			}
			return userData
		}
	}
	return nil
}
