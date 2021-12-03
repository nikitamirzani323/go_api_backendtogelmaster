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
type Model_companydetail struct {
	Company_name   string `json:"company_name"`
	Company_url    string `json:"company_url"`
	Company_status string `json:"company_status"`
	Company_create string `json:"company_create"`
	Company_update string `json:"company_update"`
}

type Model_companylistadmin struct {
	Company_admin_username      string `json:"company_admin_username"`
	Company_admin_typeadmin     string `json:"company_admin_typeadmin"`
	Company_admin_name          string `json:"company_admin_nama"`
	Company_admin_status        string `json:"company_admin_status"`
	Company_admin_statuscss     string `json:"company_admin_status_css"`
	Company_admin_lastlogin     string `json:"company_admin_lastlogin"`
	Company_admin_lastippadress string `json:"company_admin_lastipaddres"`
	Company_admin_create        string `json:"company_admin_create"`
	Company_admin_update        string `json:"company_admin_update"`
}

type Controller_company struct {
	Company_search string `json:"company_search"`
}
type Controller_companydetail struct {
	Page    string `json:"page" validate:"required"`
	Sdata   string `json:"sData" validate:"required"`
	Company string `json:"company" validate:"required"`
}
