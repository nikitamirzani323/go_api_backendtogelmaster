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

type Model_companylistpasaranonline struct {
	Company_pasaran_onlineid int    `json:"company_pasaranonline_id"`
	Company_pasaran_harian   string `json:"company_pasaranonline_hari"`
}
type Model_companypasaranconf struct {
	Company_Pasaran_diundi                    string  `json:"pasaran_diundi"`
	Company_Pasaran_url                       string  `json:"pasaran_url"`
	Company_Pasaran_jamtutup                  string  `json:"pasaran_jamtutup"`
	Company_Pasaran_jamjadwal                 string  `json:"pasaran_jamjadwal"`
	Company_Pasaran_jamopen                   string  `json:"pasaran_jamopen"`
	Company_Pasaran_statusactive              string  `json:"pasaran_statusactive"`
	Company_Royaltyfee                        float32 `json:"royaltyfee"`
	Company_Limitline4d                       int     `json:"limitline_4d"`
	Company_Limitline3d                       int     `json:"limitline_3d"`
	Company_Limitline3dd                      int     `json:"limitline_3dd"`
	Company_Limitline2d                       int     `json:"limitline_2d"`
	Company_Limitline2dd                      int     `json:"limitline_2dd"`
	Company_Limitline2dt                      int     `json:"limitline_2dt"`
	Company_Bbfs                              int     `json:"bbfs"`
	Company_Minbet_432d                       float32 `json:"minbet_432d"`
	Company_Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Company_Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Company_Maxbet3dd_432d                    float32 `json:"maxbet3dd_432d"`
	Company_Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Company_Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Company_Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Company_Maxbet4d_fullbb_432d              float32 `json:"maxbet4d_fullbb_432d"`
	Company_Maxbet3d_fullbb_432d              float32 `json:"maxbet3d_fullbb_432d"`
	Company_Maxbet3dd_fullbb_432d             float32 `json:"maxbet3dd_fullbb_432d"`
	Company_Maxbet2d_fullbb_432d              float32 `json:"maxbet2d_fullbb_432d"`
	Company_Maxbet2dd_fullbb_432d             float32 `json:"maxbet2dd_fullbb_432d"`
	Company_Maxbet2dt_fullbb_432d             float32 `json:"maxbet2dt_fullbb_432d"`
	Company_Maxbuy4d_432d                     float32 `json:"maxbuy4d_432d"`
	Company_Maxbuy3d_432d                     float32 `json:"maxbuy3d_432d"`
	Company_Maxbuy3dd_432d                    float32 `json:"maxbuy3dd_432d"`
	Company_Maxbuy2d_432d                     float32 `json:"maxbuy2d_432d"`
	Company_Maxbuy2dd_432d                    float32 `json:"maxbuy2dd_432d"`
	Company_Maxbuy2dt_432d                    float32 `json:"maxbuy2dt_432d"`
	Company_Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Company_Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Company_Limitotal3dd_432d                 float32 `json:"limitotal3dd_432d"`
	Company_Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Company_Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Company_Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Company_Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Company_Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Company_Limitglobal3dd_432d               float32 `json:"limitglobal3dd_432d"`
	Company_Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Company_Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Company_Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Company_Limitotal4d_fullbb_432d           float32 `json:"limitotal4d_fullbb_432d"`
	Company_Limitotal3d_fullbb_432d           float32 `json:"limitotal3d_fullbb_432d"`
	Company_Limitotal3dd_fullbb_432d          float32 `json:"limitotal3dd_fullbb_432d"`
	Company_Limitotal2d_fullbb_432d           float32 `json:"limitotal2d_fullbb_432d"`
	Company_Limitotal2dd_fullbb_432d          float32 `json:"limitotal2dd_fullbb_432d"`
	Company_Limitotal2dt_fullbb_432d          float32 `json:"limitotal2dt_fullbb_432d"`
	Company_Limitglobal4d_fullbb_432d         float32 `json:"limitglobal4d_fullbb_432d"`
	Company_Limitglobal3d_fullbb_432d         float32 `json:"limitglobal3d_fullbb_432d"`
	Company_Limitglobal3dd_fullbb_432d        float32 `json:"limitglobal3dd_fullbb_432d"`
	Company_Limitglobal2d_fullbb_432d         float32 `json:"limitglobal2d_fullbb_432d"`
	Company_Limitglobal2dd_fullbb_432d        float32 `json:"limitglobal2dd_fullbb_432d"`
	Company_Limitglobal2dt_fullbb_432d        float32 `json:"limitglobal2dt_fullbb_432d"`
	Company_Disc4d_432d                       float32 `json:"disc4d_432d"`
	Company_Disc3d_432d                       float32 `json:"disc3d_432d"`
	Company_Disc3dd_432d                      float32 `json:"disc3dd_432d"`
	Company_Disc2d_432d                       float32 `json:"disc2d_432d"`
	Company_Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Company_Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Company_Win4d_432d                        float32 `json:"win4d_432d"`
	Company_Win3d_432d                        float32 `json:"win3d_432d"`
	Company_Win3dd_432d                       float32 `json:"win3dd_432d"`
	Company_Win2d_432d                        float32 `json:"win2d_432d"`
	Company_Win2dd_432d                       float32 `json:"win2dd_432d"`
	Company_Win2dt_432d                       float32 `json:"win2dt_432d"`
	Company_Win4dnodisc_432d                  float32 `json:"win4dnodisc_432d"`
	Company_Win3dnodisc_432d                  float32 `json:"win3dnodisc_432d"`
	Company_Win3ddnodisc_432d                 float32 `json:"win3ddnodisc_432d"`
	Company_Win2dnodisc_432d                  float32 `json:"win2dnodisc_432d"`
	Company_Win2ddnodisc_432d                 float32 `json:"win2ddnodisc_432d"`
	Company_Win2dtnodisc_432d                 float32 `json:"win2dtnodisc_432d"`
	Company_Win4dbb_kena_432d                 float32 `json:"win4dbb_kena_432d"`
	Company_Win3dbb_kena_432d                 float32 `json:"win3dbb_kena_432d"`
	Company_Win3ddbb_kena_432d                float32 `json:"win3ddbb_kena_432d"`
	Company_Win2dbb_kena_432d                 float32 `json:"win2dbb_kena_432d"`
	Company_Win2ddbb_kena_432d                float32 `json:"win2ddbb_kena_432d"`
	Company_Win2dtbb_kena_432d                float32 `json:"win2dtbb_kena_432d"`
	Company_Win4dbb_432d                      float32 `json:"win4dbb_432d"`
	Company_Win3dbb_432d                      float32 `json:"win3dbb_432d"`
	Company_Win3ddbb_432d                     float32 `json:"win3ddbb_432d"`
	Company_Win2dbb_432d                      float32 `json:"win2dbb_432d"`
	Company_Win2ddbb_432d                     float32 `json:"win2ddbb_432d"`
	Company_Win2dtbb_432d                     float32 `json:"win2dtbb_432d"`
	Company_Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Company_Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Company_Maxbuy_cbebas                     float32 `json:"maxbuy_cbebas"`
	Company_Win_cbebas                        float32 `json:"win_cbebas"`
	Company_Disc_cbebas                       float32 `json:"disc_cbebas"`
	Company_Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Company_Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Company_Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Company_Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Company_Maxbuy_cmacau                     float32 `json:"maxbuy_cmacau"`
	Company_Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Company_Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Company_Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Company_Disc_cmacau                       float32 `json:"disc_cmacau"`
	Company_Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Company_Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Company_Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Company_Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Company_Maxbuy_cnaga                      float32 `json:"maxbuy_cnaga"`
	Company_Win3_cnaga                        float32 `json:"win3_cnaga"`
	Company_Win4_cnaga                        float32 `json:"win4_cnaga"`
	Company_Disc_cnaga                        float32 `json:"disc_cnaga"`
	Company_Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Company_Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Company_Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Company_Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Company_Maxbuy_cjitu                      float32 `json:"maxbuy_cjitu"`
	Company_Winas_cjitu                       float32 `json:"winas_cjitu"`
	Company_Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Company_Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Company_Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Company_Desc_cjitu                        float32 `json:"desc_cjitu"`
	Company_Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Company_Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Company_Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Company_Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Company_Maxbuy_5050umum                   float32 `json:"maxbuy_5050umum"`
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
	Company_Maxbuy_5050special                float32 `json:"maxbuy_5050special"`
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
	Company_Maxbuy_5050kombinasi              float32 `json:"maxbuy_5050kombinasi"`
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
	Company_Maxbuy_kombinasi                  float32 `json:"maxbuy_kombinasi"`
	Company_Win_kombinasi                     float32 `json:"win_kombinasi"`
	Company_Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Company_Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Company_Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Company_Minbet_dasar                      float32 `json:"minbet_dasar"`
	Company_Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Company_Maxbuy_dasar                      float32 `json:"maxbuy_dasar"`
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
	Company_Maxbuy_shio                       float32 `json:"maxbuy_shio"`
	Company_Win_shio                          float32 `json:"win_shio"`
	Company_Disc_shio                         float32 `json:"disc_shio"`
	Company_Shioyear_shio                     string  `json:"shioyear_shio"`
	Company_Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Company_Limittotal_shio                   float32 `json:"limittotal_shio"`
}
type Model_companylistkeluaran struct {
	Company_Pasaran_no                int     `json:"company_pasaran_no"`
	Company_Pasaran_idtrxkeluaran     int     `json:"company_pasaran_invoice"`
	Company_Pasaran_idcomppasaran     int     `json:"company_pasaran_idcompp"`
	Company_Pasaran_pasarancode       string  `json:"company_pasaran_code"`
	Company_Pasaran_keluaranperiode   string  `json:"company_pasaran_periode"`
	Company_Pasaran_nmpasaran         string  `json:"company_pasaran_name"`
	Company_Pasaran_tanggalperiode    string  `json:"company_pasaran_tanggal"`
	Company_Pasaran_keluarantogel     string  `json:"company_pasaran_keluaran"`
	Company_Pasaran_status            string  `json:"company_pasaran_status"`
	Company_Pasaran_status_css        string  `json:"company_pasaran_status_css"`
	Company_Pasaran_total_Member      float32 `json:"company_pasaran_totalmember"`
	Company_Pasaran_total_bet         float32 `json:"company_pasaran_totalbet"`
	Company_Pasaran_total_outstanding float32 `json:"company_pasaran_totaloutstanding"`
	Company_Pasaran_total_cancelbet   float32 `json:"company_pasaran_totalcancelbet"`
	Company_Pasaran_winlose           float32 `json:"company_pasaran_winlose"`
	Company_Pasaran_winlosetemp       int     `json:"company_pasaran_winlosetemp"`
	Company_Pasaran_revisi            int     `json:"company_pasaran_revisi"`
	Company_Pasaran_msgrevisi         string  `json:"company_pasaran_msgrevisi"`
}
type Model_invoicelistMember struct {
	Member         string `json:"member"`
	Totalbet       int    `json:"totalbet"`
	Totalbayar     int    `json:"totalbayar"`
	Totalcancelbet int    `json:"totalcancelbet"`
	Totalwin       int    `json:"totalwin"`
}
type Model_invoicelistpermainan struct {
	Bet_id           int     `json:"bet_id"`
	Bet_datetime     string  `json:"bet_datetime"`
	Bet_ipaddress    string  `json:"bet_ipaddress"`
	Bet_device       string  `json:"bet_device"`
	Bet_timezone     string  `json:"bet_timezone"`
	Bet_username     string  `json:"bet_username"`
	Bet_typegame     string  `json:"bet_typegame"`
	Bet_nomortogel   string  `json:"bet_nomortogel"`
	Bet_posisitogel  string  `json:"bet_posisitogel"`
	Bet_bet          int     `json:"bet_bet"`
	Bet_diskon       int     `json:"bet_diskon"`
	Bet_diskonpercen int     `json:"bet_diskonpercen"`
	Bet_kei          int     `json:"bet_kei"`
	Bet_keipercen    int     `json:"bet_keipercen"`
	Bet_win          float32 `json:"bet_win"`
	Bet_totalwin     int     `json:"bet_totalwin"`
	Bet_bayar        int     `json:"bet_bayar"`
	Bet_status       string  `json:"bet_status"`
	Bet_statuscss    string  `json:"bet_statuscss"`
	Bet_create       string  `json:"bet_create"`
	Bet_createDate   string  `json:"bet_createdate"`
	Bet_update       string  `json:"bet_update"`
	Bet_updateDate   string  `json:"bet_updatedate"`
}
type Model_invoicelistGroupPermainan struct {
	Permainan string `json:"permainan"`
}

