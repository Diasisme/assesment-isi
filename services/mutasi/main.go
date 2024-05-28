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
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	viper := config.NewViper()
	e := echo.New()

	varenv := config.NewEnvVar(viper)
	logger := logging.NewLogger(varenv.Service)
	redis, err := datastore.NewRedis(varenv)
	if err != nil {
		logger.Fatal(logrus.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		os.Exit(2)
	}

	ds := datastore.InitDB(varenv, logger)

	err = ds.MigrationDB()
	if err != nil {
		logger.Fatal(logrus.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		os.Exit(2)
	}

	appRoute := app.InitApp(ds, logger)
	apiRoute := api.InitApi(appRoute, logger, redis.Conn)

	apiRoute.Mutasi()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8101"))

	e.Start(fmt.Sprintf(":%s", varenv.ServicePort))
}
