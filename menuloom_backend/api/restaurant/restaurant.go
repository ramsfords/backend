package restaurant_api

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/ramsfords/backend/menuloom_backend/services"
)

type restaurantApi struct {
	services services.Services
}

func New(echo *echo.Echo, services services.Services) {
	grp := echo.Group("/restaurant")
	grp.Use(apis.RequireAdminOrRecordAuth())
	menuHandler := restaurantApi{services: services}
	grp.POST("", menuHandler.createRestaurant)
	grp.GET("/:id", menuHandler.get)
	grp.PUT("", menuHandler.update)
	grp.DELETE("", menuHandler.delete)

}
