package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) EchoGetAllBusinesses(ctx echo.Context) error {
	res, err := business.GetBusiness(ctx.Request().Context(), "&req")
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusAccepted, res)

}

func (business Business) GetBusiness(ctx context.Context, req string) (*v1.Business, error) {
	businessData, err := business.services.Db.GetBusiness(ctx, req)
	if err != nil {
		logger.Error(err, "GetAllBusinesses : error in getting all businesses")
		return nil, errs.ErrStoreInternal
	}

	return businessData, nil
}
