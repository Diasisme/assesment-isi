package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Diasisme/asssesment-march-ihsan.git/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (f accountApi) Mutasi() error {
	startTime := time.Now()
	var request models.Mutasi
	var response helpers.Response
	var result helpers.Response

	f.log.Info(logrus.Fields{
		"request": request,
	}, request, "log info")

	for {
		entries, err := f.redis.XRead(&redis.XReadArgs{
			Streams: []string{"transaction", "0-0"},
			Count:   2,
			Block:   0,
		}).Result()
		if err != nil {
			remark := "Cannot read entries"
			f.log.Fatal(logrus.Fields{
				"error": err,
			}, nil, remark)
			response.Message = remark
			response.Status = http.StatusBadRequest
			response.Data = nil
			return err
		}

		f.log.Info(logrus.Fields{"entries": entries}, nil, "entries data")

		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values

			data := request

			data.NomorRekening = values["account_number"].(string)

			valueAmount := values["amount"].(string)
			amount, _ := strconv.ParseFloat(valueAmount, 64)
			data.Nominal = amount

			valueDate := values["date"].(string)
			date, _ := time.Parse("2006-01-02 15:04:05", valueDate)
			data.TanggalTransaksi = date

			data.JenisTransaksi = values["jenis_transaksi"].(string)

			result, errRes := f.app.Mutasi(data)
			if errRes != nil {
				remark := result.Message
				f.log.Error(logrus.Fields{
					"error": errRes.Error(),
				}, nil, remark)

				response.Message = result.Message
				response.Status = result.Status
				response.Data = nil

				err = status.Error(codes.OK, errRes.Error())
				return err
			}

			f.redis.XDel("transaction", messageID)
		}
	}

	elapsedTime := time.Since(startTime)
	f.log.Info(logrus.Fields{
		"Request":     request,
		"Result":      result,
		"error":       nil,
		"elapsedTime": elapsedTime,
	}, nil, "log info")

	return nil
}
