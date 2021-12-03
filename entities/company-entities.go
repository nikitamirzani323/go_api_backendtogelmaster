package entities

type Model_company struct {
	Company_no          int    `json:"company_no"`
	Company_idcompany   string `json:"company_idcompany"`
	Company_startjoin   string `json:"company_startjoin"`
	Company_endjoin     string `json:"company_endjoin"`
	Company_curr        string `json:"company_curr"`
	Company_name        string `json:"company_name"`
	Company_periode     string `json:"company_periode"`
	Company_winlose     int    `json:"company_winlose"`
	Company_winlosetemp int    `json:"company_winlosetemp"`
	Company_status      string `json:"company_status"`
	Company_statuscss   string `json:"company_statuscss"`
}
