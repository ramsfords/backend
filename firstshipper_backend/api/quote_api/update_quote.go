package quote_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
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
		qt.services.Logger.Error("UpdateQuote Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = qt.services.UpdateQuote(ctx, quoteReq)
	if err != nil {
		qt.services.Logger.Error("UpdateQuote : error in updating quote into the database: %s", err)
		return nil, errs.ErrLocationUpdationFailed
	}

	return nil, nil
}