type Controller_company struct {
	Company_search string `json:"company_search"`
}
type Controller_companydetail struct {
	Page    string `json:"page" validate:"required"`
	Sdata   string `json:"sData" validate:"required"`
	Company string `json:"company" validate:"required"`
}

type Controller_companypasaranconf struct {
	Page              string `json:"page" validate:"required"`
	Sdata             string `json:"sData" validate:"required"`
	Company           string `json:"company" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
}
type Controller_companylistkeluaran struct {
	Company string `json:"company" validate:"required"`
	Periode string `json:"periode" validate:"required"`
	Year    string `json:"year" validate:"required"`
	Pasaran int    `json:"pasaran" validate:"required"`
}
type Controller_companyinvoice struct {
	Company  string `json:"company" validate:"required"`
	Username string `json:"username" `
	Invoice  int    `json:"invoice" validate:"required"`
	Month    string `json:"month" validate:"required"`
	Year     string `json:"year" validate:"required"`
	Pasaran  string `json:"pasaran" validate:"required"`
}
type Controller_companysave struct {
	Sdata     string `json:"sdata" validate:"required"`
	Company   string `json:"company" validate:"required"`
	Master    string `json:"master" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Urldomain string `json:"urldomain" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
type Controller_companyadminsave struct {
	Sdata          string `json:"sdata" validate:"required"`
	Company        string `json:"company" validate:"required"`
	Master         string `json:"master" validate:"required"`
	Admin_username string `json:"admin_username" validate:"required,alphanum"`
	Admin_password string `json:"admin_password" `
	Admin_name     string `json:"admin_name" validate:"required,alphanum"`
	Admin_status   string `json:"admin_status" validate:"required"`
}
type Controller_companypasaransave struct {
	Sdata      string `json:"sdata" validate:"required"`
	Company    string `json:"company" validate:"required"`
	Master     string `json:"master" validate:"required"`
	Pasaran_id string `json:"pasaran_id" validate:"required"`
}
type Controller_companypasaran struct {
	Company              string `json:"company" validate:"required"`
	Master               string `json:"master" validate:"required"`
	Companypasaran_id    int    `json:"companypasaran_id" validate:"required"`
	Pasaran_diundi       string `json:"pasaran_diundi" validate:"required"`
	Pasaran_url          string `json:"pasaran_url" validate:"required"`
	Pasaran_jamtutup     string `json:"pasaran_jamtutup" validate:"required"`
	Pasaran_jamjadwal    string `json:"pasaran_jamjadwal" validate:"required"`
	Pasaran_jamopen      string `json:"pasaran_jamopen" validate:"required"`
	Pasaran_statusactive string `json:"pasaran_statusactive" validate:"required"`
}
type Controller_companypasaranroyaltyfee struct {
	Company           string  `json:"company" validate:"required"`
	Master            string  `json:"master" validate:"required"`
	Companypasaran_id int     `json:"companypasaran_id" validate:"required"`
	Royaltyfee        float32 `json:"royaltyfee" validate:"required"`
}
type Controller_companypasaranline struct {
	Company               string `json:"company" validate:"required"`
	Master                string `json:"master" validate:"required"`
	Pasaran_id            string `json:"pasaran_id" validate:"required"`
	Companypasaran_id     int    `json:"companypasaran_id" validate:"required"`
	Pasaran_limitline_4d  int    `json:"pasaran_limitline_4d" validate:"required"`
	Pasaran_limitline_3d  int    `json:"pasaran_limitline_3d" validate:"required"`
	Pasaran_limitline_3dd int    `json:"pasaran_limitline_3dd" validate:"required"`
	Pasaran_limitline_2d  int    `json:"pasaran_limitline_2d" validate:"required"`
	Pasaran_limitline_2dd int    `json:"pasaran_limitline_2dd" validate:"required"`
	Pasaran_limitline_2dt int    `json:"pasaran_limitline_2dt" validate:"required"`
	Pasaran_bbfs          int    `json:"pasaran_bbfs" validate:"required"`
}
type Controller_companypasaran432 struct {
	Company                            string  `json:"company" validate:"required"`
	Master                             string  `json:"master" validate:"required"`
	Pasaran_id                         string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id                  int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_432d                int     `json:"pasaran_minbet_432d" validate:"required,numeric"`
	Pasaran_maxbet4d_432d              int     `json:"pasaran_maxbet4d_432d" validate:"required,numeric"`
	Pasaran_maxbet3d_432d              int     `json:"pasaran_maxbet3d_432d" validate:"required,numeric"`
	Pasaran_maxbet3dd_432d             int     `json:"pasaran_maxbet3dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2d_432d              int     `json:"pasaran_maxbet2d_432d" validate:"required,numeric"`
	Pasaran_maxbet2dd_432d             int     `json:"pasaran_maxbet2dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2dt_432d             int     `json:"pasaran_maxbet2dt_432d" validate:"required,numeric"`
	Pasaran_maxbet4d_fullbb_432d       int     `json:"pasaran_maxbet4d_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbet3d_fullbb_432d       int     `json:"pasaran_maxbet3d_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbet3dd_fullbb_432d      int     `json:"pasaran_maxbet3dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbet2d_fullbb_432d       int     `json:"pasaran_maxbet2d_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbet2dd_fullbb_432d      int     `json:"pasaran_maxbet2dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbet2dt_fullbb_432d      int     `json:"pasaran_maxbet2dt_fullbb_432d" validate:"required,numeric"`
	Pasaran_maxbuy4d_432d              int     `json:"pasaran_maxbuy4d_432d" validate:"required,numeric"`
	Pasaran_maxbuy3d_432d              int     `json:"pasaran_maxbuy3d_432d" validate:"required,numeric"`
	Pasaran_maxbuy3dd_432d             int     `json:"pasaran_maxbuy3dd_432d" validate:"required,numeric"`
	Pasaran_maxbuy2d_432d              int     `json:"pasaran_maxbuy2d_432d" validate:"required,numeric"`
	Pasaran_maxbuy2dd_432d             int     `json:"pasaran_maxbuy2dd_432d" validate:"required,numeric"`
	Pasaran_maxbuy2dt_432d             int     `json:"pasaran_maxbuy2dt_432d" validate:"required,numeric"`
	Pasaran_limitotal4d_432d           int     `json:"pasaran_limitotal4d_432d" validate:"required,numeric"`
	Pasaran_limitotal3d_432d           int     `json:"pasaran_limitotal3d_432d" validate:"required,numeric"`
	Pasaran_limitotal3dd_432d          int     `json:"pasaran_limitotal3dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2d_432d           int     `json:"pasaran_limitotal2d_432d" validate:"required,numeric"`
	Pasaran_limitotal2dd_432d          int     `json:"pasaran_limitotal2dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2dt_432d          int     `json:"pasaran_limitotal2dt_432d" validate:"required,numeric"`
	Pasaran_limitglobal4d_432d         int     `json:"pasaran_limitglobal4d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3d_432d         int     `json:"pasaran_limitglobal3d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3dd_432d        int     `json:"pasaran_limitglobal3dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2d_432d         int     `json:"pasaran_limitglobal2d_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dd_432d        int     `json:"pasaran_limitglobal2dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dt_432d        int     `json:"pasaran_limitglobal2dt_432d" validate:"required,numeric"`
	Pasaran_limitotal4d_fullbb_432d    int     `json:"pasaran_limitotal4d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitotal3d_fullbb_432d    int     `json:"pasaran_limitotal3d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitotal3dd_fullbb_432d   int     `json:"pasaran_limitotal3dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitotal2d_fullbb_432d    int     `json:"pasaran_limitotal2d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitotal2dd_fullbb_432d   int     `json:"pasaran_limitotal2dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitotal2dt_fullbb_432d   int     `json:"pasaran_limitotal2dt_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal4d_fullbb_432d  int     `json:"pasaran_limitglobal4d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal3d_fullbb_432d  int     `json:"pasaran_limitglobal3d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal3dd_fullbb_432d int     `json:"pasaran_limitglobal3dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal2d_fullbb_432d  int     `json:"pasaran_limitglobal2d_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dd_fullbb_432d int     `json:"pasaran_limitglobal2dd_fullbb_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dt_fullbb_432d int     `json:"pasaran_limitglobal2dt_fullbb_432d" validate:"required,numeric"`
	Pasaran_win4d_432d                 int     `json:"pasaran_win4d_432d" validate:"required,numeric"`
	Pasaran_win3d_432d                 int     `json:"pasaran_win3d_432d" validate:"required,numeric"`
	Pasaran_win3dd_432d                int     `json:"pasaran_win3dd_432d" validate:"required,numeric"`
	Pasaran_win2d_432d                 int     `json:"pasaran_win2d_432d" validate:"required,numeric"`
	Pasaran_win2dd_432d                int     `json:"pasaran_win2dd_432d" validate:"required,numeric"`
	Pasaran_win2dt_432d                int     `json:"pasaran_win2dt_432d" validate:"required,numeric"`
	Pasaran_win4dnodisc_432d           int     `json:"pasaran_win4dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3dnodisc_432d           int     `json:"pasaran_win3dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3ddnodisc_432d          int     `json:"pasaran_win3ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dnodisc_432d           int     `json:"pasaran_win2dnodisc_432d" validate:"required,numeric"`
	Pasaran_win2ddnodisc_432d          int     `json:"pasaran_win2ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dtnodisc_432d          int     `json:"pasaran_win2dtnodisc_432d" validate:"required,numeric"`
	Pasaran_win4dbb_kena_432d          int     `json:"pasaran_win4dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win3dbb_kena_432d          int     `json:"pasaran_win3dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win3ddbb_kena_432d         int     `json:"pasaran_win3ddbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2dbb_kena_432d          int     `json:"pasaran_win2dbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2ddbb_kena_432d         int     `json:"pasaran_win2ddbb_kena_432d" validate:"required,numeric"`
	Pasaran_win2dtbb_kena_432d         int     `json:"pasaran_win2dtbb_kena_432d" validate:"required,numeric"`
	Pasaran_win4dbb_432d               int     `json:"pasaran_win4dbb_432d" validate:"required,numeric"`
	Pasaran_win3dbb_432d               int     `json:"pasaran_win3dbb_432d" validate:"required,numeric"`
	Pasaran_win3ddbb_432d              int     `json:"pasaran_win3ddbb_432d" validate:"required,numeric"`
	Pasaran_win2dbb_432d               int     `json:"pasaran_win2dbb_432d" validate:"required,numeric"`
	Pasaran_win2ddbb_432d              int     `json:"pasaran_win2ddbb_432d" validate:"required,numeric"`
	Pasaran_win2dtbb_432d              int     `json:"pasaran_win2dtbb_432d" validate:"required,numeric"`
	Pasaran_disc4d_432d                float32 `json:"pasaran_disc4d_432d" validate:"required,numeric"`
	Pasaran_disc3d_432d                float32 `json:"pasaran_disc3d_432d" validate:"required,numeric"`
	Pasaran_disc3dd_432d               float32 `json:"pasaran_disc3dd_432d" validate:"required,numeric"`
	Pasaran_disc2d_432d                float32 `json:"pasaran_disc2d_432d" validate:"required,numeric"`
	Pasaran_disc2dd_432d               float32 `json:"pasaran_disc2dd_432d" validate:"required,numeric"`
	Pasaran_disc2dt_432d               float32 `json:"pasaran_disc2dt_432d" validate:"required,numeric"`
}
type Controller_companypasarancolokbebas struct {
	Company                    string  `json:"company" validate:"required"`
	Master                     string  `json:"master" validate:"required"`
	Pasaran_id                 string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id          int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbuy_cbebas      int     `json:"pasaran_maxbuy_cbebas" validate:"required,numeric"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" validate:"required,numeric"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" validate:"required,numeric"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" validate:"required,numeric"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" validate:"required,numeric"`
}
type Controller_companypasarancolokmacau struct {
	Company                    string  `json:"company" validate:"required"`
	Master                     string  `json:"master" validate:"required"`
	Pasaran_id                 string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id          int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" validate:"required,numeric"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" validate:"required,numeric"`
	Pasaran_maxbuy_cmacau      int     `json:"pasaran_maxbuy_cmacau" validate:"required,numeric"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau" validate:"required,numeric"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" validate:"required,numeric"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau" validate:"required,numeric"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau" validate:"required,numeric"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau" validate:"required,numeric"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau" validate:"required,numeric"`
}
type Controller_companypasarancoloknaga struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" validate:"required,numeric"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" validate:"required,numeric"`
	Pasaran_maxbuy_cnaga      int     `json:"pasaran_maxbuy_cnaga" validate:"required,numeric"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" validate:"required,numeric"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga" validate:"required,numeric"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga" validate:"required,numeric"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga" validate:"required,numeric"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga" validate:"required,numeric"`
}
type Controller_companypasarancolokjitu struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu" validate:"required,numeric"`
	Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu" validate:"required,numeric"`
	Pasaran_maxbuy_cjitu      int     `json:"pasaran_maxbuy_cjitu" validate:"required,numeric"`
	Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu" validate:"required,numeric"`
	Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu" validate:"required,numeric"`
	Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu" validate:"required,numeric"`
	Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu" validate:"required,numeric"`
	Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu" validate:"required,numeric"`
	Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu" validate:"required,numeric"`
	Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu" validate:"required,numeric"`
}
type Controller_companypasaran5050umum struct {
	Company                      string  `json:"company" validate:"required"`
	Master                       string  `json:"master" validate:"required"`
	Pasaran_id                   string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id            int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum" validate:"required,numeric"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum" validate:"required,numeric"`
	Pasaran_maxbuy_5050umum      int     `json:"pasaran_maxbuy_5050umum" validate:"required,numeric"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" validate:"required,numeric"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" validate:"required,numeric"`
	Pasaran_keibesar_5050umum    float64 `json:"pasaran_keibesar_5050umum" validate:"required,numeric"`
	Pasaran_keikecil_5050umum    float64 `json:"pasaran_keikecil_5050umum" validate:"required,numeric"`
	Pasaran_keigenap_5050umum    float64 `json:"pasaran_keigenap_5050umum" validate:"required,numeric"`
	Pasaran_keiganjil_5050umum   float64 `json:"pasaran_keiganjil_5050umum" validate:"required,numeric"`
	Pasaran_keitengah_5050umum   float64 `json:"pasaran_keitengah_5050umum" validate:"required,numeric"`
	Pasaran_keitepi_5050umum     float64 `json:"pasaran_keitepi_5050umum" validate:"required,numeric"`
	Pasaran_discbesar_5050umum   float64 `json:"pasaran_discbesar_5050umum" validate:"numeric"`
	Pasaran_disckecil_5050umum   float64 `json:"pasaran_disckecil_5050umum" validate:"numeric"`
	Pasaran_discgenap_5050umum   float64 `json:"pasaran_discgenap_5050umum" validate:"numeric"`
	Pasaran_discganjil_5050umum  float64 `json:"pasaran_discganjil_5050umum" validate:"numeric"`
	Pasaran_disctengah_5050umum  float64 `json:"pasaran_disctengah_5050umum" validate:"numeric"`
	Pasaran_disctepi_5050umum    float64 `json:"pasaran_disctepi_5050umum" validate:"numeric"`
}
type Controller_companypasaran5050special struct {
	Company                              string  `json:"company" validate:"required"`
	Master                               string  `json:"master" validate:"required"`
	Pasaran_id                           string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id                    int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special" validate:"required,numeric"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special" validate:"required,numeric"`
	Pasaran_maxbuy_5050special           int     `json:"pasaran_maxbuy_5050special" validate:"required,numeric"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special" validate:"required,numeric"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special" validate:"required,numeric"`
	Pasaran_keiasganjil_5050special      float64 `json:"pasaran_keiasganjil_5050special" validate:"required,numeric"`
	Pasaran_keiasgenap_5050special       float64 `json:"pasaran_keiasgenap_5050special" validate:"required,numeric"`
	Pasaran_keiasbesar_5050special       float64 `json:"pasaran_keiasbesar_5050special" validate:"required,numeric"`
	Pasaran_keiaskecil_5050special       float64 `json:"pasaran_keiaskecil_5050special" validate:"required,numeric"`
	Pasaran_keikopganjil_5050special     float64 `json:"pasaran_keikopganjil_5050special" validate:"required,numeric"`
	Pasaran_keikopgenap_5050special      float64 `json:"pasaran_keikopgenap_5050special" validate:"required,numeric"`
	Pasaran_keikopbesar_5050special      float64 `json:"pasaran_keikopbesar_5050special" validate:"required,numeric"`
	Pasaran_keikopkecil_5050special      float64 `json:"pasaran_keikopkecil_5050special" validate:"required,numeric"`
	Pasaran_keikepalaganjil_5050special  float64 `json:"pasaran_keikepalaganjil_5050special" validate:"required,numeric"`
	Pasaran_keikepalagenap_5050special   float64 `json:"pasaran_keikepalagenap_5050special" validate:"required,numeric"`
	Pasaran_keikepalabesar_5050special   float64 `json:"pasaran_keikepalabesar_5050special" validate:"required,numeric"`
	Pasaran_keikepalakecil_5050special   float64 `json:"pasaran_keikepalakecil_5050special" validate:"required,numeric"`
	Pasaran_keiekorganjil_5050special    float64 `json:"pasaran_keiekorganjil_5050special" validate:"required,numeric"`
	Pasaran_keiekorgenap_5050special     float64 `json:"pasaran_keiekorgenap_5050special" validate:"required,numeric"`
	Pasaran_keiekorbesar_5050special     float64 `json:"pasaran_keiekorbesar_5050special" validate:"required,numeric"`
	Pasaran_keiekorkecil_5050special     float64 `json:"pasaran_keiekorkecil_5050special" validate:"required,numeric"`
	Pasaran_discasganjil_5050special     float64 `json:"pasaran_discasganjil_5050special" validate:"numeric"`
	Pasaran_discasgenap_5050special      float64 `json:"pasaran_discasgenap_5050special" validate:"numeric"`
	Pasaran_discasbesar_5050special      float64 `json:"pasaran_discasbesar_5050special" validate:"numeric"`
	Pasaran_discaskecil_5050special      float64 `json:"pasaran_discaskecil_5050special" validate:"numeric"`
	Pasaran_disckopganjil_5050special    float64 `json:"pasaran_disckopganjil_5050special" validate:"numeric"`
	Pasaran_disckopgenap_5050special     float64 `json:"pasaran_disckopgenap_5050special" validate:"numeric"`
	Pasaran_disckopbesar_5050special     float64 `json:"pasaran_disckopbesar_5050special" validate:"numeric"`
	Pasaran_disckopkecil_5050special     float64 `json:"pasaran_disckopkecil_5050special" validate:"numeric"`
	Pasaran_disckepalaganjil_5050special float64 `json:"pasaran_disckepalaganjil_5050special" validate:"numeric"`
	Pasaran_disckepalagenap_5050special  float64 `json:"pasaran_disckepalagenap_5050special" validate:"numeric"`
	Pasaran_disckepalabesar_5050special  float64 `json:"pasaran_disckepalabesar_5050special" validate:"numeric"`
	Pasaran_disckepalakecil_5050special  float64 `json:"pasaran_disckepalakecil_5050special" validate:"numeric"`
	Pasaran_discekorganjil_5050special   float64 `json:"pasaran_discekorganjil_5050special" validate:"numeric"`
	Pasaran_discekorgenap_5050special    float64 `json:"pasaran_discekorgenap_5050special" validate:"numeric"`
	Pasaran_discekorbesar_5050special    float64 `json:"pasaran_discekorbesar_5050special" validate:"numeric"`
	Pasaran_discekorkecil_5050special    float64 `json:"pasaran_discekorkecil_5050special" validate:"numeric"`
}
type Controller_companypasaran5050kombinasi struct {
	Company                                   string  `json:"company" validate:"required"`
	Master                                    string  `json:"master" validate:"required"`
	Pasaran_id                                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id                         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_maxbuy_5050kombinasi              int     `json:"pasaran_maxbuy_5050kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeimono_5050kombinasi     float64 `json:"pasaran_belakangkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeistereo_5050kombinasi   float64 `json:"pasaran_belakangkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembang_5050kombinasi  float64 `json:"pasaran_belakangkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikempis_5050kombinasi   float64 `json:"pasaran_belakangkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembar_5050kombinasi   float64 `json:"pasaran_belakangkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeimono_5050kombinasi       float64 `json:"pasaran_tengahkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeistereo_5050kombinasi     float64 `json:"pasaran_tengahkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembang_5050kombinasi    float64 `json:"pasaran_tengahkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikempis_5050kombinasi     float64 `json:"pasaran_tengahkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembar_5050kombinasi     float64 `json:"pasaran_tengahkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeimono_5050kombinasi        float64 `json:"pasaran_depankeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeistereo_5050kombinasi      float64 `json:"pasaran_depankeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembang_5050kombinasi     float64 `json:"pasaran_depankeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikempis_5050kombinasi      float64 `json:"pasaran_depankeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembar_5050kombinasi      float64 `json:"pasaran_depankeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdiscmono_5050kombinasi    float64 `json:"pasaran_belakangdiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdiscstereo_5050kombinasi  float64 `json:"pasaran_belakangdiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckembang_5050kombinasi float64 `json:"pasaran_belakangdisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckempis_5050kombinasi  float64 `json:"pasaran_belakangdisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_belakangdisckembar_5050kombinasi  float64 `json:"pasaran_belakangdisckembar_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdiscmono_5050kombinasi      float64 `json:"pasaran_tengahdiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdiscstereo_5050kombinasi    float64 `json:"pasaran_tengahdiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckembang_5050kombinasi   float64 `json:"pasaran_tengahdisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckempis_5050kombinasi    float64 `json:"pasaran_tengahdisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_tengahdisckembar_5050kombinasi    float64 `json:"pasaran_tengahdisckembar_5050kombinasi" validate:"numeric"`
	Pasaran_depandiscmono_5050kombinasi       float64 `json:"pasaran_depandiscmono_5050kombinasi" validate:"numeric"`
	Pasaran_depandiscstereo_5050kombinasi     float64 `json:"pasaran_depandiscstereo_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckembang_5050kombinasi    float64 `json:"pasaran_depandisckembang_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckempis_5050kombinasi     float64 `json:"pasaran_depandisckempis_5050kombinasi" validate:"numeric"`
	Pasaran_depandisckembar_5050kombinasi     float64 `json:"pasaran_depandisckembar_5050kombinasi" validate:"numeric"`
}

