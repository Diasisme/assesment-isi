package datastore

import (
	"fmt"

	"github.com/Diasisme/asssesment-march-ihsan.git/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseData struct {
	DB  *gorm.DB
	log *logging.Logger
}

func InitDB(host, user, password, database, port string, log *logging.Logger) *databaseData {
	var err error
	dsn := fmt.Sprintf("host=assessment-march-db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", user, password, database, port)
	print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &databaseData{
		DB:  db,
		log: log,
	}
}
