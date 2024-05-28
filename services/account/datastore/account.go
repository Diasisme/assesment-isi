package datastore

import (
	"math"
	"time"

	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (f *DatabaseData) Daftar(request models.Nasabah) error {
	return f.DB.Create(&request).Error
}

func (f *DatabaseData) BuatTabung(request models.Tabungan) error {
	return f.DB.Create(&request).Error
}

func (f *DatabaseData) TambahTabung(request models.Tabungan) error {

	f.log.Info(logrus.Fields{"request": request}, nil, "Log info: TambahTabung")
	err := f.redis.XAdd(&redis.XAddArgs{
		Stream: "transaction",
		Values: map[string]any{"date": time.Now().Format("2006-01-02 15:04:05"), "jenis_transaksi": "D", "amount": math.Abs(request.Nominal), "account_number": request.NomorRekening},
	}).Err()
	if err != nil {
		f.log.Error(map[string]any{"account_number": request.NomorRekening, "amount": request.Nominal, "error": err}, nil, "database error")
	}
	return f.DB.Model(&models.Tabungan{}).Where("nomor_rekening = ?", request.NomorRekening).Update("nominal", gorm.Expr("nominal + ?", request.Nominal)).Error
}

func (f *DatabaseData) KurangTabung(request models.Tabungan) error {
	f.log.Info(logrus.Fields{"request": request}, nil, "Log info: KurangTabung")
	err := f.redis.XAdd(&redis.XAddArgs{
		Stream: "transaction",
		Values: map[string]any{"date": time.Now().Format("2006-01-02 15:04:05"), "jenis_transaksi": "C", "amount": math.Abs(request.Nominal), "account_number": request.NomorRekening},
	}).Err()
	if err != nil {
		f.log.Error(map[string]any{"account_number": request.NomorRekening, "amount": request.Nominal, "error": err}, nil, "database error")
	}
	return f.DB.Model(&models.Tabungan{}).Where("nomor_rekening = ?", request.NomorRekening).Update("nominal", gorm.Expr("nominal - ?", request.Nominal)).Error
}

func (f *DatabaseData) Transaksi(request models.Transaksi) error {
	return f.DB.Create(&request).Error
}

func (f *DatabaseData) GetDataAccount(nomor_rekening string) (models.Nasabah, error) {
	var data models.Nasabah
	result := f.DB.Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *DatabaseData) GetDataTabungan(nomor_rekening string) (models.Tabungan, error) {
	var data models.Tabungan
	result := f.DB.Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *DatabaseData) GetSaldoTabungan(nomor_rekening string) (models.Tabungan, error) {
	var data models.Tabungan
	result := f.DB.Select("nominal").Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *DatabaseData) Mutasi(nomor_rekening string) ([]models.Transaksi, error) {
	var data []models.Transaksi
	result := f.DB.Find(&data)
	return data, result.Error
}
