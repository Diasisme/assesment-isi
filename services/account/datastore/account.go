package datastore

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/services/models"
	"gorm.io/gorm"
)

func (f *databaseData) Daftar(request models.Nasabah) error {
	return f.DB.Create(&request).Error
}

func (f *databaseData) BuatTabung(request models.Tabungan) error {
	return f.DB.Create(&request).Error
}

func (f *databaseData) TambahTabung(request models.Tabungan) error {
	return f.DB.Model(&models.Tabungan{}).Where("nomor_rekening = ?", request.NomorRekening).Update("nominal", gorm.Expr("nominal + ?", request.Nominal)).Error
}

func (f *databaseData) KurangTabung(request models.Tabungan) error {
	return f.DB.Model(&models.Tabungan{}).Where("nomor_rekening = ?", request.NomorRekening).Update("nominal", gorm.Expr("nominal - ?", request.Nominal)).Error
}

func (f *databaseData) Transaksi(request models.Transaksi) error {
	return f.DB.Create(&request).Error
}

func (f *databaseData) GetDataAccount(nomor_rekening string) (models.Nasabah, error) {
	var data models.Nasabah
	result := f.DB.Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *databaseData) GetDataTabungan(nomor_rekening string) (models.Tabungan, error) {
	var data models.Tabungan
	result := f.DB.Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *databaseData) GetSaldoTabungan(nomor_rekening string) (models.Tabungan, error) {
	var data models.Tabungan
	result := f.DB.Select("nominal").Where("nomor_rekening", nomor_rekening).First(&data)
	return data, result.Error
}

func (f *databaseData) Mutasi(nomor_rekening string) ([]models.Transaksi, error) {
	var data []models.Transaksi
	result := f.DB.Find(&data)
	return data, result.Error
}
