package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoGetAllQuotes(ctx echo.Context) error {
	req := v1.Empty{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := qt.GetAllQuotes(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (qt Quote) GetAllQuotes(ctx context.Context, qtReq *v1.Empty) (*v1.QuotesResponse, error) {
	err := qtReq.Validate()
	if err != nil {
		return nil, errs.ErrInvalidInputs
	}
	//qts, err := qt.Database.GetAllQuotesByBusinessId(ctx, qtReq.BusinessId)
	if err != nil {
		return nil, errs.ErrInvalidInputs
	}
	fmt.Println("qts")
	return &v1.QuotesResponse{}, nil
}
