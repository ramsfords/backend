package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
)

type AddBusinessName struct {
	BusinessName string `json:"businessName"`
}

func (business Business) UpdateBusinessName(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	businessName := &AddBusinessName{}
	err = ctx.Bind(businessName)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if businessName.BusinessName == "" || len(businessName.BusinessName) < 3 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	err = business.services.Db.UpdateBusinessName(newContext, authContext.UserMetadata.OrganizationId, businessName.BusinessName)
	if err != nil {
		logger.Error(err, "adding address to database failded")
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusCreated)
}

// func (business Business) AddAddress(ctx context.Context, address *v1.Address) (*v1.Ok, error) {

// }
