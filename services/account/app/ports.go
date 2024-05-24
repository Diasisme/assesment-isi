package app

import (
	"github.com/Diasisme/asssesment-march-ihsan.git/services/helpers"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/models"
	"github.com/Diasisme/asssesment-march-ihsan.git/services/payload"
)

type AccountDatastore interface {
	Daftar(request models.Nasabah) error
	BuatTabung(request models.Tabungan) error
	TambahTabung(request models.Tabungan) error
	KurangTabung(request models.Tabungan) error
	Transaksi(request models.Transaksi) error
	GetDataAccount(nomor_rekening string) (models.Nasabah, error)
	GetDataTabungan(nomor_rekening string) (models.Tabungan, error)
	GetSaldoTabungan(nomor_rekening string) (models.Tabungan, error)
	Mutasi(nomor_rekening string) ([]models.Transaksi, error)
}

type AccountApp interface {
	Daftar(request models.Nasabah) (response helpers.Response, err error)
	Tabung(request models.Tabungan) (response helpers.Response, err error)
	Tarik(request models.Tabungan) (response helpers.Response, err error)
	Transfer(request payload.TransferReq) (response helpers.Response, err error)
	GetSaldoTabungan(request payload.GetTransaksiReq) (response helpers.Response, err error)
	GetMutasi(request payload.GetTransaksiReq) (response helpers.Response, err error)
}
