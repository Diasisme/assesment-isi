package api

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/app"
	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging"
	v1 "github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
)

type accountApi struct {
	app      app.AccountApp
	validate v1.Validate
	log      *logging.Logger
	redis    *redis.Client
}

func InitApi(app app.AccountApp, log *logging.Logger, rdb *redis.Client) *accountApi {
	return &accountApi{
		app:      app,
		validate: *v1.New(),
		log:      log,
		redis:    rdb,
	}
}
