package api

import (
	v1 "github.com/go-playground/validator/v10"

	"github.com/Diasisme/asssesment-march-ihsan.git/pkg/logging"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/account/app"
)

type accountApi struct {
	app      app.AccountApp
	validate v1.Validate
	log      *logging.Logger
}

func InitApi(app app.AccountApp, log *logging.Logger) *accountApi {
	return &accountApi{
		app:      app,
		validate: *v1.New(),
		log:      log,
	}
}
