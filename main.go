package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/ramsfords/backend/api"
	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/services"
)

func Runner(services *services.Services, echoRouter *echo.Echo) {
	go func() error {
		err := services.Rapid.Login(&models.AuthRequestPayload{
			Username: services.Conf.SitesSettings.FirstShipper.RapidShipLTL.UserName,
			Password: services.Conf.SitesSettings.FirstShipper.RapidShipLTL.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}()

	// OR send a completely different email template
	// app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(services))
	// // OR send a completely different email template
	// app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(services.Conf))
	// this sets auth token to cloudflare KV on every login

}

func main() {
	conf := configs.GetConfig()
	servicesInstance := services.New(conf)
	// app := pocketbase.New()
	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	// grp.Use(apis.RequireAdminOrRecordAuth())
	// 	e.Router.GET("/ping", func(ctx echo.Context) error {
	// 		return ctx.JSON(200, echo.Map{
	// 			"message": "pong",
	// 			"status":  "ok",
	// 			"code":    200,
	// 		})
	// 	})
	// 	e.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 		AllowOrigins: []string{"https://localhost:3000", "https://firstshipper.com", "https://www.firstshipper.com", "https://menuloom.com", "https://localhost:3001", "https://127.0.0.1:3000", "https://127.0.0.1:8787", "https://api.firstshipper.com"},
	// 		AllowHeaders: []string{
	// 			echo.HeaderOrigin,
	// 			echo.HeaderContentType,
	// 			echo.HeaderAccept,
	// 			echo.HeaderAuthorization,
	// 			"first-access-token",
	// 			"first-refresh-token",
	// 			echo.HeaderAccessControlAllowHeaders,
	// 			echo.HeaderAccessControlRequestHeaders,
	// 			echo.HeaderAccessControlAllowOrigin,
	// 			echo.HeaderAccessControlAllowCredentials,
	// 		},
	// 		AllowMethods: []string{
	// 			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead, http.MethodPatch},
	// 		AllowCredentials: true,
	// 	}))
	// 	// e.Router.OPTIONS("/*", func(c echo.Context) error {
	// 	// 	orgin := c.Request().Header.Get("Origin")
	// 	// 	fmt.Println(orgin)
	// 	// 	c.Request().Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	// 	c.Request().Header.Set("Access-Control-Allow-Credentials", "true")
	// 	// 	c.Request().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	// 	c.Request().Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 	// 	return c.NoContent(http.StatusOK)
	// 	// })

	//
	// 	return nil

	// })
	echos := echo.New()

	echos.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://localhost:3000", "https://firstshipper.com", "https://www.firstshipper.com", "https://menuloom.com", "https://localhost:3001", "https://127.0.0.1:3000", "https://127.0.0.1:8787", "https://api.firstshipper.com"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"first-access-token",
			"first-refresh-token",
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlRequestHeaders,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials,
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead, http.MethodPatch},
		AllowCredentials: true,
	}))
	echos.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "pong",
			"status":  "ok",
			"code":    200,
		})
	})
	api.SetUpAPi(echos, servicesInstance)
	if err := echos.Start(":8090"); err != nil {
		log.Fatal(err)
	}
}

// func Process(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Request().Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Request().Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Request().Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Request().Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		c.Request().Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		if err := next(c); err != nil {
// 			fmt.Println(err)
// 		}
// 		status := strconv.Itoa(c.Response().Status)
// 		fmt.Println(status)
// 		return nil
// 	}
// }
