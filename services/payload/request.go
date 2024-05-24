package payload

type DaftarReq struct {
	Nama string `json:"nama" validate:"required"`
	Nik  string `json:"nik" validate:"required"`
	NoHp string `json:"no_hp"`
	Pin  string `json:"pin" validate:"required"`
}

type TabunganReq struct {
	NomorRekening string  `json:"nomor_rekening" validate:"required"`
	Nominal       float64 `json:"nominal" validate:"required"`
}

type TransferReq struct {
	NomorRekeningAsal   string  `json:"nomor_rekening_asal" validate:"required"`
	NomorRekeningTujuan string  `json:"nomor_rekening_tujuan" validate:"required"`
	Nominal             float64 `json:"nominal" validate:"required"`
}

type GetTransaksiReq struct {
	NomorRekening string `json:"nomor_rekening"`
}
