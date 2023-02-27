package utils

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/services"
)

type Ok struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Body       interface{} `json:"body"`
}

func Respond(provider services.Services, ctx echo.Context, result Ok, err error) {
	if err == nil {
		if !result.Success {
			result.Success = true
		}
		if result.StatusCode == 0 {
			result.StatusCode = http.StatusOK
		}
		if result.Message == "" {
			result.Message = "ok"
		}
		ctx.JSON(result.StatusCode, result)
		return
	}
	if err != nil {
		newErr, ok := err.(errs.ApiErr)
		if result.StatusCode == 302 {
			ctx.NoContent(result.StatusCode)
			return
		}
		if !ok {
			ctx.NoContent(http.StatusInternalServerError)
			return
		} else {
			if newErr.Cod > 500 {
				ctx.NoContent(http.StatusInternalServerError)
				return
			} else {
				ctx.NoContent(newErr.Cod)
				return
			}

		}
	}
}
