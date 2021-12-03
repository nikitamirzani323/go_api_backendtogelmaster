package controller

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
)

var ctx = context.Background()

type companyhome struct {
	Master string `json:"master" validate:"required"`
}
type companydetail struct {
	Company string `json:"company" validate:"required"`
	Master  string `json:"master" validate:"required"`
}
type companydetailonline struct {
	Company           string `json:"company" validate:"required"`
	Master            string `json:"master" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
}
type companydetailonlinesave struct {
	Company           string `json:"company" validate:"required"`
	Master            string `json:"master" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
	Pasaran_hari      string `json:"pasaran_hari" validate:"required"`
}
type companydetailonlinedelete struct {
	Company              string `json:"company" validate:"required"`
	Master               string `json:"master" validate:"required"`
	Companypasaran_id    int    `json:"companypasaran_id" validate:"required"`
	Companypasaran_idoff int    `json:"companypasaran_idoffline" validate:"required"`
}
type companypasaranconf struct {
	Company           string `json:"company" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
}
type companylistkeluaran struct {
	Company string `json:"company" validate:"required"`
	Periode string `json:"periode" validate:"required"`
	Pasaran int    `json:"pasaran" validate:"required"`
}
type companyinvoice struct {
	Company  string `json:"company" validate:"required"`
	Username string `json:"username" `
	Invoice  int    `json:"invoice" validate:"required"`
}
type companyinvoicelistpermainan struct {
	Company   string `json:"company" validate:"required"`
	Invoice   int    `json:"invoice" validate:"required"`
	Permainan string `json:"permainan" validate:"required"`
}
type companyinvoicelistpermainanstatus struct {
	Company string `json:"company" validate:"required"`
	Invoice int    `json:"invoice" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
type companyinvoicelistpermainanusername struct {
	Company   string `json:"company" validate:"required"`
	Invoice   int    `json:"invoice" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Permainan string `json:"permainan" validate:"required"`
}
type companysave struct {
	Sdata     string `json:"sdata" validate:"required"`
	Company   string `json:"company" validate:"required"`
	Master    string `json:"master" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Urldomain string `json:"urldomain" validate:"required"`
	Status    string `json:"status" validate:"required"`
}
type companyadminsave struct {
	Sdata          string `json:"sdata" validate:"required"`
	Company        string `json:"company" validate:"required"`
	Master         string `json:"master" validate:"required"`
	Admin_username string `json:"admin_username" validate:"required,alphanum"`
	Admin_password string `json:"admin_password" `
	Admin_name     string `json:"admin_name" validate:"required,alphanum"`
	Admin_status   string `json:"admin_status" validate:"required"`
}
type companypasaransave struct {
	Sdata      string `json:"sdata" validate:"required"`
	Company    string `json:"company" validate:"required"`
	Master     string `json:"master" validate:"required"`
	Pasaran_id string `json:"pasaran_id" validate:"required"`
}
type companyfetchpasaran432 struct {
	Company           string `json:"company" validate:"required"`
	Master            string `json:"master" validate:"required"`
	Pasaran_id        string `json:"pasaran_id" validate:"required"`
	Companypasaran_id int    `json:"companypasaran_id" validate:"required"`
}
type companypasaran struct {
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
type companypasaranline struct {
	Company               string `json:"company" validate:"required"`
	Master                string `json:"master" validate:"required"`
	Pasaran_id            string `json:"pasaran_id" validate:"required"`
	Companypasaran_id     int    `json:"companypasaran_id" validate:"required"`
	Pasaran_limitline_4d  int    `json:"pasaran_limitline_4d" validate:"required"`
	Pasaran_limitline_3d  int    `json:"pasaran_limitline_3d" validate:"required"`
	Pasaran_limitline_2d  int    `json:"pasaran_limitline_2d" validate:"required"`
	Pasaran_limitline_2dd int    `json:"pasaran_limitline_2dd" validate:"required"`
	Pasaran_limitline_2dt int    `json:"pasaran_limitline_2dt" validate:"required"`
	Pasaran_bbfs          int    `json:"pasaran_bbfs" validate:"required"`
}
type companypasaran432 struct {
	Company                     string  `json:"company" validate:"required"`
	Master                      string  `json:"master" validate:"required"`
	Pasaran_id                  string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id           int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_432d         int     `json:"pasaran_minbet_432d" validate:"required,numeric"`
	Pasaran_maxbet4d_432d       int     `json:"pasaran_maxbet4d_432d" validate:"required,numeric"`
	Pasaran_maxbet3d_432d       int     `json:"pasaran_maxbet3d_432d" validate:"required,numeric"`
	Pasaran_maxbet2d_432d       int     `json:"pasaran_maxbet2d_432d" validate:"required,numeric"`
	Pasaran_maxbet2dd_432d      int     `json:"pasaran_maxbet2dd_432d" validate:"required,numeric"`
	Pasaran_maxbet2dt_432d      int     `json:"pasaran_maxbet2dt_432d" validate:"required,numeric"`
	Pasaran_limitotal4d_432d    int     `json:"pasaran_limitotal4d_432d" validate:"required,numeric"`
	Pasaran_limitotal3d_432d    int     `json:"pasaran_limitotal3d_432d" validate:"required,numeric"`
	Pasaran_limitotal2d_432d    int     `json:"pasaran_limitotal2d_432d" validate:"required,numeric"`
	Pasaran_limitotal2dd_432d   int     `json:"pasaran_limitotal2dd_432d" validate:"required,numeric"`
	Pasaran_limitotal2dt_432d   int     `json:"pasaran_limitotal2dt_432d" validate:"required,numeric"`
	Pasaran_limitglobal4d_432d  int     `json:"pasaran_limitglobal4d_432d" validate:"required,numeric"`
	Pasaran_limitglobal3d_432d  int     `json:"pasaran_limitglobal3d_432d" validate:"required,numeric"`
	Pasaran_limitglobal2d_432d  int     `json:"pasaran_limitglobal2d_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dd_432d int     `json:"pasaran_limitglobal2dd_432d" validate:"required,numeric"`
	Pasaran_limitglobal2dt_432d int     `json:"pasaran_limitglobal2dt_432d" validate:"required,numeric"`
	Pasaran_win4d_432d          int     `json:"pasaran_win4d_432d" validate:"required,numeric"`
	Pasaran_win3d_432d          int     `json:"pasaran_win3d_432d" validate:"required,numeric"`
	Pasaran_win2d_432d          int     `json:"pasaran_win2d_432d" validate:"required,numeric"`
	Pasaran_win2dd_432d         int     `json:"pasaran_win2dd_432d" validate:"required,numeric"`
	Pasaran_win2dt_432d         int     `json:"pasaran_win2dt_432d" validate:"required,numeric"`
	Pasaran_disc4d_432d         float32 `json:"pasaran_disc4d_432d" validate:"required,numeric"`
	Pasaran_disc3d_432d         float32 `json:"pasaran_disc3d_432d" validate:"required,numeric"`
	Pasaran_disc2d_432d         float32 `json:"pasaran_disc2d_432d" validate:"required,numeric"`
	Pasaran_disc2dd_432d        float32 `json:"pasaran_disc2dd_432d" validate:"required,numeric"`
	Pasaran_disc2dt_432d        float32 `json:"pasaran_disc2dt_432d" validate:"required,numeric"`
}
type companypasarancolokbebas struct {
	Company                    string  `json:"company" validate:"required"`
	Master                     string  `json:"master" validate:"required"`
	Pasaran_id                 string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id          int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" validate:"required,numeric"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" validate:"required,numeric"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" validate:"required,numeric"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" validate:"required,numeric"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" validate:"required,numeric"`
}
type companypasarancolokmacau struct {
	Company                    string  `json:"company" validate:"required"`
	Master                     string  `json:"master" validate:"required"`
	Pasaran_id                 string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id          int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cmacau      int     `json:"pasaran_minbet_cmacau" validate:"required,numeric"`
	Pasaran_maxbet_cmacau      int     `json:"pasaran_maxbet_cmacau" validate:"required,numeric"`
	Pasaran_limitotal_cmacau   int     `json:"pasaran_limitotal_cmacau" validate:"required,numeric"`
	Pasaran_limitglobal_cmacau int     `json:"pasaran_limitglobal_cmacau" validate:"required,numeric"`
	Pasaran_win2_cmacau        float32 `json:"pasaran_win2_cmacau" validate:"required,numeric"`
	Pasaran_win3_cmacau        float32 `json:"pasaran_win3_cmacau" validate:"required,numeric"`
	Pasaran_win4_cmacau        float32 `json:"pasaran_win4_cmacau" validate:"required,numeric"`
	Pasaran_disc_cmacau        float32 `json:"pasaran_disc_cmacau" validate:"required,numeric"`
}
type companypasarancoloknaga struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_cnaga      int     `json:"pasaran_minbet_cnaga" validate:"required,numeric"`
	Pasaran_maxbet_cnaga      int     `json:"pasaran_maxbet_cnaga" validate:"required,numeric"`
	Pasaran_limittotal_cnaga  int     `json:"pasaran_limittotal_cnaga" validate:"required,numeric"`
	Pasaran_limitglobal_cnaga int     `json:"pasaran_limitglobal_cnaga" validate:"required,numeric"`
	Pasaran_win3_cnaga        float32 `json:"pasaran_win3_cnaga" validate:"required,numeric"`
	Pasaran_win4_cnaga        float32 `json:"pasaran_win4_cnaga" validate:"required,numeric"`
	Pasaran_disc_cnaga        float32 `json:"pasaran_disc_cnaga" validate:"required,numeric"`
}
type companypasarancolokjitu struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
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
type companypasaran5050umum struct {
	Company                      string  `json:"company" validate:"required"`
	Master                       string  `json:"master" validate:"required"`
	Pasaran_id                   string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id            int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050umum      int     `json:"pasaran_minbet_5050umum" validate:"required,numeric"`
	Pasaran_maxbet_5050umum      int     `json:"pasaran_maxbet_5050umum" validate:"required,numeric"`
	Pasaran_limittotal_5050umum  int     `json:"pasaran_limittotal_5050umum" validate:"required,numeric"`
	Pasaran_limitglobal_5050umum int     `json:"pasaran_limitglobal_5050umum" validate:"required,numeric"`
	Pasaran_keibesar_5050umum    float64 `json:"pasaran_keibesar_5050umum" validate:"required,numeric"`
	Pasaran_keikecil_5050umum    float64 `json:"pasaran_keikecil_5050umum" validate:"required,numeric"`
	Pasaran_keigenap_5050umum    float64 `json:"pasaran_keigenap_5050umum" validate:"required,numeric"`
	Pasaran_keiganjil_5050umum   float64 `json:"pasaran_keiganjil_5050umum" validate:"required,numeric"`
	Pasaran_keitengah_5050umum   float64 `json:"pasaran_keitengah_5050umum" validate:"required,numeric"`
	Pasaran_keitepi_5050umum     float64 `json:"pasaran_keitepi_5050umum" validate:"required,numeric"`
	Pasaran_discbesar_5050umum   float64 `json:"pasaran_discbesar_5050umum" validate:"required,numeric"`
	Pasaran_disckecil_5050umum   float64 `json:"pasaran_disckecil_5050umum" validate:"required,numeric"`
	Pasaran_discgenap_5050umum   float64 `json:"pasaran_discgenap_5050umum" validate:"required,numeric"`
	Pasaran_discganjil_5050umum  float64 `json:"pasaran_discganjil_5050umum" validate:"required,numeric"`
	Pasaran_disctengah_5050umum  float64 `json:"pasaran_disctengah_5050umum" validate:"required,numeric"`
	Pasaran_disctepi_5050umum    float64 `json:"pasaran_disctepi_5050umum" validate:"required,numeric"`
}
type companypasaran5050special struct {
	Company                              string  `json:"company" validate:"required"`
	Master                               string  `json:"master" validate:"required"`
	Pasaran_id                           string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id                    int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050special           int     `json:"pasaran_minbet_5050special" validate:"required,numeric"`
	Pasaran_maxbet_5050special           int     `json:"pasaran_maxbet_5050special" validate:"required,numeric"`
	Pasaran_limitglobal_5050special      int     `json:"pasaran_limitglobal_5050special" validate:"required,numeric"`
	Pasaran_limittotal_5050special       int     `json:"pasaran_limittotal_5050special" validate:"required,numeric"`
	Pasaran_keiasganjil_5050special      float64 `json:"pasaran_keiasganjil_5050special" validate:"numeric"`
	Pasaran_keiasgenap_5050special       float64 `json:"pasaran_keiasgenap_5050special" validate:"numeric"`
	Pasaran_keiasbesar_5050special       float64 `json:"pasaran_keiasbesar_5050special" validate:"numeric"`
	Pasaran_keiaskecil_5050special       float64 `json:"pasaran_keiaskecil_5050special" validate:"numeric"`
	Pasaran_keikopganjil_5050special     float64 `json:"pasaran_keikopganjil_5050special" validate:"numeric"`
	Pasaran_keikopgenap_5050special      float64 `json:"pasaran_keikopgenap_5050special" validate:"numeric"`
	Pasaran_keikopbesar_5050special      float64 `json:"pasaran_keikopbesar_5050special" validate:"numeric"`
	Pasaran_keikopkecil_5050special      float64 `json:"pasaran_keikopkecil_5050special" validate:"numeric"`
	Pasaran_keikepalaganjil_5050special  float64 `json:"pasaran_keikepalaganjil_5050special" validate:"numeric"`
	Pasaran_keikepalagenap_5050special   float64 `json:"pasaran_keikepalagenap_5050special" validate:"numeric"`
	Pasaran_keikepalabesar_5050special   float64 `json:"pasaran_keikepalabesar_5050special" validate:"numeric"`
	Pasaran_keikepalakecil_5050special   float64 `json:"pasaran_keikepalakecil_5050special" validate:"numeric"`
	Pasaran_keiekorganjil_5050special    float64 `json:"pasaran_keiekorganjil_5050special" validate:"numeric"`
	Pasaran_keiekorgenap_5050special     float64 `json:"pasaran_keiekorgenap_5050special" validate:"numeric"`
	Pasaran_keiekorbesar_5050special     float64 `json:"pasaran_keiekorbesar_5050special" validate:"numeric"`
	Pasaran_keiekorkecil_5050special     float64 `json:"pasaran_keiekorkecil_5050special" validate:"numeric"`
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
type companypasaran5050kombinasi struct {
	Company                                   string  `json:"company" validate:"required"`
	Master                                    string  `json:"master" validate:"required"`
	Pasaran_id                                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id                         int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_5050kombinasi              int     `json:"pasaran_minbet_5050kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_5050kombinasi              int     `json:"pasaran_maxbet_5050kombinasi" validate:"required,numeric"`
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
	Pasaran_belakangdiscmono_5050kombinasi    float64 `json:"pasaran_belakangdiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdiscstereo_5050kombinasi  float64 `json:"pasaran_belakangdiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckembang_5050kombinasi float64 `json:"pasaran_belakangdisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckempis_5050kombinasi  float64 `json:"pasaran_belakangdisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_belakangdisckembar_5050kombinasi  float64 `json:"pasaran_belakangdisckembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdiscmono_5050kombinasi      float64 `json:"pasaran_tengahdiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdiscstereo_5050kombinasi    float64 `json:"pasaran_tengahdiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckembang_5050kombinasi   float64 `json:"pasaran_tengahdisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckempis_5050kombinasi    float64 `json:"pasaran_tengahdisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_tengahdisckembar_5050kombinasi    float64 `json:"pasaran_tengahdisckembar_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandiscmono_5050kombinasi       float64 `json:"pasaran_depandiscmono_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandiscstereo_5050kombinasi     float64 `json:"pasaran_depandiscstereo_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckembang_5050kombinasi    float64 `json:"pasaran_depandisckembang_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckempis_5050kombinasi     float64 `json:"pasaran_depandisckempis_5050kombinasi" validate:"required,numeric"`
	Pasaran_depandisckembar_5050kombinasi     float64 `json:"pasaran_depandisckembar_5050kombinasi" validate:"required,numeric"`
}
type companypasarankombinasi struct {
	Company                       string  `json:"company" validate:"required"`
	Master                        string  `json:"master" validate:"required"`
	Pasaran_id                    string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id             int     `json:"companypasaran_id" validate:"required"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" validate:"required,numeric"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" validate:"required,numeric"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" validate:"required,numeric"`
}
type companypasarandasar struct {
	Company                   string  `json:"company" validate:"required"`
	Master                    string  `json:"master" validate:"required"`
	Pasaran_id                string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id         int     `json:"companypasaran_id" validate:"required"`
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
type companypasaranshio struct {
	Company                  string  `json:"company" validate:"required"`
	Master                   string  `json:"master" validate:"required"`
	Pasaran_id               string  `json:"pasaran_id" validate:"required"`
	Companypasaran_id        int     `json:"companypasaran_id" validate:"required"`
	Pasaran_shioyear_shio    string  `json:"pasaran_shioyear_shio" validate:"required"`
	Pasaran_minbet_shio      int     `json:"pasaran_minbet_shio" validate:"required,numeric"`
	Pasaran_maxbet_shio      int     `json:"pasaran_maxbet_shio" validate:"required,numeric"`
	Pasaran_limitglobal_shio int     `json:"pasaran_limitglobal_shio" validate:"required,numeric"`
	Pasaran_limittotal_shio  int     `json:"pasaran_limittotal_shio" validate:"required,numeric"`
	Pasaran_disc_shio        float32 `json:"pasaran_disc_shio" validate:"numeric"`
	Pasaran_win_shio         float32 `json:"pasaran_win_shio" validate:"numeric"`
}
type rediscompanyhome struct {
	No          int    `json:"company_no"`
	Idcompany   string `json:"company_idcompany"`
	Startjoin   string `json:"company_startjoin"`
	Endjoin     string `json:"company_endjoin"`
	Curr        string `json:"company_curr"`
	Name        string `json:"company_name"`
	Periode     string `json:"company_periode"`
	Winlose     int    `json:"company_winlose"`
	Winlosetemp int    `json:"company_winlosetemp"`
	Status      string `json:"company_status"`
	Statuscss   string `json:"company_statuscss"`
}

func CompanyHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyhome)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	field_redis := "LISTCOMPANY_MASTER"
	render_page := time.Now()
	var obj rediscompanyhome
	var arraobj []rediscompanyhome
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_no, _ := jsonparser.GetInt(value, "company_no")
		company_idcompany, _ := jsonparser.GetString(value, "company_idcompany")
		company_startjoin, _ := jsonparser.GetString(value, "company_startjoin")
		company_endjoin, _ := jsonparser.GetString(value, "company_endjoin")
		company_curr, _ := jsonparser.GetString(value, "company_curr")
		company_name, _ := jsonparser.GetString(value, "company_name")
		company_periode, _ := jsonparser.GetString(value, "company_periode")
		company_winlose, _ := jsonparser.GetInt(value, "company_winlose")
		company_winlosetemp, _ := jsonparser.GetInt(value, "company_winlosetemp")
		company_status, _ := jsonparser.GetString(value, "company_status")
		company_statuscss, _ := jsonparser.GetString(value, "company_statuscss")

		obj.No = int(company_no)
		obj.Idcompany = company_idcompany
		obj.Startjoin = company_startjoin
		obj.Endjoin = company_endjoin
		obj.Curr = company_curr
		obj.Name = company_name
		obj.Periode = company_periode
		obj.Winlose = int(company_winlose)
		obj.Winlosetemp = int(company_winlosetemp)
		obj.Status = company_status
		obj.Statuscss = company_statuscss
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company()
		helpers.SetRedis(field_redis, result, 30*time.Minute)
		log.Println("COMPANY MYSQL")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		return c.JSON(result)
	} else {
		log.Println("COMPANY CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetail)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_companyDetail(client.Company)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyDetailListAdmin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetail)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_listadmin(client.Company)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyDetailListPasaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetail)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_listpasaran(client.Company)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyDetailListPasaranOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetailonline)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_listpasaranonline(client.Company, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyDetailListPasaranConf(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaranconf)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_listpasaranConf(client.Company, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyListKeluaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companylistkeluaran)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_listkeluaran(client.Company, client.Periode, client.Pasaran)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoiceMember(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoice)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_member(client.Company, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoiceMemberTemp(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoice)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_membertemp(client.Company, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoiceMemberSync(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoice)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_membersync(client.Company, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoiceGroupPermainan(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoice)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_grouppermainan(client.Company, client.Username, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoicelistpermainan(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoicelistpermainan)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_listpermainan(client.Company, client.Permainan, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoicelistpermainanbystatus(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoicelistpermainanstatus)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_listpermainanbystatus(client.Company, client.Status, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanyInvoicelistpermainanbyusername(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyinvoicelistpermainanusername)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Fetch_company_invoice_listpermainanbyusername(client.Company, client.Username, client.Permainan, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanySave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companysave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_company(
		client.Sdata,
		client.Master,
		client.Company,
		client.Name, client.Urldomain, client.Status)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanySaveNewAdmin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyadminsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyNewAdmin(
		client.Sdata, client.Master, client.Company,
		client.Admin_username, client.Admin_password, client.Admin_name, client.Admin_status)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func CompanySaveNewPasaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaransave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyNewPasaran(client.Master, client.Company, client.Pasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_agent_periode := helpers.DeleteRedis("LISTPERIODE_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PERIODE: %d", val_agent_periode)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanySaveNewPasaranHariOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetailonlinesave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyInsertPasaranharionline(client.Master, client.Company, client.Companypasaran_id, client.Pasaran_hari)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyDeletePasaranHariOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companydetailonlinedelete)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Delete_companyPasaranharionline(client.Master, client.Company, client.Companypasaran_id, client.Companypasaran_idoff)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaranlimitline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaranlimitline(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_4-3-2")
	log.Printf("REDIS DELETE FRONTEND CONFIG 4-3-2: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaran432(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaran432(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_4-3-2")
	log.Printf("REDIS DELETE FRONTEND CONFIG 4-3-2: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasarancolokbebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasarancolokbebas(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasarancolokmacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasarancolokmacau(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasarancoloknaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasarancoloknaga(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasarancolokjitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasarancolokjitu(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaran5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaran5050umum(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaran5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaran5050special(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaran5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaran5050kombinasi(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaranmacaukombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaranmacau(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_macaukombinasi")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasarandasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasarandasar(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_dasar")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyFetchPasaranshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companyfetchpasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	result, err := models.Fetch_companyPasaranshio(client.Master, client.Company, client.Pasaran_id, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_shio")
	log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdate(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaran)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaran(
		client.Master, client.Company, client.Pasaran_diundi,
		client.Pasaran_url, client.Pasaran_jamtutup, client.Pasaran_jamjadwal,
		client.Pasaran_jamopen, client.Pasaran_statusactive, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdateLimitline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaranline)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaranLine(
		client.Master, client.Company, client.Companypasaran_id, client.Pasaran_limitline_4d,
		client.Pasaran_limitline_3d, client.Pasaran_limitline_2d, client.Pasaran_limitline_2dd,
		client.Pasaran_limitline_2dt, client.Pasaran_bbfs)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_4-3-2")
	log.Printf("REDIS DELETE FRONTEND CONFIG 4-3-2: %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdate432(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaran432)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaran432(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_432d, client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d,
		client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
		client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
		client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
		client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
		client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_4-3-2")
	log.Printf("REDIS DELETE FRONTEND CONF 4-3-2 : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokbebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarancolokbebas)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarancolokbebas(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_cbebas, client.Pasaran_maxbet_cbebas,
		client.Pasaran_win_cbebas, client.Pasaran_disc_cbebas,
		client.Pasaran_limitglobal_cbebas, client.Pasaran_limitotal_cbebas)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONF COLOK : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokmacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarancolokmacau)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarancolokmacau(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_cmacau, client.Pasaran_maxbet_cmacau,
		client.Pasaran_win2_cmacau, client.Pasaran_win3_cmacau, client.Pasaran_win4_cmacau, client.Pasaran_disc_cmacau,
		client.Pasaran_limitglobal_cmacau, client.Pasaran_limitotal_cmacau)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONF COLOK : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatecoloknaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarancoloknaga)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarancoloknaga(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_cnaga, client.Pasaran_maxbet_cnaga,
		client.Pasaran_win3_cnaga, client.Pasaran_win4_cnaga, client.Pasaran_disc_cnaga,
		client.Pasaran_limitglobal_cnaga, client.Pasaran_limittotal_cnaga)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONF COLOK : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokjitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarancolokjitu)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarancolokjitu(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_cjitu, client.Pasaran_maxbet_cjitu,
		client.Pasaran_winas_cjitu, client.Pasaran_winkop_cjitu, client.Pasaran_winkepala_cjitu, client.Pasaran_winekor_cjitu, client.Pasaran_desc_cjitu,
		client.Pasaran_limitglobal_cjitu, client.Pasaran_limittotal_cjitu)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_colok")
	log.Printf("REDIS DELETE FRONTEND CONF COLOK : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdate5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaran5050umum)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaran5050umum(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_5050umum, client.Pasaran_maxbet_5050umum,
		client.Pasaran_keibesar_5050umum, client.Pasaran_keikecil_5050umum, client.Pasaran_keigenap_5050umum, client.Pasaran_keiganjil_5050umum, client.Pasaran_keitengah_5050umum, client.Pasaran_keitepi_5050umum,
		client.Pasaran_discbesar_5050umum, client.Pasaran_disckecil_5050umum, client.Pasaran_discgenap_5050umum, client.Pasaran_discganjil_5050umum, client.Pasaran_disctengah_5050umum, client.Pasaran_disctepi_5050umum,
		client.Pasaran_limitglobal_5050umum, client.Pasaran_limittotal_5050umum)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONF 5050 : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdate5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaran5050special)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaran5050special(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_5050special, client.Pasaran_maxbet_5050special,
		client.Pasaran_keiasganjil_5050special, client.Pasaran_keiasgenap_5050special, client.Pasaran_keiasbesar_5050special, client.Pasaran_keiaskecil_5050special,
		client.Pasaran_keikopganjil_5050special, client.Pasaran_keikopgenap_5050special, client.Pasaran_keikopbesar_5050special, client.Pasaran_keikopkecil_5050special,
		client.Pasaran_keikepalaganjil_5050special, client.Pasaran_keikepalagenap_5050special, client.Pasaran_keikepalabesar_5050special, client.Pasaran_keikepalakecil_5050special,
		client.Pasaran_keiekorganjil_5050special, client.Pasaran_keiekorgenap_5050special, client.Pasaran_keiekorbesar_5050special, client.Pasaran_keiekorkecil_5050special,
		client.Pasaran_discasganjil_5050special, client.Pasaran_discasgenap_5050special, client.Pasaran_discasbesar_5050special, client.Pasaran_discaskecil_5050special,
		client.Pasaran_disckopganjil_5050special, client.Pasaran_disckopgenap_5050special, client.Pasaran_disckopbesar_5050special, client.Pasaran_disckopkecil_5050special,
		client.Pasaran_disckepalaganjil_5050special, client.Pasaran_disckepalagenap_5050special, client.Pasaran_disckepalabesar_5050special, client.Pasaran_disckepalakecil_5050special,
		client.Pasaran_discekorganjil_5050special, client.Pasaran_discekorgenap_5050special, client.Pasaran_discekorbesar_5050special, client.Pasaran_discekorkecil_5050special,
		client.Pasaran_limitglobal_5050special, client.Pasaran_limittotal_5050special)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONF 5050 : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdate5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaran5050kombinasi)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaran5050kombinasi(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_5050kombinasi, client.Pasaran_maxbet_5050kombinasi,
		client.Pasaran_belakangkeimono_5050kombinasi, client.Pasaran_belakangkeistereo_5050kombinasi, client.Pasaran_belakangkeikembang_5050kombinasi, client.Pasaran_belakangkeikempis_5050kombinasi, client.Pasaran_belakangkeikembar_5050kombinasi,
		client.Pasaran_tengahkeimono_5050kombinasi, client.Pasaran_tengahkeistereo_5050kombinasi, client.Pasaran_tengahkeikembang_5050kombinasi, client.Pasaran_tengahkeikempis_5050kombinasi, client.Pasaran_tengahkeikembar_5050kombinasi,
		client.Pasaran_depankeimono_5050kombinasi, client.Pasaran_depankeistereo_5050kombinasi, client.Pasaran_depankeikembang_5050kombinasi, client.Pasaran_depankeikempis_5050kombinasi, client.Pasaran_depankeikembar_5050kombinasi,
		client.Pasaran_belakangdiscmono_5050kombinasi, client.Pasaran_belakangdiscstereo_5050kombinasi, client.Pasaran_belakangdisckembang_5050kombinasi, client.Pasaran_belakangdisckempis_5050kombinasi, client.Pasaran_belakangdisckembar_5050kombinasi,
		client.Pasaran_tengahdiscmono_5050kombinasi, client.Pasaran_tengahdiscstereo_5050kombinasi, client.Pasaran_tengahdisckembang_5050kombinasi, client.Pasaran_tengahdisckempis_5050kombinasi, client.Pasaran_tengahdisckembar_5050kombinasi,
		client.Pasaran_depandiscmono_5050kombinasi, client.Pasaran_depandiscstereo_5050kombinasi, client.Pasaran_depandisckembang_5050kombinasi, client.Pasaran_depandisckempis_5050kombinasi, client.Pasaran_depandisckembar_5050kombinasi,
		client.Pasaran_limitglobal_5050kombinasi, client.Pasaran_limittotal_5050kombinasi)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_5050")
	log.Printf("REDIS DELETE FRONTEND CONF 5050 : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatekombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarankombinasi)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarankombinasi(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_kombinasi, client.Pasaran_maxbet_kombinasi,
		client.Pasaran_win_kombinasi, client.Pasaran_disc_kombinasi,
		client.Pasaran_limitglobal_kombinasi, client.Pasaran_limittotal_kombinasi)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_macaukombinasi")
	log.Printf("REDIS DELETE FRONTEND CONF MACAU KOMBINASI : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdatedasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasarandasar)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasarandasar(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_minbet_dasar, client.Pasaran_maxbet_dasar,
		client.Pasaran_keibesar_dasar, client.Pasaran_keikecil_dasar, client.Pasaran_keigenap_dasar, client.Pasaran_keiganjil_dasar,
		client.Pasaran_discbesar_dasar, client.Pasaran_disckecil_dasar, client.Pasaran_discgenap_dasar, client.Pasaran_discganjil_dasar,
		client.Pasaran_limitglobal_dasar, client.Pasaran_limittotal_dasar)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_dasar")
	log.Printf("REDIS DELETE FRONTEND CONF DASAR : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
func CompanyPasaranUpdateshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(companypasaranshio)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, err := models.Save_companyUpdatePasaranshio(
		client.Master, client.Company, client.Companypasaran_id,
		client.Pasaran_shioyear_shio,
		client.Pasaran_minbet_shio, client.Pasaran_maxbet_shio,
		client.Pasaran_win_shio, client.Pasaran_disc_shio, client.Pasaran_limitglobal_shio, client.Pasaran_limittotal_shio)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val := helpers.DeleteRedis("config_" + client.Company + "_" + client.Pasaran_id + "_shio")
	log.Printf("REDIS DELETE FRONTEND CONF SHIO : %d", val)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company)
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)
	val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + client.Company)
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)
	return c.JSON(result)
}
