package app

import (
	"net/http"

	"github.com/Diasisme/asssesment-march-ihsan.git/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *accountApp) Mutasi(request models.Mutasi) (response helpers.Response, err error) {

	err = a.accRepo.Mutasi(request)
	if err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
	}

	response.Status = 200
	response.Message = "Berhasil melakukan transaksi"

	return
}
