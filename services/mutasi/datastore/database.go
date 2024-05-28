package datastore

import (
	"fmt"

	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseData struct {
	DB  *gorm.DB
	log *logging.Logger
}

func InitDB(varenv models.VarEnviroment, log *logging.Logger) *DatabaseData {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", varenv.Host, varenv.User, varenv.Pass, varenv.DB, varenv.Port)
	print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DatabaseData{
		DB:  db,
		log: log,
	}
}
