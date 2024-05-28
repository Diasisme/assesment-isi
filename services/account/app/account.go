package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Diasisme/asssesment-march-ihsan.git/config/logging/utils"
	"github.com/Diasisme/asssesment-march-ihsan.git/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// Daftar implements AccountApp.
func (a *accountApp) Daftar(request models.Nasabah) (response helpers.Response, err error) {
	var tabungData models.Tabungan

	noRek := utils.GenerateRandomNumber(8)
	hashPin, _ := utils.HashPin(request.Pin)
	request.Pin = hashPin
	request.NomorRekening = strconv.Itoa(noRek)

	if err = a.accRepo.Daftar(request); err != nil {
		remark := "Gagal memasukkan data, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	getData, err := a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	tabungData = models.Tabungan{
		NomorRekening: request.NomorRekening,
		Nominal:       0,
		NasabahID:     getData.ID,
	}
	err = a.accRepo.BuatTabung(tabungData)
	if err != nil {
		remark := "Gagal membuat tabungan, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	response.Status = 200
	response.Message = "Berhasil memasukkan data"
	response.Data = request

	return
}

// Tabung implements AccountApp.
func (a *accountApp) Tabung(request models.Tabungan) (response helpers.Response, err error) {
	var transaksiData models.Transaksi
	// var dataAccount models.Nasabah

	getData, err := a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem, silahkan coba lagi.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err = a.accRepo.TambahTabung(request); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	transaksiData = models.Transaksi{
		NomorRekening: request.NomorRekening,
		Nominal:       request.Nominal,
		KodeTransaksi: "C",
		NasabahID:     getData.ID,
	}

	err = a.accRepo.Transaksi(transaksiData)
	if err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	response.Status = 200
	response.Message = "Berhasil melakukan transaksi"
	response.Data = request

	return
}

func (a *accountApp) Tarik(request models.Tabungan) (response helpers.Response, err error) {
	var transaksiData models.Transaksi
	var dataTabungan models.Tabungan

	_, err = a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	dataTabungan, err = a.accRepo.GetDataTabungan(request.NomorRekening)
	if err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Tabungan nasabah dengan nomor rekening %s tidak ditemukan", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if dataTabungan.Nominal < request.Nominal {
		remark := fmt.Sprintf("Saldo rekening tabungan %s tidak cukup", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		return
	}

	if err = a.accRepo.KurangTabung(request); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	transaksiData = models.Transaksi{
		NomorRekening: request.NomorRekening,
		Nominal:       request.Nominal,
		KodeTransaksi: "D",
	}

	if err = a.accRepo.Transaksi(transaksiData); err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = http.StatusBadRequest
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		response.Status = http.StatusBadRequest
		response.Message = remark
	}

	response.Status = 200
	response.Message = "Berhasil melakukan transaksi"
	response.Data = request

	return
}
