package payload

type GetSaldoTabunganResp struct {
	NomorRekening string  `json:"nomor_rekening"`
	Saldo         float64 `json:"saldo"`
}
