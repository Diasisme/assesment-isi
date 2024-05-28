package payload

type MutasiReq struct {
	NomorRekening    string  `json:"nomor_rekening" validate:"required"`
	TanggalTransaksi string  `json:"tanggal_transaksi" validate:"required"`
	JenisTransaksi   string  `json:"jenis_transaksi" validate:"required"`
	Nominal          float64 `json:"nominal" validate:"required"`
}
