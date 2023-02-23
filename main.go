package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/ramsfords/backend/api"
	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"golang.org/x/time/rate"
)

const serviceName = "firstshipper-api"

var Tracer *trace.TracerProvider

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
}

func main() {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := newExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	Tracer = trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := Tracer.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(Tracer)
	conf := configs.GetConfig()
	servicesInstance := services.New(conf)
	echos := echo.New()
	RegisterMiddleware(echos, servicesInstance)
	api.SetUpAPi(echos, servicesInstance)
	go func() error {
		err := servicesInstance.Rapid.Login(&models.AuthRequestPayload{
			Username: servicesInstance.Conf.SitesSettings.FirstShipper.RapidShipLTL.UserName,
			Password: servicesInstance.Conf.SitesSettings.FirstShipper.RapidShipLTL.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}()
	err = echos.Start("0.0.0.0:8090")
	if err != nil {
		log.Fatal(err)
	}
}
func RegisterMiddleware(echos *echo.Echo, services *services.Services) {
	var reqLimit rate.Limit = 20
	if services.Conf.Env == "dev" {
		reqLimit = 1000
	}
	echos.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	echos.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(reqLimit)))
	echos.Use(middleware.Recover())
	echos.Use(mid.CORS())
	echos.Use(mid.Tracer(services.Logger.Application))
	echos.Use(middleware.RequestID())
	echos.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			services.Logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}