type Controller_companypasarankombinasi struct {
	Company                       string  `json:"company" validate:"required"`
	Master                        string  `json:"master" validate:"required"`
	Pasaran_id                    string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id             int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbuy_kombinasi      int     `json:"pasaran_maxbuy_kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" validate:"required,numeric"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" validate:"required,numeric"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" validate:"required,numeric"`
}
type Controller_companypasarandasar struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar" validate:"required,numeric"`
	Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar" validate:"required,numeric"`
	Pasaran_maxbuy_dasar      int     `json:"pasaran_maxbuy_dasar" validate:"required,numeric"`
	Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar" validate:"required,numeric"`
	Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar" validate:"required,numeric"`
	Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar" validate:"numeric"`
	Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar" validate:"numeric"`
	Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar" validate:"numeric"`
	Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar" validate:"numeric"`
	Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar" validate:"numeric"`
	Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar" validate:"numeric"`
	Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar" validate:"numeric"`
	Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar" validate:"numeric"`
}
type Controller_companypasaranshio struct {
	Company                  string  `json:"company" validate:"required"`
	Master                   string  `json:"master" validate:"required"`
	Pasaran_id               string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id        int     `json:"companypasaran_id" validate:"required"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio" validate:"required"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio" validate:"required,numeric"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio" validate:"required,numeric"`
	Pasaran_maxbuy_shio      int     `json:"pasaran_maxbuy_shio" validate:"required,numeric"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio" validate:"required,numeric"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio" validate:"required,numeric"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio" validate:"numeric"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio" validate:"required,numeric"`
}
type Controller_companydetailonlinesave struct {
	Company           string `json:"company" validate:"required"`
	Master            string `json:"master" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
	Pasaran_hari      string `json:"pasaran_hari" validate:"required"`
}
type Controller_companydeletepasaranonline struct {
	Company              string `json:"company" validate:"required"`
	Master               string `json:"master" validate:"required"`
	Companypasaran_id    int    `json:"companypasaran_id" validate:"required"`
	Companypasaran_idoff int    `json:"companypasaran_idoffline" validate:"required"`
}
type Controller_companyfetchpasaran432 struct {
	Company           string `json:"company" validate:"required"`
	Master            string `json:"master" validate:"required"`
	Pasaran_id        string `json:"pasaran_id" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
}
type Controller_companyinvoicelistpermainan struct {
	Company   string `json:"company" validate:"required"`
	Invoice   int    `json:"invoice" validate:"required"`
	Permainan string `json:"permainan" validate:"required"`
}
type Controller_companyinvoicelistpermainanstatus struct {
	Company string `json:"company" validate:"required"`
	Invoice int    `json:"invoice" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
type Controller_companyinvoicelistpermainanusername struct {
	Company   string `json:"company" validate:"required"`
	Invoice   int    `json:"invoice" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Permainan string `json:"permainan" validate:"required"`
}
