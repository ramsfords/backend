package quote_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoUpdateQuote(ctx echo.Context) error {
	req := v1.QuoteRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := qt.UpdateQuote(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)

}

func (qt Quote) UpdateQuote(ctx context.Context, quoteReq *v1.QuoteRequest) (*v1.QuoteRequest, error) {
	err := quoteReq.Validate()
	if err != nil {
		logger.Error(err, "UpdateQuote Validate : req data validation failed")
		return nil, errs.ErrInputDataNotValid
	}
	err = qt.services.Db.UpdateQuote(ctx, quoteReq)
	if err != nil {
		logger.Error(err, "UpdateQuote : error in updating quote into the database")
		return nil, errs.ErrLocationUpdationFailed
	}

	return nil, nil
}
