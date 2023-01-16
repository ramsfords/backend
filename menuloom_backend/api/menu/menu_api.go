package menu_api

import (
	"context"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/menuloom_backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Menu interface {
	createItem(ctx echo.Context) error
	createItems(ctx echo.Context) error
	createCategory(ctx echo.Context) error
	createCategories(ctx echo.Context) error
	getItem(ctx echo.Context) error
	getItems(ctx echo.Context) error
	getCategory(ctx echo.Context) error
	getCategories(ctx echo.Context) error
	getMenu(ctx echo.Context) error
	updateItem(ctx echo.Context) error
	updateItems(ctx echo.Context) error
	updateCategory(ctx echo.Context) error
	updateCategories(ctx echo.Context) error
	deleteItem(ctx echo.Context) error
	deleteItems(ctx echo.Context) error
	deleteCategory(ctx echo.Context) error
	deleteCategories(ctx echo.Context) error
	CreateCategories(ctx context.Context, data *v1.Categories) (*v1.ItemResponse, error)
}

type menuApi struct {
	services services.Services
}

func New(echo *echo.Group, services services.Services) {
	grp := echo.Group("/menu")
	var menuHandler Menu = menuApi{services: services}
	grp.POST("/item", menuHandler.createItem)
	grp.POST("/items", menuHandler.createItems)
	grp.POST("/category", menuHandler.createCategory)
	grp.POST("/categories", menuHandler.createItem)
	grp.GET("/item/:id", menuHandler.getItem)
	grp.GET("/items", menuHandler.getItems)
	grp.GET("/:id", menuHandler.getMenu)
	grp.GET("/category/:id", menuHandler.getCategory)
	grp.GET("/categories/:id", menuHandler.getCategories)
	grp.PATCH("/item", menuHandler.updateItem)
	grp.PATCH("/items", menuHandler.updateItems)
	grp.PATCH("/category", menuHandler.updateCategory)
	grp.PATCH("/categories", menuHandler.updateCategories)
	grp.DELETE("/item/:id", menuHandler.deleteItem)
	grp.DELETE("/items", menuHandler.deleteItems)
	grp.DELETE("/category/:id", menuHandler.deleteCategory)
	grp.DELETE("/categories", menuHandler.deleteCategories)
}
