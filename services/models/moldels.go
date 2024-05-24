package models

import "gorm.io/gorm"

type Tabungan struct {
	gorm.Model
	Nominal       float64
	NomorRekening string
}

type Transaksi struct {
	gorm.Model
	Nominal       float64
	NomorRekening string
	KodeTransaksi string
}

type Nasabah struct {
	gorm.Model
	NomorRekening string `gorm:"Unique"`
	Nama          string
	Nik           string
	NoHp          string
	Pin           string
	Tabungan      Tabungan  `gorm:"foreignKey:NomorRekening"`
	Transaksi     Transaksi `gorm:"foreignKey:NomorRekening"`
}
