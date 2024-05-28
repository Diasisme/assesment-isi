package middleware

import (
	"net/http"

	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging"
	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging/utils"
	"github.com/Diasisme/asssesment-march-ihsan.git/datastore"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MiddlewareApp struct {
	database *datastore.DatabaseData
	log      *logging.Logger
}

func (m *MiddlewareApp) BasicAuthMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nomor_rekening := c.Request().Header.Get("X-Nomor_rekening")
		pin := c.Request().Header.Get("X-Pin")
		db := m.database.DB
		if nomor_rekening == "" || pin == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Username and password required"})
		}

		m.log.Info(logrus.Fields{
			"nomor_rekening": nomor_rekening,
		}, nil, "")

		var data models.Nasabah
		if err := db.Where("nomor_rekening", nomor_rekening).First(&data).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				m.log.Error(logrus.Fields{"err": err}, nil, "error: Invalid username or password")
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Database error"})
		}

		if !utils.CheckPinHash(data.Pin, pin) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
		}

		c.Set("nomor_rekening:", data.NomorRekening)
		return next(c)
	}
}

func InitMiddleWare(db datastore.DatabaseData, logger *logging.Logger) *MiddlewareApp {
	return &MiddlewareApp{
		database: &db,
		log:      logger,
	}
}
