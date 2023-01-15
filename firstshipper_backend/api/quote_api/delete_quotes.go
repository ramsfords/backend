package quote_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoDeleteAllQuotes(ctx echo.Context) error {
	req := v1.Ids{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := qt.DeleteQuotes(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (qt Quote) DeleteQuotes(ctx context.Context, req *v1.Ids) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		qt.services.Logger.Error("DeleteQuote Validate : req data validation failed: %s", err)
		return nil, errs.ErrInvalidInputs
	}
	for _, j := range req.Ids {
		err = qt.services.DeleteQuote(ctx, j.Id)
		if err != nil {
			qt.services.Logger.Error("DeleteAllQuotesByBusinessId : error in deleting quotes : %s", err)
			return nil, errs.ErrLocationCreationFailed
		}
	}
	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
