package controller

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
)

type pasaranhome struct {
	Master string `json:"master" validate:"required"`
}
type pasarandetail struct {
	Pasarancode string `json:"pasarancode" validate:"required"`
	Master      string `json:"master" validate:"required"`
}
type pasaransave struct {
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
type pasaransavelimitline struct {
	Master               string `json:"master" validate:"required"`
	Idrecord             string `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_bbfs         int    `json:"pasaran_bbfs" validate:"required,numeric"`
	Pasaran_limitline4d  int    `json:"pasaran_limitline4d" validate:"required,numeric"`
	Pasaran_limitline3d  int    `json:"pasaran_limitline3d" validate:"required,numeric"`
	Pasaran_limitline2d  int    `json:"pasaran_limitline2d" validate:"required,numeric"`
	Pasaran_limitline2dd int    `json:"pasaran_limitline2dd" validate:"required,numeric"`
	Pasaran_limitline2dt int    `json:"pasaran_limitline2dt" validate:"required,numeric"`
}
type pasaransaveconf432d struct {
	Master                      string  `json:"master" validate:"required"`
	Idrecord                    string  `json:"idrecord" validate:"required,min=2,max=10"`
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
type pasaransaveconfcolokbebas struct {
	Master                     string  `json:"master" validate:"required"`
	Idrecord                   string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_cbebas      int     `json:"pasaran_minbet_cbebas" validate:"required,numeric"`
	Pasaran_maxbet_cbebas      int     `json:"pasaran_maxbet_cbebas" validate:"required,numeric"`
	Pasaran_limitotal_cbebas   int     `json:"pasaran_limitotal_cbebas" validate:"required,numeric"`
	Pasaran_limitglobal_cbebas int     `json:"pasaran_limitglobal_cbebas" validate:"required,numeric"`
	Pasaran_win_cbebas         float32 `json:"pasaran_win_cbebas" validate:"required,numeric"`
	Pasaran_disc_cbebas        float32 `json:"pasaran_disc_cbebas" validate:"required,numeric"`
}
type pasaransaveconfcolokmacau struct {
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
type pasaransaveconfcoloknaga struct {
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
type pasaransaveconfcolokjitu struct {
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
type pasaransaveconf5050umum struct {
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
type pasaransaveconf5050special struct {
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
type pasaransaveconf5050kombinasi struct {
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
type pasaransaveconfmacaukombinasi struct {
	Master                        string  `json:"master" validate:"required"`
	Idrecord                      string  `json:"idrecord" validate:"required,min=2,max=10"`
	Pasaran_minbet_kombinasi      int     `json:"pasaran_minbet_kombinasi" validate:"required,numeric"`
	Pasaran_maxbet_kombinasi      int     `json:"pasaran_maxbet_kombinasi" validate:"required,numeric"`
	Pasaran_limitglobal_kombinasi int     `json:"pasaran_limitglobal_kombinasi" validate:"required,numeric"`
	Pasaran_limittotal_kombinasi  int     `json:"pasaran_limittotal_kombinasi" validate:"required,numeric"`
	Pasaran_win_kombinasi         float32 `json:"pasaran_win_kombinasi" validate:"required,numeric"`
	Pasaran_disc_kombinasi        float32 `json:"pasaran_disc_kombinasi" validate:"required,numeric"`
}
type pasaransaveconfdasar struct {
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
type pasaransaveconfshio struct {
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
type redispasaranhome struct {
	No             int    `json:"pasaran_no"`
	Idpasarantogel string `json:"pasaran_idpasarantogel"`
	Nmpasarantogel string `json:"pasaran_nmpasarantogel"`
	Tipepasaran    string `json:"pasaran_tipepasaran"`
	Urlpasaran     string `json:"pasaran_urlpasaran"`
	Pasarandiundi  string `json:"pasaran_pasarandiundi"`
	Jamtutup       string `json:"pasaran_jamtutup"`
	Jamjadwal      string `json:"pasaran_jamjadwal"`
	Jamopen        string `json:"pasaran_jamopen"`
}

func PasaranHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaranhome)
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
	field_redis := "LISTPASARAN_MASTER"
	render_page := time.Now()
	var obj redispasaranhome
	var arraobj []redispasaranhome
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_no, _ := jsonparser.GetInt(value, "pasaran_no")
		pasaran_idpasarantogel, _ := jsonparser.GetString(value, "pasaran_idpasarantogel")
		pasaran_nmpasarantogel, _ := jsonparser.GetString(value, "pasaran_nmpasarantogel")
		pasaran_tipepasaran, _ := jsonparser.GetString(value, "pasaran_tipepasaran")
		pasaran_urlpasaran, _ := jsonparser.GetString(value, "pasaran_urlpasaran")
		pasaran_pasarandiundi, _ := jsonparser.GetString(value, "pasaran_pasarandiundi")
		pasaran_jamtutup, _ := jsonparser.GetString(value, "pasaran_jamtutup")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_jamopen, _ := jsonparser.GetString(value, "pasaran_jamopen")

		obj.No = int(pasaran_no)
		obj.Idpasarantogel = pasaran_idpasarantogel
		obj.Nmpasarantogel = pasaran_nmpasarantogel
		obj.Tipepasaran = pasaran_tipepasaran
		obj.Urlpasaran = pasaran_urlpasaran
		obj.Pasarandiundi = pasaran_pasarandiundi
		obj.Jamtutup = pasaran_jamtutup
		obj.Jamjadwal = pasaran_jamjadwal
		obj.Jamopen = pasaran_jamopen
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaran()
		helpers.SetRedis(field_redis, result, 0)
		log.Println("PASARAN MYSQL")
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
		log.Println("PASARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PasaranDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandetail)
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

	result, err := models.Fetch_pasaranDetail(client.Pasarancode)
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
func PasaranDetailConf(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasarandetail)
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

	result, err := models.Fetch_pasaranDetailConf(client.Pasarancode)
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
func PasaranSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransave)
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
	field_redis := "LISTPASARAN_MASTER"
	result, err := models.Save_pasaran(client.Sdata, client.Master, client.Idrecord, client.Name, client.Tipe, client.Url, client.Diundi, client.Jamtutup, client.Jamjadwal, client.Jamopen)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_master := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete MASTER - PASARAN status: %d", val_master)
	return c.JSON(result)
}
func PasaranSaveLimitLine(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransavelimitline)
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
	result, err := models.Save_pasaranlimitline(
		client.Master, client.Idrecord,
		client.Pasaran_limitline4d, client.Pasaran_limitline3d, client.Pasaran_limitline2d,
		client.Pasaran_limitline2dd, client.Pasaran_limitline2dt, client.Pasaran_bbfs)
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
func PasaranSaveConf432d(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconf432d)
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

	result, err := models.Save_pasaranConf432(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfColokBebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfcolokbebas)
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

	result, err := models.Save_pasaranConfColokBebas(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfColokMacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfcolokmacau)
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

	result, err := models.Save_pasaranConfColokMacau(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfColokNaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfcoloknaga)
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

	result, err := models.Save_pasaranConfColokNaga(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfColokJitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfcolokjitu)
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

	result, err := models.Save_pasaranConfColokJitu(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConf5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconf5050umum)
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

	result, err := models.Save_pasaranConf5050umum(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConf5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconf5050special)
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

	result, err := models.Save_pasaranConf5050special(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConf5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconf5050kombinasi)
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

	result, err := models.Save_pasaranConf5050kombinasi(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfmacaukombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfmacaukombinasi)
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

	result, err := models.Save_pasaranConfmacaukombinasi(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfdasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfdasar)
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

	result, err := models.Save_pasaranConfdasar(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
func PasaranSaveConfshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(pasaransaveconfshio)
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

	result, err := models.Save_pasaranConfshio(
		client.Master, client.Idrecord,
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
	return c.JSON(result)
}
