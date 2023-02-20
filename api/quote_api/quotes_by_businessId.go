package quote_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/utils"
	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/foundations/errs"
)

func (qt Quote) EchoGetQuotesByBusinessId(ctx echo.Context) error {
	authContext, err := utils.GetAuthContext(ctx)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}
	res, err := qt.GetQuotesByBusinessId(ctx.Request().Context(), authContext.OrganizationId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (qt Quote) GetQuotesByBusinessId(ctx context.Context, businessId string) ([]*model.QuoteRequest, error) {
	qts, err := qt.services.Db.GetAllQuotesByBusinessId(ctx, businessId)
	if err != nil {
		return nil, errs.ErrInvalidInputs
	}
	return qts, nil
}
