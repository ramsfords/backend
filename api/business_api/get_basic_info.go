package business_api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

type BasicInfo struct {
	Business v1.Business `json:"business"`
	Token    string      `json:"token"`
}

func (business Business) GetBasicInfo(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	data, err := business.services.Db.GetAllDataByBusinessId(ctx.Request().Context(), authContext.UserMetadata.OrganizationId)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	shipments, err := business.services.Db.GetAllBookingsByBusinessId(ctx.Request().Context(), authContext.UserMetadata.OrganizationId)
	if err != nil {
		logger.Error(err, "error getting shipments")
	}
	removeSaltedPassword(data.Users)
	resdata := model.FrontEndBusinessData{
		Business:      data.Business,
		Users:         sanitizeUserToFrontEnd(data.Users),
		QuoteRequests: data.QuoteRequests,
		User:          &v1.FrontEndUser{},
		Shipments:     shipments,
	}
	resdata.User = getCurrentSanitizedUser(resdata.Users, authContext.Email)
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
