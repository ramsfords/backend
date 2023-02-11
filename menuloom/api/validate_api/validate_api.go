package validate_api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	"github.com/ramsfords/backend/menuloom/services"
)

type Validate struct {
	services *services.Services
}

func New(echo *echo.Group, services *services.Services) {
	validate := &Validate{services}
	api := echo.Group("/validate")
	api.GET("/validate/:id", validate.getValid)
	api.PUT("/validate/:id", validate.updateValid)
}

func (api Validate) getValid(ctx echo.Context) error {
	id := "pk#" + ctx.PathParam("id")
	record, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)
	fmt.Println(record)
	res, err := api.services.Db.GetValidate(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api Validate) updateValid(ctx echo.Context) error {
	id := "pk#" + ctx.PathParam("id")
	var data bool
	if err := ctx.Bind(&data); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	record, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)
	fmt.Println(record)
	err := api.services.Db.UpdateValidate(ctx.Request().Context(), id, data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.NoContent(http.StatusOK)
}
