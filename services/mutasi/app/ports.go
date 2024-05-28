package app

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
)

type AccountDatastore interface {
	Mutasi(request models.Mutasi) error
}

type AccountApp interface {
	Mutasi(request models.Mutasi) (response helpers.Response, err error)
}
