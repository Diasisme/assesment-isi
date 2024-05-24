package app

import (
	"fmt"
	"strconv"

	"github.com/Diasisme/asssesment-march-ihsan.git/pkg/logging/utils"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/models"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/payload"
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
		remark := "Gagal memasukkan data, , silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	tabungData = models.Tabungan{
		NomorRekening: request.NomorRekening,
		Nominal:       0,
	}
	if err = a.accRepo.BuatTabung(tabungData); err != nil {
		remark := "Gagal membuat tabungan, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		response.Status = 500
		response.Message = remark
	}

	response.Status = 200
	response.Message = "Berhasil memasukkan data"

	return
}

// Tabung implements AccountApp.
func (a *accountApp) Tabung(request models.Tabungan) (response helpers.Response, err error) {
	var transaksiData models.Transaksi
	// var dataAccount models.Nasabah

	_, err = a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem, silahkan coba lagi.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err = a.accRepo.TambahTabung(request); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	transaksiData = models.Transaksi{
		NomorRekening: request.NomorRekening,
		Nominal:       request.Nominal,
		KodeTransaksi: "C",
	}

	err = a.accRepo.Transaksi(transaksiData)
	if err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
	}

	response.Status = 200
	response.Message = "Berhasil melakukan transaksi"

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
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
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
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Tabungan nasabah dengan nomor rekening %s tidak ditemukan", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if dataTabungan.Nominal < request.Nominal {
		remark := fmt.Sprintf("Saldo rekening tabungan %s tidak cukup", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		return
	}

	if err = a.accRepo.KurangTabung(request); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
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
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		response.Status = 500
		response.Message = remark
	}

	response.Status = 200
	response.Message = "Berhasil melakukan transaksi"

	return
}

func (a *accountApp) Transfer(request payload.TransferReq) (response helpers.Response, err error) {
	var transaksiData models.Transaksi
	var dataTabungan models.Tabungan

	_, err = a.accRepo.GetDataAccount(request.NomorRekeningAsal)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nomor rekening nasabah %s tidak ditemukan", request.NomorRekeningAsal)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	_, err = a.accRepo.GetDataAccount(request.NomorRekeningTujuan)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}
	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nomor rekening nasabah %s tidak ditemukan", request.NomorRekeningTujuan)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	dataTabungan, err = a.accRepo.GetDataTabungan(request.NomorRekeningAsal)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Tabungan nasabah dengan nomor rekening %s tidak ditemukan", request.NomorRekeningTujuan)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
	}

	if dataTabungan.Nominal < request.Nominal {
		remark := fmt.Sprintf("Saldo rekening tabungan %s tidak cukup", request.NomorRekeningAsal)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		return
	}

	dataTabungan = models.Tabungan{
		NomorRekening: request.NomorRekeningAsal,
		Nominal:       request.Nominal,
	}

	if err = a.accRepo.KurangTabung(dataTabungan); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	dataTabungan = models.Tabungan{
		NomorRekening: request.NomorRekeningTujuan,
		Nominal:       request.Nominal,
	}

	if err = a.accRepo.TambahTabung(dataTabungan); err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	transaksiData = models.Transaksi{
		NomorRekening: request.NomorRekeningAsal,
		Nominal:       request.Nominal,
		KodeTransaksi: "T",
	}

	if err = a.accRepo.Transaksi(transaksiData); err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	transaksiData = models.Transaksi{
		NomorRekening: request.NomorRekeningTujuan,
		Nominal:       request.Nominal,
		KodeTransaksi: "T",
	}

	if err = a.accRepo.Transaksi(transaksiData); err != nil {
		remark := "Gagal mencatat transaksi, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
	}

	response.Message = "Berhasil melakukan transfer"
	response.Status = 200

	return

}

func (a *accountApp) GetSaldoTabungan(request payload.GetTransaksiReq) (response helpers.Response, err error) {
	var respData payload.GetSaldoTabunganResp

	_, err = a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	_, err = a.accRepo.GetDataTabungan(request.NomorRekening)
	if err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Tabungan nasabah dengan nomor rekening %s tidak ditemukan", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	result, err := a.accRepo.GetSaldoTabungan(request.NomorRekening)
	if err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	respData = payload.GetSaldoTabunganResp{
		NomorRekening: request.NomorRekening,
		Saldo:         result.Nominal,
	}

	response.Status = 200
	response.Message = "Data berhasil didapatkan"
	response.Data = respData

	return
}

func (a *accountApp) GetMutasi(request payload.GetTransaksiReq) (response helpers.Response, err error) {
	_, err = a.accRepo.GetDataAccount(request.NomorRekening)
	if err != nil && err != gorm.ErrRecordNotFound {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	if err == gorm.ErrRecordNotFound {
		remark := fmt.Sprintf("Nasabah dengan nomor %s tidak terdaftar di sistem.", request.NomorRekening)
		a.log.Error(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	result, err := a.accRepo.Mutasi(request.NomorRekening)
	if err != nil {
		remark := "Terdapat kesalahan pada database, silahkan coba lagi."
		a.log.Warn(logrus.Fields{
			"err": err,
		}, nil, remark)
		response.Status = 500
		response.Message = remark
		err = status.Error(codes.OK, err.Error())
		return
	}

	response.Status = 200
	response.Message = "Data berhasil didapatkan"
	response.Data = result

	return
}
