package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) EchoGetAllBusiness(ctx echo.Context) error {
	res, err := business.GetBusinesses(ctx.Request().Context())
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (business Business) GetBusinesses(ctx context.Context) (*v1.Business, error) {
	res, err := business.services.Db.GetBusiness(ctx, "req.Id")
	if err != nil {
		logger.Error(err, "GetAllBusinesses : error in getting all businesses")
		return nil, errs.ErrStoreInternal
	}

	return res, nil
}
