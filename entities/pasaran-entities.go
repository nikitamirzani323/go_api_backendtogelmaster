package entities

type Model_pasaran struct {
	Pasaran_no             int    `json:"pasaran_no"`
	Pasaran_idpasarantogel string `json:"pasaran_idpasarantogel"`
	Pasaran_nmpasarantogel string `json:"pasaran_nmpasarantogel"`
	Pasaran_tipepasaran    string `json:"pasaran_tipepasaran"`
	Pasaran_urlpasaran     string `json:"pasaran_urlpasaran"`
	Pasaran_pasarandiundi  string `json:"pasaran_pasarandiundi"`
	Pasaran_jamtutup       string `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal      string `json:"pasaran_jamjadwal"`
	Pasaran_jamopen        string `json:"pasaran_jamopen"`
}

type Controller_pasaran struct {
	Master string `json:"master" validate:"required"`
}
