package models

type Tabungan struct {
	ID            uint `json:"id" db:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Nominal       float64
	NomorRekening string
	NasabahID     uint `json:"foreignKey:id"`
}

type Transaksi struct {
	ID            uint `json:"id" db:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Nominal       float64
	KodeTransaksi string
	NomorRekening string
	NasabahID     uint `json:"foreignKey:id"`
}

type Nasabah struct {
	ID            uint   `json:"id" db:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	NomorRekening string `gorm:"uniqueIndex;not null"`
	Nama          string
	Nik           string
	NoHp          string
	Pin           string      `gorm:"not null"`
	Transaksi     []Transaksi `gorm:"foreignKey:NasabahID"`
	Tabungan      []Tabungan  `gorm:"foreignKey:NasabahID"`
}
