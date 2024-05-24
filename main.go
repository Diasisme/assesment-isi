package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Diasisme/asssesment-march-ihsan.git/pkg/logging"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/account/api"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/account/app"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/account/datastore"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatalf("err loading: %v", err)
	}

	SERVICE := os.Getenv("CONTAINER_ID_NAME")
	DB_HOST := os.Getenv("POSTGRES_HOST")
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_PORT := os.Getenv("POSTGRES_DB_PORT")
	DB_DATABASE := os.Getenv("POSTGRES_DB")
	SVC_PORT := os.Getenv("SVC_PORT")

	logger := logging.NewLogger(SERVICE)
	ds := datastore.InitDB(DB_HOST, DB_USER, DB_PASSWORD, DB_DATABASE, DB_PORT, logger)
	app := app.InitApp(ds, logger)
	apiRoute := api.InitApi(app, logger)

	ds.DB.AutoMigrate(&models.Tabungan{}, &models.Transaksi{})
	ds.DB.AutoMigrate(&models.Nasabah{})

	e.POST("/daftar", apiRoute.Daftar)
	e.POST("/tabung", apiRoute.Tabung)
	e.POST("/tarik", apiRoute.Tarik)
	e.POST("/transfer", apiRoute.Transfer)
	e.GET("/saldo/:nomor_rekening", apiRoute.Saldo)
	e.GET("/mutasi/:nomor_rekening", apiRoute.Mutasi)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8100"))

	e.Start(fmt.Sprintf(":%s", SVC_PORT))
}
