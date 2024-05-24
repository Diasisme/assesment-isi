package api

import (
	"net/http"
	"time"

	"github.com/Diasisme/asssesment-march-ihsan.git/services/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/models"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/payload"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f accountApi) Daftar(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.Nasabah
	var response helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	payloadValidator := new(payload.DaftarReq)

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	err = copier.Copy(&request, payloadValidator)
	if err != nil {
		remark := "Tidak dapat melakukan proses copy data, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error":            err,
			"source copy":      payloadValidator,
			"destination copy": request,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.Daftar(request)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f accountApi) Tabung(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.Tabungan
	var response helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	payloadValidator := new(payload.TabunganReq)

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind"
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	err = copier.Copy(&request, payloadValidator)
	if err != nil {
		remark := "Tidak dapat melakukan proses copy data, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error":            err,
			"source copy":      payloadValidator,
			"destination copy": request,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.Tabung(request)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f accountApi) Tarik(c echo.Context) (err error) {
	startTime := time.Now()
	var request models.Tabungan
	var response helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	payloadValidator := new(payload.TabunganReq)

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	err = copier.Copy(&request, payloadValidator)
	if err != nil {
		remark := "Tidak dapat melakukan proses copy data, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error":            err,
			"source copy":      payloadValidator,
			"destination copy": request,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.Tarik(request)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f accountApi) Transfer(c echo.Context) (err error) {
	startTime := time.Now()
	var response helpers.Response

	payloadValidator := new(payload.TransferReq)

	f.log.Info(logrus.Fields{
		"request": payloadValidator,
	}, payloadValidator, "log info")

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.Transfer(*payloadValidator)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     payloadValidator,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f accountApi) Saldo(c echo.Context) (err error) {
	startTime := time.Now()
	var response helpers.Response

	payloadValidator := new(payload.GetTransaksiReq)

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.GetSaldoTabungan(*payloadValidator)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     payloadValidator,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}

func (f accountApi) Mutasi(c echo.Context) (err error) {
	startTime := time.Now()
	var response helpers.Response

	payloadValidator := new(payload.GetTransaksiReq)

	f.log.Info(logrus.Fields{
		"request": payloadValidator,
	}, payloadValidator, "log info")

	if err = c.Bind(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses bind, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	if err = f.validate.Struct(payloadValidator); err != nil {
		remark := "Tidak dapat melakukan proses validasi struct, silahkan coba lagi."
		f.log.Error(logrus.Fields{
			"error": err,
			"data":  payloadValidator,
		}, nil, remark)
		response.Message = remark
		response.Status = 500
		response.Data = nil
		return err
	}

	result, err := f.app.GetMutasi(*payloadValidator)
	if err != nil {
		remark := result.Message
		f.log.Error(logrus.Fields{
			"error": err.Error(),
		}, nil, remark)

		response.Message = result.Message
		response.Status = result.Status
		response.Data = nil

		err = status.Error(codes.OK, err.Error())
		return
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     payloadValidator,
		"Result":      result,
		"error":       err,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return c.JSON(http.StatusOK, result)
}
