package quote_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoDeleteQuote(ctx echo.Context) error {
	req := v1.Id{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := qt.DeleteQuote(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (qt Quote) DeleteQuote(ctx context.Context, req *v1.Id) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		qt.services.Logger.Error("DeleteQuote Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = qt.services.DeleteQuote(ctx, req.Id)
	if err != nil {
		qt.services.Logger.Error("DeleteQuoteByQuoteId : error in deleting quote : %s", err)
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
