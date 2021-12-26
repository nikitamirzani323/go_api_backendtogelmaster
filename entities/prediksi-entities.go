package entities

type Model_pasaranlist struct {
	Pasaranlist_idpasarantogel string `json:"pasaranlist_idpasarantogel"`
	Pasaranlist_nmpasarantogel string `json:"pasaranlist_nmpasarantogel"`
}
type Model_prediksiwajib struct {
	Prediksi_idcompany   string      `json:"prediksi_idcompany"`
	Prediksi_nmcompany   string      `json:"prediksi_nmcompany"`
	Prediksi_result      interface{} `json:"prediksi_result"`
	Prediksi_totalbet    int         `json:"prediksi_totalbet"`
	Prediksi_subtotal    int         `json:"prediksi_subtotal"`
	Prediksi_subtotalwin int         `json:"prediksi_subtotalwin"`
}
type Model_listPrediksi struct {
	Prediksi_tanggal      string  `json:"prediksi_tanggal"`
	Prediksi_username     string  `json:"prediksi_username"`
	Prediksi_permainan    string  `json:"prediksi_permainan"`
	Prediksi_nomor        string  `json:"prediksi_nomor"`
	Prediksi_bet          int     `json:"prediksi_bet"`
	Prediksi_diskon       int     `json:"prediksi_diskon"`
	Prediksi_diskonpercen float32 `json:"prediksi_diskonpercen"`
	Prediksi_kei          int     `json:"prediksi_kei"`
	Prediksi_keipercen    float32 `json:"prediksi_keipercen"`
	Prediksi_bayar        int     `json:"prediksi_bayar"`
	Prediksi_win          float32 `json:"prediksi_win"`
	Prediksi_totalwin     int     `json:"prediksi_totalwin"`
	Prediksi_status       string  `json:"prediksi_status"`
	Prediksi_statuscss    string  `json:"prediksi_statuscss"`
}

type Controller_pasaranlist struct {
	Master string `json:"master" validate:"required"`
}
type Controller_prediksi struct {
	Master         string `json:"master" validate:"required"`
	Idpasarantogel string `json:"idpasarantogel" validate:"required"`
	Nomorprediksi  string `json:"nomorprediksi" validate:"required"`
}
