package entities

type Model_pasaran struct {
	Pasaran_idpasarantogel string `json:"pasaran_idpasarantogel"`
	Pasaran_nmpasarantogel string `json:"pasaran_nmpasarantogel"`
	Pasaran_tipepasaran    string `json:"pasaran_tipepasaran"`
	Pasaran_urlpasaran     string `json:"pasaran_urlpasaran"`
	Pasaran_pasarandiundi  string `json:"pasaran_pasarandiundi"`
	Pasaran_jamtutup       string `json:"pasaran_jamtutup"`
	Pasaran_jamjadwal      string `json:"pasaran_jamjadwal"`
	Pasaran_jamopen        string `json:"pasaran_jamopen"`
	Pasaran_create         string `json:"pasaran_create"`
	Pasaran_update         string `json:"pasaran_update"`
}

type Model_pasaranDetailConf struct {
	Limitline4d                       int     `json:"limitline_4d"`
	Limitline3d                       int     `json:"limitline_3d"`
	Limitline3dd                      int     `json:"limitline_3dd"`
	Limitline2d                       int     `json:"limitline_2d"`
	Limitline2dd                      int     `json:"limitline_2dd"`
	Limitline2dt                      int     `json:"limitline_2dt"`
	Bbfs                              int     `json:"bbfs"`
	Minbet_432d                       float32 `json:"minbet_432d"`
	Maxbet4d_432d                     float32 `json:"maxbet4d_432d"`
	Maxbet3d_432d                     float32 `json:"maxbet3d_432d"`
	Maxbet3dd_432d                    float32 `json:"maxbet3dd_432d"`
	Maxbet2d_432d                     float32 `json:"maxbet2d_432d"`
	Maxbet2dd_432d                    float32 `json:"maxbet2dd_432d"`
	Maxbet2dt_432d                    float32 `json:"maxbet2dt_432d"`
	Limitotal4d_432d                  float32 `json:"limitotal4d_432d"`
	Limitotal3d_432d                  float32 `json:"limitotal3d_432d"`
	Limitotal3dd_432d                 float32 `json:"limitotal3dd_432d"`
	Limitotal2d_432d                  float32 `json:"limitotal2d_432d"`
	Limitotal2dd_432d                 float32 `json:"limitotal2dd_432d"`
	Limitotal2dt_432d                 float32 `json:"limitotal2dt_432d"`
	Limitglobal4d_432d                float32 `json:"limitglobal4d_432d"`
	Limitglobal3d_432d                float32 `json:"limitglobal3d_432d"`
	Limitglobal3dd_432d               float32 `json:"limitglobal3dd_432d"`
	Limitglobal2d_432d                float32 `json:"limitglobal2d_432d"`
	Limitglobal2dd_432d               float32 `json:"limitglobal2dd_432d"`
	Limitglobal2dt_432d               float32 `json:"limitglobal2dt_432d"`
	Disc4d_432d                       float32 `json:"disc4d_432d"`
	Disc3d_432d                       float32 `json:"disc3d_432d"`
	Disc3dd_432d                      float32 `json:"disc3dd_432d"`
	Disc2d_432d                       float32 `json:"disc2d_432d"`
	Disc2dd_432d                      float32 `json:"disc2dd_432d"`
	Disc2dt_432d                      float32 `json:"disc2dt_432d"`
	Win4d_432d                        float32 `json:"win4d_432d"`
	Win3d_432d                        float32 `json:"win3d_432d"`
	Win3dd_432d                       float32 `json:"win3dd_432d"`
	Win2d_432d                        float32 `json:"win2d_432d"`
	Win2dd_432d                       float32 `json:"win2dd_432d"`
	Win2dt_432d                       float32 `json:"win2dt_432d"`
	Win4dnodisc_432d                  float32 `json:"win4dnodisc_432d"`
	Win3dnodisc_432d                  float32 `json:"win3dnodisc_432d"`
	Win3ddnodisc_432d                 float32 `json:"win3ddnodisc_432d"`
	Win2dnodisc_432d                  float32 `json:"win2dnodisc_432d"`
	Win2ddnodisc_432d                 float32 `json:"win2ddnodisc_432d"`
	Win2dtnodisc_432d                 float32 `json:"win2dtnodisc_432d"`
	Minbet_cbebas                     float32 `json:"minbet_cbebas"`
	Maxbet_cbebas                     float32 `json:"maxbet_cbebas"`
	Win_cbebas                        float32 `json:"win_cbebas"`
	Disc_cbebas                       float32 `json:"disc_cbebas"`
	Limitglobal_cbebas                float32 `json:"limitglobal_cbebas"`
	Limittotal_cbebas                 float32 `json:"limittotal_cbebas"`
	Minbet_cmacau                     float32 `json:"minbet_cmacau"`
	Maxbet_cmacau                     float32 `json:"maxbet_cmacau"`
	Win2d_cmacau                      float32 `json:"win2d_cmacau"`
	Win3d_cmacau                      float32 `json:"win3d_cmacau"`
	Win4d_cmacau                      float32 `json:"win4d_cmacau"`
	Disc_cmacau                       float32 `json:"disc_cmacau"`
	Limitglobal_cmacau                float32 `json:"limitglobal_cmacau"`
	Limitotal_cmacau                  float32 `json:"limitotal_cmacau"`
	Minbet_cnaga                      float32 `json:"minbet_cnaga"`
	Maxbet_cnaga                      float32 `json:"maxbet_cnaga"`
	Win3_cnaga                        float32 `json:"win3_cnaga"`
	Win4_cnaga                        float32 `json:"win4_cnaga"`
	Disc_cnaga                        float32 `json:"disc_cnaga"`
	Limitglobal_cnaga                 float32 `json:"limitglobal_cnaga"`
	Limittotal_cnaga                  float32 `json:"limittotal_cnaga"`
	Minbet_cjitu                      float32 `json:"minbet_cjitu"`
	Maxbet_cjitu                      float32 `json:"maxbet_cjitu"`
	Winas_cjitu                       float32 `json:"winas_cjitu"`
	Winkop_cjitu                      float32 `json:"winkop_cjitu"`
	Winkepala_cjitu                   float32 `json:"winkepala_cjitu"`
	Winekor_cjitu                     float32 `json:"winekor_cjitu"`
	Desc_cjitu                        float32 `json:"desc_cjitu"`
	Limitglobal_cjitu                 float32 `json:"limitglobal_cjitu"`
	Limittotal_cjitu                  float32 `json:"limittotal_cjitu"`
	Minbet_5050umum                   float32 `json:"minbet_5050umum"`
	Maxbet_5050umum                   float32 `json:"maxbet_5050umum"`
	Keibesar_5050umum                 float32 `json:"keibesar_5050umum"`
	Keikecil_5050umum                 float32 `json:"keikecil_5050umum"`
	Keigenap_5050umum                 float32 `json:"keigenap_5050umum"`
	Keiganjil_5050umum                float32 `json:"keiganjil_5050umum"`
	Keitengah_5050umum                float32 `json:"keitengah_5050umum"`
	Keitepi_5050umum                  float32 `json:"keitepi_5050umum"`
	Discbesar_5050umum                float32 `json:"discbesar_5050umum"`
	Disckecil_5050umum                float32 `json:"disckecil_5050umum"`
	Discgenap_5050umum                float32 `json:"discgenap_5050umum"`
	Discganjil_5050umum               float32 `json:"discganjil_5050umum"`
	Disctengah_5050umum               float32 `json:"disctengah_5050umum"`
	Disctepi_5050umum                 float32 `json:"disctepi_5050umum"`
	Limitglobal_5050umum              float32 `json:"limitglobal_5050umum"`
	Limittotal_5050umum               float32 `json:"limittotal_5050umum"`
	Minbet_5050special                float32 `json:"minbet_5050special"`
	Maxbet_5050special                float32 `json:"maxbet_5050special"`
	Keiasganjil_5050special           float32 `json:"keiasganjil_5050special"`
	Keiasgenap_5050special            float32 `json:"keiasgenap_5050special"`
	Keiasbesar_5050special            float32 `json:"keiasbesar_5050special"`
	Keiaskecil_5050special            float32 `json:"keiaskecil_5050special"`
	Keikopganjil_5050special          float32 `json:"keikopganjil_5050special"`
	Keikopgenap_5050special           float32 `json:"keikopgenap_5050special"`
	Keikopbesar_5050special           float32 `json:"keikopbesar_5050special"`
	Keikopkecil_5050special           float32 `json:"keikopkecil_5050special"`
	Keikepalaganjil_5050special       float32 `json:"keikepalaganjil_5050special"`
	Keikepalagenap_5050special        float32 `json:"keikepalagenap_5050special"`
	Keikepalabesar_5050special        float32 `json:"keikepalabesar_5050special"`
	Keikepalakecil_5050special        float32 `json:"keikepalakecil_5050special"`
	Keiekorganjil_5050special         float32 `json:"keiekorganjil_5050special"`
	Keiekorgenap_5050special          float32 `json:"keiekorgenap_5050special"`
	Keiekorbesar_5050special          float32 `json:"keiekorbesar_5050special"`
	Keiekorkecil_5050special          float32 `json:"keiekorkecil_5050special"`
	Discasganjil_5050special          float32 `json:"discasganjil_5050special"`
	Discasgenap_5050special           float32 `json:"discasgenap_5050special"`
	Discasbesar_5050special           float32 `json:"discasbesar_5050special"`
	Discaskecil_5050special           float32 `json:"discaskecil_5050special"`
	Disckopganjil_5050special         float32 `json:"disckopganjil_5050special"`
	Disckopgenap_5050special          float32 `json:"disckopgenap_5050special"`
	Disckopbesar_5050special          float32 `json:"disckopbesar_5050special"`
	Disckopkecil_5050special          float32 `json:"disckopkecil_5050special"`
	Disckepalaganjil_5050special      float32 `json:"disckepalaganjil_5050special"`
	Disckepalagenap_5050special       float32 `json:"disckepalagenap_5050special"`
	Disckepalabesar_5050special       float32 `json:"disckepalabesar_5050special"`
	Disckepalakecil_5050special       float32 `json:"disckepalakecil_5050special"`
	Discekorganjil_5050special        float32 `json:"discekorganjil_5050special"`
	Discekorgenap_5050special         float32 `json:"discekorgenap_5050special"`
	Discekorbesar_5050special         float32 `json:"discekorbesar_5050special"`
	Discekorkecil_5050special         float32 `json:"discekorkecil_5050special"`
	Limitglobal_5050special           float32 `json:"limitglobal_5050special"`
	Limittotal_5050special            float32 `json:"limittotal_5050special"`
	Minbet_5050kombinasi              float32 `json:"minbet_5050kombinasi"`
	Maxbet_5050kombinasi              float32 `json:"maxbet_5050kombinasi"`
	Belakangkeimono_5050kombinasi     float32 `json:"belakangkeimono_5050kombinasi"`
	Belakangkeistereo_5050kombinasi   float32 `json:"belakangkeistereo_5050kombinasi"`
	Belakangkeikembang_5050kombinasi  float32 `json:"belakangkeikembang_5050kombinasi"`
	Belakangkeikempis_5050kombinasi   float32 `json:"belakangkeikempis_5050kombinasi"`
	Belakangkeikembar_5050kombinasi   float32 `json:"belakangkeikembar_5050kombinasi"`
	Tengahkeimono_5050kombinasi       float32 `json:"tengahkeimono_5050kombinasi"`
	Tengahkeistereo_5050kombinasi     float32 `json:"tengahkeistereo_5050kombinasi"`
	Tengahkeikembang_5050kombinasi    float32 `json:"tengahkeikembang_5050kombinasi"`
	Tengahkeikempis_5050kombinasi     float32 `json:"tengahkeikempis_5050kombinasi"`
	Tengahkeikembar_5050kombinasi     float32 `json:"tengahkeikembar_5050kombinasi"`
	Depankeimono_5050kombinasi        float32 `json:"depankeimono_5050kombinasi"`
	Depankeistereo_5050kombinasi      float32 `json:"depankeistereo_5050kombinasi"`
	Depankeikembang_5050kombinasi     float32 `json:"depankeikembang_5050kombinasi"`
	Depankeikempis_5050kombinasi      float32 `json:"depankeikempis_5050kombinasi"`
	Depankeikembar_5050kombinasi      float32 `json:"depankeikembar_5050kombinasi"`
	Belakangdiscmono_5050kombinasi    float32 `json:"belakangdiscmono_5050kombinasi"`
	Belakangdiscstereo_5050kombinasi  float32 `json:"belakangdiscstereo_5050kombinasi"`
	Belakangdisckembang_5050kombinasi float32 `json:"belakangdisckembang_5050kombinasi"`
	Belakangdisckempis_5050kombinasi  float32 `json:"belakangdisckempis_5050kombinasi"`
	Belakangdisckembar_5050kombinasi  float32 `json:"belakangdisckembar_5050kombinasi"`
	Tengahdiscmono_5050kombinasi      float32 `json:"tengahdiscmono_5050kombinasi"`
	Tengahdiscstereo_5050kombinasi    float32 `json:"tengahdiscstereo_5050kombinasi"`
	Tengahdisckembang_5050kombinasi   float32 `json:"tengahdisckembang_5050kombinasi"`
	Tengahdisckempis_5050kombinasi    float32 `json:"tengahdisckempis_5050kombinasi"`
	Tengahdisckembar_5050kombinasi    float32 `json:"tengahdisckembar_5050kombinasi"`
	Depandiscmono_5050kombinasi       float32 `json:"depandiscmono_5050kombinasi"`
	Depandiscstereo_5050kombinasi     float32 `json:"depandiscstereo_5050kombinasi"`
	Depandisckembang_5050kombinasi    float32 `json:"depandisckembang_5050kombinasi"`
	Depandisckempis_5050kombinasi     float32 `json:"depandisckempis_5050kombinasi"`
	Depandisckembar_5050kombinasi     float32 `json:"depandisckembar_5050kombinasi"`
	Limitglobal_5050kombinasi         float32 `json:"limitglobal_5050kombinasi"`
	Limittotal_5050kombinasi          float32 `json:"limittotal_5050kombinasi"`
	Minbet_kombinasi                  float32 `json:"minbet_kombinasi"`
	Maxbet_kombinasi                  float32 `json:"maxbet_kombinasi"`
	Win_kombinasi                     float32 `json:"win_kombinasi"`
	Disc_kombinasi                    float32 `json:"disc_kombinasi"`
	Limitglobal_kombinasi             float32 `json:"limitglobal_kombinasi"`
	Limittotal_kombinasi              float32 `json:"limittotal_kombinasi"`
	Minbet_dasar                      float32 `json:"minbet_dasar"`
	Maxbet_dasar                      float32 `json:"maxbet_dasar"`
	Keibesar_dasar                    float32 `json:"keibesar_dasar"`
	Keikecil_dasar                    float32 `json:"keikecil_dasar"`
	Keigenap_dasar                    float32 `json:"keigenap_dasar"`
	Keiganjil_dasar                   float32 `json:"keiganjil_dasar"`
	Discbesar_dasar                   float32 `json:"discbesar_dasar"`
	Disckecil_dasar                   float32 `json:"disckecil_dasar"`
	Discgenap_dasar                   float32 `json:"discgenap_dasar"`
	Discganjil_dasar                  float32 `json:"discganjil_dasar"`
	Limitglobal_dasar                 float32 `json:"limitglobal_dasar"`
	Limittotal_dasar                  float32 `json:"limittotal_dasar"`
	Minbet_shio                       float32 `json:"minbet_shio"`
	Maxbet_shio                       float32 `json:"maxbet_shio"`
	Win_shio                          float32 `json:"win_shio"`
	Disc_shio                         float32 `json:"disc_shio"`
	Shioyear_shio                     string  `json:"shioyear_shio"`
	Limitglobal_shio                  float32 `json:"limitglobal_shio"`
	Limittotal_shio                   float32 `json:"limittotal_shio"`
}
type Controller_pasaran struct {
	Master string `json:"master" validate:"required"`
}
type Controller_pasarandetail struct {
	Pasarancode string `json:"pasarancode" validate:"required"`
	Master      string `json:"master" validate:"required"`
}
type Controller_pasaransave struct {
	Sdata     string `json:"sdata" validate:"required"`
	Master    string `json:"master" validate:"required"`
	Idrecord  string `json:"idrecord" validate:"required,min=2,max=10"`
	Name      string `json:"pasaran_name" validate:"required"`
	Diundi    string `json:"pasaran_diundi" validate:"required"`
	Url       string `json:"pasaran_url" validate:"required"`
	Jamtutup  string `json:"pasaran_jamtutup" validate:"required"`
	Jamjadwal string `json:"pasaran_jamjadwal" validate:"required"`
	Jamopen   string `json:"pasaran_jamopen" validate:"required"`
	Tipe      string `json:"pasaran_tipe" `
}
type Controller_pasaransavelimitline struct {
	Master               string `json:"master" validate:"required"`
	Idrecord             string `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_bbfs         int    `json:"pasaran_bbfs" validate:"required,numeric"`
	Pasaran_limitline4d  int    `json:"pasaran_limitline4d" validate:"required,numeric"`
	Pasaran_limitline3d  int    `json:"pasaran_limitline3d" validate:"required,numeric"`
	Pasaran_limitline3dd int    `json:"pasaran_limitline3dd" validate:"required,numeric"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d" validate:"required,numeric"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd" validate:"required,numeric"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt" validate:"required,numeric"`
}
type Controller_pasaransaveconf432d struct {
	Master                      string  `json:"master" validate:"required"`
	Idrecord                    string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d" validate:"required,numeric"`
	Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d" validate:"required,numeric"`
	Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d" validate:"required,numeric"`
	Pasaran_maxbet3dd_432d      int     `json:"pasaran_maxbet3dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d" validate:"required,numeric"`
	Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d" validate:"required,numeric"`
	Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d" validate:"required,numeric"`
	Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d" validate:"required,numeric"`
	Pasaran_limitotal3dd_432d   int     `json:"pasaran_limitotal3dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d" validate:"required,numeric"`
	Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d" validate:"required,numeric"`
	Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3dd_432d int     `json:"pasaran_limitglobal3dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d" validate:"required,numeric"`
	Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d" validate:"required,numeric"`
	Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d" validate:"required,numeric"`
	Pasaran_win3dd_432d         int     `json:"pasaran_win3dd_432d" validate:"required,numeric"`
	Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d" validate:"required,numeric"`
	Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d" validate:"required,numeric"`
	Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d" validate:"required,numeric"`
	Pasaran_win4dnodisc_432d    int     `json:"pasaran_win4dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3dnodisc_432d    int     `json:"pasaran_win3dnodisc_432d" validate:"required,numeric"`
	Pasaran_win3ddnodisc_432d   int     `json:"pasaran_win3ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dnodisc_432d    int     `json:"pasaran_win2dnodisc_432d" validate:"required,numeric"`
	Pasaran_win2ddnodisc_432d   int     `json:"pasaran_win2ddnodisc_432d" validate:"required,numeric"`
	Pasaran_win2dtnodisc_432d   int     `json:"pasaran_win2dtnodisc_432d" validate:"required,numeric"`
	Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d" validate:"required,numeric"`
	Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d" validate:"required,numeric"`
	Pasaran_disc3dd_432d        float32 `json:"pasaran_disc3dd_432d" validate:"required,numeric"`
	Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d" validate:"required,numeric"`
	Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d" validate:"required,numeric"`
	Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d" validate:"required,numeric"`
}
type Controller_pasaransaveconfcolokbebas struct {
	Master                     string  `json:"master" validate:"required"`
	Idrecord                   string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" validate:"required,numeric"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" validate:"required,numeric"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" validate:"required,numeric"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" validate:"required,numeric"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" validate:"required,numeric"`
}
type Controller_pasaransaveconfcolokmacau struct {
	Master                     string  `json:"master" validate:"required"`
	Idrecord                   string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" validate:"required,numeric"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" validate:"required,numeric"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau" validate:"required,numeric"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" validate:"required,numeric"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau" validate:"required,numeric"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau" validate:"required,numeric"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau" validate:"required,numeric"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau" validate:"required,numeric"`
}
type Controller_pasaransaveconfcoloknaga struct {
	Master                    string  `json:"master" validate:"required"`
	Idrecord                  string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" validate:"required,numeric"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" validate:"required,numeric"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" validate:"required,numeric"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga" validate:"required,numeric"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga" validate:"required,numeric"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga" validate:"required,numeric"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga" validate:"required,numeric"`
}
type Controller_pasaransaveconfcolokjitu struct {
	Master                    string  `json:"master" validate:"required"`
	Idrecord                  string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_cjitu      int     `json:"pasaran_minbet_cjitu" validate:"required,numeric"`
	Pasaran_maxbet_cjitu      int     `json:"pasaran_maxbet_cjitu" validate:"required,numeric"`
	Pasaran_limittotal_cjitu  int     `json:"pasaran_limittotal_cjitu" validate:"required,numeric"`
	Pasaran_limitglobal_cjitu int     `json:"pasaran_limitglobal_cjitu" validate:"required,numeric"`
	Pasaran_winas_cjitu       float32 `json:"pasaran_winas_cjitu" validate:"required,numeric"`
	Pasaran_winkop_cjitu      float32 `json:"pasaran_winkop_cjitu" validate:"required,numeric"`
	Pasaran_winkepala_cjitu   float32 `json:"pasaran_winkepala_cjitu" validate:"required,numeric"`
	Pasaran_winekor_cjitu     float32 `json:"pasaran_winekor_cjitu" validate:"required,numeric"`
	Pasaran_desc_cjitu        float32 `json:"pasaran_desc_cjitu" validate:"required,numeric"`
}
type Controller_pasaransaveconf5050umum struct {
	Master                       string  `json:"master" validate:"required"`
	Idrecord                     string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum" validate:"required,numeric"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum" validate:"required,numeric"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" validate:"required,numeric"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" validate:"required,numeric"`
	Pasaran_keibesar_5050umum    float32 `json:"pasaran_keibesar_5050umum" validate:"required,numeric"`
	Pasaran_keikecil_5050umum    float32 `json:"pasaran_keikecil_5050umum" validate:"required,numeric"`
	Pasaran_keigenap_5050umum    float32 `json:"pasaran_keigenap_5050umum" validate:"required,numeric"`
	Pasaran_keiganjil_5050umum   float32 `json:"pasaran_keiganjil_5050umum" validate:"required,numeric"`
	Pasaran_keitengah_5050umum   float32 `json:"pasaran_keitengah_5050umum" validate:"required,numeric"`
	Pasaran_keitepi_5050umum     float32 `json:"pasaran_keitepi_5050umum" validate:"required,numeric"`
	Pasaran_discbesar_5050umum   float32 `json:"pasaran_discbesar_5050umum" validate:"required,numeric"`
	Pasaran_disckecil_5050umum   float32 `json:"pasaran_disckecil_5050umum" validate:"required,numeric"`
	Pasaran_discgenap_5050umum   float32 `json:"pasaran_discgenap_5050umum" validate:"required,numeric"`
	Pasaran_discganjil_5050umum  float32 `json:"pasaran_discganjil_5050umum" validate:"required,numeric"`
	Pasaran_disctengah_5050umum  float32 `json:"pasaran_disctengah_5050umum" validate:"required,numeric"`
	Pasaran_disctepi_5050umum    float32 `json:"pasaran_disctepi_5050umum" validate:"required,numeric"`
}
type Controller_pasaransaveconf5050special struct {
	Master                               string  `json:"master" validate:"required"`
	Idrecord                             string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special" validate:"required,numeric"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special" validate:"required,numeric"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special" validate:"required,numeric"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special" validate:"required,numeric"`
	Pasaran_keiasganjil_5050special      float32 `json:"pasaran_keiasganjil_5050special" validate:"numeric"`
	Pasaran_keiasgenap_5050special       float32 `json:"pasaran_keiasgenap_5050special" validate:"numeric"`
	Pasaran_keiasbesar_5050special       float32 `json:"pasaran_keiasbesar_5050special" validate:"numeric"`
	Pasaran_keiaskecil_5050special       float32 `json:"pasaran_keiaskecil_5050special" validate:"numeric"`
	Pasaran_keikopganjil_5050special     float32 `json:"pasaran_keikopganjil_5050special" validate:"numeric"`
	Pasaran_keikopgenap_5050special      float32 `json:"pasaran_keikopgenap_5050special" validate:"numeric"`
	Pasaran_keikopbesar_5050special      float32 `json:"pasaran_keikopbesar_5050special" validate:"numeric"`
	Pasaran_keikopkecil_5050special      float32 `json:"pasaran_keikopkecil_5050special" validate:"numeric"`
	Pasaran_keikepalaganjil_5050special  float32 `json:"pasaran_keikepalaganjil_5050special" validate:"numeric"`
	Pasaran_keikepalagenap_5050special   float32 `json:"pasaran_keikepalagenap_5050special" validate:"numeric"`
	Pasaran_keikepalabesar_5050special   float32 `json:"pasaran_keikepalabesar_5050special" validate:"numeric"`
	Pasaran_keikepalakecil_5050special   float32 `json:"pasaran_keikepalakecil_5050special" validate:"numeric"`
	Pasaran_keiekorganjil_5050special    float32 `json:"pasaran_keiekorganjil_5050special" validate:"numeric"`
	Pasaran_keiekorgenap_5050special     float32 `json:"pasaran_keiekorgenap_5050special" validate:"numeric"`
	Pasaran_keiekorbesar_5050special     float32 `json:"pasaran_keiekorbesar_5050special" validate:"numeric"`
	Pasaran_keiekorkecil_5050special     float32 `json:"pasaran_keiekorkecil_5050special" validate:"numeric"`
	Pasaran_discasganjil_5050special     float32 `json:"pasaran_discasganjil_5050special" validate:"numeric"`
	Pasaran_discasgenap_5050special      float32 `json:"pasaran_discasgenap_5050special" validate:"numeric"`
	Pasaran_discasbesar_5050special      float32 `json:"pasaran_discasbesar_5050special" validate:"numeric"`
	Pasaran_discaskecil_5050special      float32 `json:"pasaran_discaskecil_5050special" validate:"numeric"`
	Pasaran_disckopganjil_5050special    float32 `json:"pasaran_disckopganjil_5050special" validate:"numeric"`
	Pasaran_disckopgenap_5050special     float32 `json:"pasaran_disckopgenap_5050special" validate:"numeric"`
	Pasaran_disckopbesar_5050special     float32 `json:"pasaran_disckopbesar_5050special" validate:"numeric"`
	Pasaran_disckopkecil_5050special     float32 `json:"pasaran_disckopkecil_5050special" validate:"numeric"`
	Pasaran_disckepalaganjil_5050special float32 `json:"pasaran_disckepalaganjil_5050special" validate:"numeric"`
	Pasaran_disckepalagenap_5050special  float32 `json:"pasaran_disckepalagenap_5050special" validate:"numeric"`
	Pasaran_disckepalabesar_5050special  float32 `json:"pasaran_disckepalabesar_5050special" validate:"numeric"`
	Pasaran_disckepalakecil_5050special  float32 `json:"pasaran_disckepalakecil_5050special" validate:"numeric"`
	Pasaran_discekorganjil_5050special   float32 `json:"pasaran_discekorganjil_5050special" validate:"numeric"`
	Pasaran_discekorgenap_5050special    float32 `json:"pasaran_discekorgenap_5050special" validate:"numeric"`
	Pasaran_discekorbesar_5050special    float32 `json:"pasaran_discekorbesar_5050special" validate:"numeric"`
	Pasaran_discekorkecil_5050special    float32 `json:"pasaran_discekorkecil_5050special" validate:"numeric"`
}
type Controller_pasaransaveconf5050kombinasi struct {
	Master                                    string  `json:"master" validate:"required"`
	Idrecord                                  string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_5050kombinasi         int     `json:"pasaran_limitglobal_5050kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_5050kombinasi          int     `json:"pasaran_limittotal_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeimono_5050kombinasi     float32 `json:"pasaran_belakangkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeistereo_5050kombinasi   float32 `json:"pasaran_belakangkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembang_5050kombinasi  float32 `json:"pasaran_belakangkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikempis_5050kombinasi   float32 `json:"pasaran_belakangkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangkeikembar_5050kombinasi   float32 `json:"pasaran_belakangkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeimono_5050kombinasi       float32 `json:"pasaran_tengahkeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeistereo_5050kombinasi     float32 `json:"pasaran_tengahkeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembang_5050kombinasi    float32 `json:"pasaran_tengahkeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikempis_5050kombinasi     float32 `json:"pasaran_tengahkeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahkeikembar_5050kombinasi     float32 `json:"pasaran_tengahkeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeimono_5050kombinasi        float32 `json:"pasaran_depankeimono_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeistereo_5050kombinasi      float32 `json:"pasaran_depankeistereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembang_5050kombinasi     float32 `json:"pasaran_depankeikembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikempis_5050kombinasi      float32 `json:"pasaran_depankeikempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_depankeikembar_5050kombinasi      float32 `json:"pasaran_depankeikembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdiscmono_5050kombinasi    float32 `json:"pasaran_belakangdiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdiscstereo_5050kombinasi  float32 `json:"pasaran_belakangdiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckembang_5050kombinasi float32 `json:"pasaran_belakangdisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckempis_5050kombinasi  float32 `json:"pasaran_belakangdisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckembar_5050kombinasi  float32 `json:"pasaran_belakangdisckembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdiscmono_5050kombinasi      float32 `json:"pasaran_tengahdiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdiscstereo_5050kombinasi    float32 `json:"pasaran_tengahdiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckembang_5050kombinasi   float32 `json:"pasaran_tengahdisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckempis_5050kombinasi    float32 `json:"pasaran_tengahdisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckembar_5050kombinasi    float32 `json:"pasaran_tengahdisckembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandiscmono_5050kombinasi       float32 `json:"pasaran_depandiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandiscstereo_5050kombinasi     float32 `json:"pasaran_depandiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckembang_5050kombinasi    float32 `json:"pasaran_depandisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckempis_5050kombinasi     float32 `json:"pasaran_depandisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckembar_5050kombinasi     float32 `json:"pasaran_depandisckembar_5050kombinasi" validate:"required,numeric"`
}
type Controller_pasaransaveconfmacaukombinasi struct {
	Master                        string  `json:"master" validate:"required"`
	Idrecord                      string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" validate:"required,numeric"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" validate:"required,numeric"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" validate:"required,numeric"`
}
type Controller_pasaransaveconfdasar struct {
	Master                    string  `json:"master" validate:"required"`
	Idrecord                  string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_dasar      int     `json:"pasaran_minbet_dasar" validate:"required,numeric"`
	Pasaran_maxbet_dasar      int     `json:"pasaran_maxbet_dasar" validate:"required,numeric"`
	Pasaran_limitglobal_dasar int     `json:"pasaran_limitglobal_dasar" validate:"required,numeric"`
	Pasaran_limittotal_dasar  int     `json:"pasaran_limittotal_dasar" validate:"required,numeric"`
	Pasaran_keibesar_dasar    float32 `json:"pasaran_keibesar_dasar" validate:"numeric"`
	Pasaran_keikecil_dasar    float32 `json:"pasaran_keikecil_dasar" validate:"numeric"`
	Pasaran_keigenap_dasar    float32 `json:"pasaran_keigenap_dasar" validate:"numeric"`
	Pasaran_keiganjil_dasar   float32 `json:"pasaran_keiganjil_dasar" validate:"numeric"`
	Pasaran_discbesar_dasar   float32 `json:"pasaran_discbesar_dasar" validate:"required,numeric"`
	Pasaran_disckecil_dasar   float32 `json:"pasaran_disckecil_dasar" validate:"required,numeric"`
	Pasaran_discgenap_dasar   float32 `json:"pasaran_discgenap_dasar" validate:"required,numeric"`
	Pasaran_discganjil_dasar  float32 `json:"pasaran_discganjil_dasar" validate:"required,numeric"`
}
type Controller_pasaransaveconfshio struct {
	Master                   string  `json:"master" validate:"required"`
	Idrecord                 string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio" validate:"required"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio" validate:"required,numeric"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio" validate:"required,numeric"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio" validate:"required,numeric"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio" validate:"required,numeric"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio" validate:"numeric"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio" validate:"numeric"`
}
