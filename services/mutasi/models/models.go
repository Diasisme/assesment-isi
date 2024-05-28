package models

import (
	"time"

	"gorm.io/gorm"
)

type Mutasi struct {
	gorm.Model
	TanggalTransaksi time.Time
	NomorRekening    string
	JenisTransaksi   string
	Nominal          float64
}