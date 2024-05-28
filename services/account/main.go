package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Diasisme/asssesment-march-ihsan.git/api"
	"github.com/Diasisme/asssesment-march-ihsan.git/app"
	"github.com/Diasisme/asssesment-march-ihsan.git/config"
	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging"
	"github.com/Diasisme/asssesment-march-ihsan.git/datastore"
	"github.com/Diasisme/asssesment-march-ihsan.git/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	viper := config.NewViper()
	var err error

	varenv := config.NewEnvVar(viper)
	logger := logging.NewLogger(varenv.Service)

	redis, err := datastore.NewRedis(varenv)
	if err != nil {
		logger.Fatal(logrus.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		os.Exit(2)
	}
	ds := datastore.InitDB(varenv, logger, redis.Conn)
	appRoute := app.InitApp(ds, logger)
	apiRoute := api.InitApi(appRoute, logger)

	err = ds.MigrationDB()
	if err != nil {
		logger.Fatal(logrus.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		os.Exit(2)
	}
	

	e.POST("/daftar", apiRoute.Daftar)

	middleware := middleware.InitMiddleWare(*ds, logger)

	protected := e.Group("/v2")
	protected.Use(middleware.BasicAuthMiddleWare)
	protected.POST("/tabung", apiRoute.Tabung)
	protected.POST("/tarik", apiRoute.Tarik)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8100"))

	e.Start(fmt.Sprintf(":%s", varenv.ServicePort))
}
