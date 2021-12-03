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

type Model_companylistpasaran struct {
	Company_pasaran_idcomppasaran   int    `json:"company_pasaran_idcomppasaran"`
	Company_pasaran_idpasarantogel  string `json:"company_pasaran_idpasarantogel"`
	Company_pasaran_nmpasarantogel  string `json:"company_pasaran_nmpasarantogel"`
	Company_pasaran_periode         string `json:"company_pasaran_periode"`
	Company_pasaran_winlose         int    `json:"company_pasaran_winlose"`
	Company_pasaran_displaypasaran  int    `json:"company_pasaran_displaypasaran"`
	Company_pasaran_status          string `json:"company_pasaran_status"`
	Company_pasaran_statuscss       string `json:"company_pasaran_statuscss"`
	Company_pasaran_statusactive    string `json:"company_pasaran_statuspasaranactive"`
	Company_pasaran_statusactivecss string `json:"company_pasaran_statuspasaranactivecss"`
}
type Model_companypasaranconf struct {
	Company_Pasaran_diundi                    string  `json:"pasaran_diundi"`
	Company_Pasaran_url                       string  `json:"pasaran_url"`
	Company_Pasaran_jamtutup                  string  `json:"pasaran_jamtutup"`
	Company_Pasaran_jamjadwal                 string  `json:"pasaran_jamjadwal"`
	Company_Pasaran_jamopen                   string  `json:"pasaran_jamopen"`
	Company_Pasaran_statusactive              string  `json:"pasaran_statusactive"`
	Company_Limitline4d                       int     `json:"limitline_4d"`
	Company_Limitline3d                       int     `json:"limitline_3d"`
	Company_Limitline2d                       int     `json:"limitline_2d"`
	Company_Limitline2dd                      int     `json:"limitline_2dd"`
	Company_Limitline2dt                      int     `json:"limitline_2dt"`
	Company_Bbfs                              int     `json:"bbfs"`
	Company_Minbet_432d                       float32 `json:"minbet_432d"`
	Company_Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Company_Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Company_Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Company_Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Company_Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Company_Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Company_Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Company_Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Company_Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Company_Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Company_Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Company_Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Company_Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Company_Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Company_Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Company_Disc4d_432d                       float32 `json:"disc4d_432d"`
	Company_Disc3d_432d                       float32 `json:"disc3d_432d"`
	Company_Disc2d_432d                       float32 `json:"disc2d_432d"`
	Company_Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Company_Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Company_Win4d_432d                        float32 `json:"win4d_432d"`
	Company_Win3d_432d                        float32 `json:"win3d_432d"`
	Company_Win2d_432d                        float32 `json:"win2d_432d"`
	Company_Win2dd_432d                       float32 `json:"win2dd_432d"`
	Company_Win2dt_432d                       float32 `json:"win2dt_432d"`
	Company_Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Company_Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Company_Win_cbebas                        float32 `json:"win_cbebas"`
	Company_Disc_cbebas                       float32 `json:"disc_cbebas"`
	Company_Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Company_Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Company_Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Company_Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Company_Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Company_Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Company_Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Company_Disc_cmacau                       float32 `json:"disc_cmacau"`
	Company_Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Company_Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Company_Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Company_Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Company_Win3_cnaga                        float32 `json:"win3_cnaga"`
	Company_Win4_cnaga                        float32 `json:"win4_cnaga"`
	Company_Disc_cnaga                        float32 `json:"disc_cnaga"`
	Company_Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Company_Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Company_Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Company_Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Company_Winas_cjitu                       float32 `json:"winas_cjitu"`
	Company_Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Company_Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Company_Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Company_Desc_cjitu                        float32 `json:"desc_cjitu"`
	Company_Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Company_Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Company_Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Company_Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Company_Keibesar_5050umum                 float32 `json:"keibesar_5050umum"`
	Company_Keikecil_5050umum                 float32 `json:"keikecil_5050umum"`
	Company_Keigenap_5050umum                 float32 `json:"keigenap_5050umum"`
	Company_Keiganjil_5050umum                float32 `json:"keiganjil_5050umum"`
	Company_Keitengah_5050umum                float32 `json:"keitengah_5050umum"`
	Company_Keitepi_5050umum                  float32 `json:"keitepi_5050umum"`
	Company_Discbesar_5050umum                float32 `json:"discbesar_5050umum"`
	Company_Disckecil_5050umum                float32 `json:"disckecil_5050umum"`
	Company_Discgenap_5050umum                float32 `json:"discgenap_5050umum"`
	Company_Discganjil_5050umum               float32 `json:"discganjil_5050umum"`
	Company_Disctengah_5050umum               float32 `json:"disctengah_5050umum"`
	Company_Disctepi_5050umum                 float32 `json:"disctepi_5050umum"`
	Company_Limitglobal_5050umum              float32 `json:"limitglobal_5050umum"`
	Company_Limittotal_5050umum               float32 `json:"limittotal_5050umum"`
	Company_Minbet_5050special                float32 `json:"minbet_5050special"`
	Company_Maxbet_5050special                float32 `json:"maxbet_5050special"`
	Company_Keiasganjil_5050special           float32 `json:"keiasganjil_5050special"`
	Company_Keiasgenap_5050special            float32 `json:"keiasgenap_5050special"`
	Company_Keiasbesar_5050special            float32 `json:"keiasbesar_5050special"`
	Company_Keiaskecil_5050special            float32 `json:"keiaskecil_5050special"`
	Company_Keikopganjil_5050special          float32 `json:"keikopganjil_5050special"`
	Company_Keikopgenap_5050special           float32 `json:"keikopgenap_5050special"`
	Company_Keikopbesar_5050special           float32 `json:"keikopbesar_5050special"`
	Company_Keikopkecil_5050special           float32 `json:"keikopkecil_5050special"`
	Company_Keikepalaganjil_5050special       float32 `json:"keikepalaganjil_5050special"`
	Company_Keikepalagenap_5050special        float32 `json:"keikepalagenap_5050special"`
	Company_Keikepalabesar_5050special        float32 `json:"keikepalabesar_5050special"`
	Company_Keikepalakecil_5050special        float32 `json:"keikepalakecil_5050special"`
	Company_Keiekorganjil_5050special         float32 `json:"keiekorganjil_5050special"`
	Company_Keiekorgenap_5050special          float32 `json:"keiekorgenap_5050special"`
	Company_Keiekorbesar_5050special          float32 `json:"keiekorbesar_5050special"`
	Company_Keiekorkecil_5050special          float32 `json:"keiekorkecil_5050special"`
	Company_Discasganjil_5050special          float32 `json:"discasganjil_5050special"`
	Company_Discasgenap_5050special           float32 `json:"discasgenap_5050special"`
	Company_Discasbesar_5050special           float32 `json:"discasbesar_5050special"`
	Company_Discaskecil_5050special           float32 `json:"discaskecil_5050special"`
	Company_Disckopganjil_5050special         float32 `json:"disckopganjil_5050special"`
	Company_Disckopgenap_5050special          float32 `json:"disckopgenap_5050special"`
	Company_Disckopbesar_5050special          float32 `json:"disckopbesar_5050special"`
	Company_Disckopkecil_5050special          float32 `json:"disckopkecil_5050special"`
	Company_Disckepalaganjil_5050special      float32 `json:"disckepalaganjil_5050special"`
	Company_Disckepalagenap_5050special       float32 `json:"disckepalagenap_5050special"`
	Company_Disckepalabesar_5050special       float32 `json:"disckepalabesar_5050special"`
	Company_Disckepalakecil_5050special       float32 `json:"disckepalakecil_5050special"`
	Company_Discekorganjil_5050special        float32 `json:"discekorganjil_5050special"`
	Company_Discekorgenap_5050special         float32 `json:"discekorgenap_5050special"`
	Company_Discekorbesar_5050special         float32 `json:"discekorbesar_5050special"`
	Company_Discekorkecil_5050special         float32 `json:"discekorkecil_5050special"`
	Company_Limitglobal_5050special           float32 `json:"limitglobal_5050special"`
	Company_Limittotal_5050special            float32 `json:"limittotal_5050special"`
	Company_Minbet_5050kombinasi              float32 `json:"minbet_5050kombinasi"`
	Company_Maxbet_5050kombinasi              float32 `json:"maxbet_5050kombinasi"`
	Company_Belakangkeimono_5050kombinasi     float32 `json:"belakangkeimono_5050kombinasi"`
	Company_Belakangkeistereo_5050kombinasi   float32 `json:"belakangkeistereo_5050kombinasi"`
	Company_Belakangkeikembang_5050kombinasi  float32 `json:"belakangkeikembang_5050kombinasi"`
	Company_Belakangkeikempis_5050kombinasi   float32 `json:"belakangkeikempis_5050kombinasi"`
	Company_Belakangkeikembar_5050kombinasi   float32 `json:"belakangkeikembar_5050kombinasi"`
	Company_Tengahkeimono_5050kombinasi       float32 `json:"tengahkeimono_5050kombinasi"`
	Company_Tengahkeistereo_5050kombinasi     float32 `json:"tengahkeistereo_5050kombinasi"`
	Company_Tengahkeikembang_5050kombinasi    float32 `json:"tengahkeikembang_5050kombinasi"`
	Company_Tengahkeikempis_5050kombinasi     float32 `json:"tengahkeikempis_5050kombinasi"`
	Company_Tengahkeikembar_5050kombinasi     float32 `json:"tengahkeikembar_5050kombinasi"`
	Company_Depankeimono_5050kombinasi        float32 `json:"depankeimono_5050kombinasi"`
	Company_Depankeistereo_5050kombinasi      float32 `json:"depankeistereo_5050kombinasi"`
	Company_Depankeikembang_5050kombinasi     float32 `json:"depankeikembang_5050kombinasi"`
	Company_Depankeikempis_5050kombinasi      float32 `json:"depankeikempis_5050kombinasi"`
	Company_Depankeikembar_5050kombinasi      float32 `json:"depankeikembar_5050kombinasi"`
	Company_Belakangdiscmono_5050kombinasi    float32 `json:"belakangdiscmono_5050kombinasi"`
	Company_Belakangdiscstereo_5050kombinasi  float32 `json:"belakangdiscstereo_5050kombinasi"`
	Company_Belakangdisckembang_5050kombinasi float32 `json:"belakangdisckembang_5050kombinasi"`
	Company_Belakangdisckempis_5050kombinasi  float32 `json:"belakangdisckempis_5050kombinasi"`
	Company_Belakangdisckembar_5050kombinasi  float32 `json:"belakangdisckembar_5050kombinasi"`
	Company_Tengahdiscmono_5050kombinasi      float32 `json:"tengahdiscmono_5050kombinasi"`
	Company_Tengahdiscstereo_5050kombinasi    float32 `json:"tengahdiscstereo_5050kombinasi"`
	Company_Tengahdisckembang_5050kombinasi   float32 `json:"tengahdisckembang_5050kombinasi"`
	Company_Tengahdisckempis_5050kombinasi    float32 `json:"tengahdisckempis_5050kombinasi"`
	Company_Tengahdisckembar_5050kombinasi    float32 `json:"tengahdisckembar_5050kombinasi"`
	Company_Depandiscmono_5050kombinasi       float32 `json:"depandiscmono_5050kombinasi"`
	Company_Depandiscstereo_5050kombinasi     float32 `json:"depandiscstereo_5050kombinasi"`
	Company_Depandisckembang_5050kombinasi    float32 `json:"depandisckembang_5050kombinasi"`
	Company_Depandisckempis_5050kombinasi     float32 `json:"depandisckempis_5050kombinasi"`
	Company_Depandisckembar_5050kombinasi     float32 `json:"depandisckembar_5050kombinasi"`
	Company_Limitglobal_5050kombinasi         float32 `json:"limitglobal_5050kombinasi"`
	Company_Limittotal_5050kombinasi          float32 `json:"limittotal_5050kombinasi"`
	Company_Minbet_kombinasi                  float32 `json:"minbet_kombinasi"`
	Company_Maxbet_kombinasi                  float32 `json:"maxbet_kombinasi"`
	Company_Win_kombinasi                     float32 `json:"win_kombinasi"`
	Company_Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Company_Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Company_Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Company_Minbet_dasar                      float32 `json:"minbet_dasar"`
	Company_Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Company_Keibesar_dasar                    float32 `json:"keibesar_dasar"`
	Company_Keikecil_dasar                    float32 `json:"keikecil_dasar"`
	Company_Keigenap_dasar                    float32 `json:"keigenap_dasar"`
	Company_Keiganjil_dasar                   float32 `json:"keiganjil_dasar"`
	Company_Discbesar_dasar                   float32 `json:"discbesar_dasar"`
	Company_Disckecil_dasar                   float32 `json:"disckecil_dasar"`
	Company_Discgenap_dasar                   float32 `json:"discgenap_dasar"`
	Company_Discganjil_dasar                  float32 `json:"discganjil_dasar"`
	Company_Limitglobal_dasar                 float32 `json:"limitglobal_dasar"`
	Company_Limittotal_dasar                  float32 `json:"limittotal_dasar"`
	Company_Minbet_shio                       float32 `json:"minbet_shio"`
	Company_Maxbet_shio                       float32 `json:"maxbet_shio"`
	Company_Win_shio                          float32 `json:"win_shio"`
	Company_Disc_shio                         float32 `json:"disc_shio"`
	Company_Shioyear_shio                     string  `json:"shioyear_shio"`
	Company_Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Company_Limittotal_shio                   float32 `json:"limittotal_shio"`
}

type Controller_company struct {
	Company_search string `json:"company_search"`
}
type Controller_companydetail struct {
	Page    string `json:"page" validate:"required"`
	Sdata   string `json:"sData" validate:"required"`
	Company string `json:"company" validate:"required"`
}
