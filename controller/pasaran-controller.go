package controller

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
)

const Fieldpasaran_home_redis = "LISTPASARAN_MASTER"
const Fieldpasarandetail_home_redis = "LISTPASARANDETAIL_MASTER"
const FieldpasarandetailCONF_home_redis = "LISTPASARANDETAILCONF_MASTER"

func PasaranHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaran)
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
	render_page := time.Now()
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	resultredis, flag := helpers.GetRedis(Fieldpasaran_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_idpasarantogel, _ := jsonparser.GetString(value, "pasaran_idpasarantogel")
		pasaran_nmpasarantogel, _ := jsonparser.GetString(value, "pasaran_nmpasarantogel")
		pasaran_tipepasaran, _ := jsonparser.GetString(value, "pasaran_tipepasaran")
		pasaran_urlpasaran, _ := jsonparser.GetString(value, "pasaran_urlpasaran")
		pasaran_pasarandiundi, _ := jsonparser.GetString(value, "pasaran_pasarandiundi")
		pasaran_jamtutup, _ := jsonparser.GetString(value, "pasaran_jamtutup")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_jamopen, _ := jsonparser.GetString(value, "pasaran_jamopen")

		obj.Pasaran_idpasarantogel = pasaran_idpasarantogel
		obj.Pasaran_nmpasarantogel = pasaran_nmpasarantogel
		obj.Pasaran_tipepasaran = pasaran_tipepasaran
		obj.Pasaran_urlpasaran = pasaran_urlpasaran
		obj.Pasaran_pasarandiundi = pasaran_pasarandiundi
		obj.Pasaran_jamtutup = pasaran_jamtutup
		obj.Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Pasaran_jamopen = pasaran_jamopen
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaran()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpasaran_home_redis, result, 60*time.Minute)
		log.Println("PASARAN MYSQL")
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
	client := new(entities.Controller_pasarandetail)
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
	render_page := time.Now()
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	resultredis, flag := helpers.GetRedis(Fieldpasarandetail_home_redis + "_" + client.Pasarancode)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_nmpasarantogel, _ := jsonparser.GetString(value, "pasaran_nmpasarantogel")
		pasaran_tipepasaran, _ := jsonparser.GetString(value, "pasaran_tipepasaran")
		pasaran_urlpasaran, _ := jsonparser.GetString(value, "pasaran_urlpasaran")
		pasaran_pasarandiundi, _ := jsonparser.GetString(value, "pasaran_pasarandiundi")
		pasaran_jamtutup, _ := jsonparser.GetString(value, "pasaran_jamtutup")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_jamopen, _ := jsonparser.GetString(value, "pasaran_jamopen")
		pasaran_create, _ := jsonparser.GetString(value, "pasaran_create")
		pasaran_update, _ := jsonparser.GetString(value, "pasaran_update")

		obj.Pasaran_nmpasarantogel = pasaran_nmpasarantogel
		obj.Pasaran_tipepasaran = pasaran_tipepasaran
		obj.Pasaran_urlpasaran = pasaran_urlpasaran
		obj.Pasaran_pasarandiundi = pasaran_pasarandiundi
		obj.Pasaran_jamtutup = pasaran_jamtutup
		obj.Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Pasaran_jamopen = pasaran_jamopen
		obj.Pasaran_create = pasaran_create
		obj.Pasaran_update = pasaran_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaranDetail(client.Pasarancode)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpasarandetail_home_redis+"_"+client.Pasarancode, result, 60*time.Minute)
		log.Println("PASARAN MYSQL")
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
func PasaranDetailConf(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasarandetail)
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
	render_page := time.Now()
	var obj entities.Model_pasaranDetailConf
	var arraobj []entities.Model_pasaranDetailConf
	resultredis, flag := helpers.GetRedis(FieldpasarandetailCONF_home_redis + "_" + client.Pasarancode)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		limitline_4d, _ := jsonparser.GetInt(value, "limitline_4d")
		limitline_3d, _ := jsonparser.GetInt(value, "limitline_3d")
		limitline_3dd, _ := jsonparser.GetInt(value, "limitline_3dd")
		limitline_2d, _ := jsonparser.GetInt(value, "limitline_2d")
		limitline_2dd, _ := jsonparser.GetInt(value, "limitline_2dd")
		limitline_2dt, _ := jsonparser.GetInt(value, "limitline_2dt")
		bbfs, _ := jsonparser.GetInt(value, "bbfs")
		minbet_432d, _ := jsonparser.GetFloat(value, "minbet_432d")
		maxbet4d_432d, _ := jsonparser.GetFloat(value, "maxbet4d_432d")
		maxbet3d_432d, _ := jsonparser.GetFloat(value, "maxbet3d_432d")
		maxbet3dd_432d, _ := jsonparser.GetFloat(value, "maxbet3dd_432d")
		maxbet2d_432d, _ := jsonparser.GetFloat(value, "maxbet2d_432d")
		maxbet2dd_432d, _ := jsonparser.GetFloat(value, "maxbet2dd_432d")
		maxbet2dt_432d, _ := jsonparser.GetFloat(value, "maxbet2dt_432d")
		maxbet4dfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet4dfullbb_432d")
		maxbet3dfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet3dfullbb_432d")
		maxbet3ddfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet3ddfullbb_432d")
		maxbet2dfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2dfullbb_432d")
		maxbet2ddfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2ddfullbb_432d")
		maxbet2dtfullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2dtfullbb_432d")
		Maxbuy4d_432d, _ := jsonparser.GetFloat(value, "Maxbuy4d_432d")
		Maxbuy3d_432d, _ := jsonparser.GetFloat(value, "Maxbuy3d_432d")
		Maxbuy3dd_432d, _ := jsonparser.GetFloat(value, "Maxbuy3dd_432d")
		Maxbuy2d_432d, _ := jsonparser.GetFloat(value, "Maxbuy2d_432d")
		Maxbuy2dd_432d, _ := jsonparser.GetFloat(value, "Maxbuy2dd_432d")
		Maxbuy2dt_432d, _ := jsonparser.GetFloat(value, "Maxbuy2dt_432d")
		limitotal4d_432d, _ := jsonparser.GetFloat(value, "limitotal4d_432d")
		limitotal3d_432d, _ := jsonparser.GetFloat(value, "limitotal3d_432d")
		limitotal3dd_432d, _ := jsonparser.GetFloat(value, "limitotal3dd_432d")
		limitotal2d_432d, _ := jsonparser.GetFloat(value, "limitotal2d_432d")
		limitotal2dd_432d, _ := jsonparser.GetFloat(value, "limitotal2dd_432d")
		limitotal2dt_432d, _ := jsonparser.GetFloat(value, "limitotal2dt_432d")
		limitglobal4d_432d, _ := jsonparser.GetFloat(value, "limitglobal4d_432d")
		limitglobal3d_432d, _ := jsonparser.GetFloat(value, "limitglobal3d_432d")
		limitglobal3dd_432d, _ := jsonparser.GetFloat(value, "limitglobal3dd_432d")
		limitglobal2d_432d, _ := jsonparser.GetFloat(value, "limitglobal2d_432d")
		limitglobal2dd_432d, _ := jsonparser.GetFloat(value, "limitglobal2dd_432d")
		limitglobal2dt_432d, _ := jsonparser.GetFloat(value, "limitglobal2dt_432d")
		limitotal4d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal4d_fullbb_432d")
		limitotal3d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal3d_fullbb_432d")
		limitotal3dd_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal3dd_fullbb_432d")
		limitotal2d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal2d_fullbb_432d")
		limitotal2dd_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal2dd_fullbb_432d")
		limitotal2dt_fullbb_432d, _ := jsonparser.GetFloat(value, "limitotal2dt_fullbb_432d")
		limitglobal4d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal4d_fullbb_432d")
		limitglobal3d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal3d_fullbb_432d")
		limitglobal3dd_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal3dd_fullbb_432d")
		limitglobal2d_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal2d_fullbb_432d")
		limitglobal2dd_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal2dd_fullbb_432d")
		limitglobal2dt_fullbb_432d, _ := jsonparser.GetFloat(value, "limitglobal2dt_fullbb_432d")
		disc4d_432d, _ := jsonparser.GetFloat(value, "disc4d_432d")
		disc3d_432d, _ := jsonparser.GetFloat(value, "disc3d_432d")
		disc3dd_432d, _ := jsonparser.GetFloat(value, "disc3dd_432d")
		disc2d_432d, _ := jsonparser.GetFloat(value, "disc2d_432d")
		disc2dd_432d, _ := jsonparser.GetFloat(value, "disc2dd_432d")
		disc2dt_432d, _ := jsonparser.GetFloat(value, "disc2dt_432d")
		win4d_432d, _ := jsonparser.GetFloat(value, "win4d_432d")
		win3d_432d, _ := jsonparser.GetFloat(value, "win3d_432d")
		win3dd_432d, _ := jsonparser.GetFloat(value, "win3dd_432d")
		win2d_432d, _ := jsonparser.GetFloat(value, "win2d_432d")
		win2dd_432d, _ := jsonparser.GetFloat(value, "win2dd_432d")
		win2dt_432d, _ := jsonparser.GetFloat(value, "win2dt_432d")
		win4dnodisc_432d, _ := jsonparser.GetFloat(value, "win4dnodisc_432d")
		win3dnodisc_432d, _ := jsonparser.GetFloat(value, "win3dnodisc_432d")
		win3ddnodisc_432d, _ := jsonparser.GetFloat(value, "win3ddnodisc_432d")
		win2dnodisc_432d, _ := jsonparser.GetFloat(value, "win2dnodisc_432d")
		win2ddnodisc_432d, _ := jsonparser.GetFloat(value, "win2ddnodisc_432d")
		win2dtnodisc_432d, _ := jsonparser.GetFloat(value, "win2dtnodisc_432d")
		win4dbb_kena_432d, _ := jsonparser.GetFloat(value, "win4dbb_kena_432d")
		win3dbb_kena_432d, _ := jsonparser.GetFloat(value, "win3dbb_kena_432d")
		win3ddbb_kena_432d, _ := jsonparser.GetFloat(value, "win3ddbb_kena_432d")
		win2dbb_kena_432d, _ := jsonparser.GetFloat(value, "win2dbb_kena_432d")
		win2ddbb_kena_432d, _ := jsonparser.GetFloat(value, "win2ddbb_kena_432d")
		win2dtbb_kena_432d, _ := jsonparser.GetFloat(value, "win2dtbb_kena_432d")
		win4dbb_432d, _ := jsonparser.GetFloat(value, "win4dbb_432d")
		win3dbb_432d, _ := jsonparser.GetFloat(value, "win3dbb_432d")
		win3ddbb_432d, _ := jsonparser.GetFloat(value, "win3ddbb_432d")
		win2dbb_432d, _ := jsonparser.GetFloat(value, "win2dbb_432d")
		win2ddbb_432d, _ := jsonparser.GetFloat(value, "win2ddbb_432d")
		win2dtbb_432d, _ := jsonparser.GetFloat(value, "win2dtbb_432d")
		minbet_cbebas, _ := jsonparser.GetFloat(value, "minbet_cbebas")
		maxbet_cbebas, _ := jsonparser.GetFloat(value, "maxbet_cbebas")
		maxbuy_cbebas, _ := jsonparser.GetFloat(value, "maxbuy_cbebas")
		win_cbebas, _ := jsonparser.GetFloat(value, "win_cbebas")
		disc_cbebas, _ := jsonparser.GetFloat(value, "disc_cbebas")
		limitglobal_cbebas, _ := jsonparser.GetFloat(value, "limitglobal_cbebas")
		limittotal_cbebas, _ := jsonparser.GetFloat(value, "limittotal_cbebas")
		minbet_cmacau, _ := jsonparser.GetFloat(value, "minbet_cmacau")
		maxbet_cmacau, _ := jsonparser.GetFloat(value, "maxbet_cmacau")
		maxbuy_cmacau, _ := jsonparser.GetFloat(value, "maxbuy_cmacau")
		win2d_cmacau, _ := jsonparser.GetFloat(value, "win2d_cmacau")
		win3d_cmacau, _ := jsonparser.GetFloat(value, "win3d_cmacau")
		win4d_cmacau, _ := jsonparser.GetFloat(value, "win4d_cmacau")
		disc_cmacau, _ := jsonparser.GetFloat(value, "disc_cmacau")
		limitglobal_cmacau, _ := jsonparser.GetFloat(value, "limitglobal_cmacau")
		limitotal_cmacau, _ := jsonparser.GetFloat(value, "limitotal_cmacau")
		minbet_cnaga, _ := jsonparser.GetFloat(value, "minbet_cnaga")
		maxbet_cnaga, _ := jsonparser.GetFloat(value, "maxbet_cnaga")
		maxbuy_cnaga, _ := jsonparser.GetFloat(value, "maxbuy_cnaga")
		win3_cnaga, _ := jsonparser.GetFloat(value, "win3_cnaga")
		win4_cnaga, _ := jsonparser.GetFloat(value, "win4_cnaga")
		disc_cnaga, _ := jsonparser.GetFloat(value, "disc_cnaga")
		limitglobal_cnaga, _ := jsonparser.GetFloat(value, "limitglobal_cnaga")
		limittotal_cnaga, _ := jsonparser.GetFloat(value, "limittotal_cnaga")
		minbet_cjitu, _ := jsonparser.GetFloat(value, "minbet_cjitu")
		maxbet_cjitu, _ := jsonparser.GetFloat(value, "maxbet_cjitu")
		maxbuy_cjitu, _ := jsonparser.GetFloat(value, "maxbuy_cjitu")
		winas_cjitu, _ := jsonparser.GetFloat(value, "winas_cjitu")
		winkop_cjitu, _ := jsonparser.GetFloat(value, "winkop_cjitu")
		winkepala_cjitu, _ := jsonparser.GetFloat(value, "winkepala_cjitu")
		winekor_cjitu, _ := jsonparser.GetFloat(value, "winekor_cjitu")
		desc_cjitu, _ := jsonparser.GetFloat(value, "desc_cjitu")
		limitglobal_cjitu, _ := jsonparser.GetFloat(value, "limitglobal_cjitu")
		limittotal_cjitu, _ := jsonparser.GetFloat(value, "limittotal_cjitu")
		minbet_5050umum, _ := jsonparser.GetFloat(value, "minbet_5050umum")
		maxbet_5050umum, _ := jsonparser.GetFloat(value, "maxbet_5050umum")
		maxbuy_5050umum, _ := jsonparser.GetFloat(value, "maxbuy_5050umum")
		keibesar_5050umum, _ := jsonparser.GetFloat(value, "keibesar_5050umum")
		keikecil_5050umum, _ := jsonparser.GetFloat(value, "keikecil_5050umum")
		Keigenap_5050umum, _ := jsonparser.GetFloat(value, "keigenap_5050umum")
		keiganjil_5050umum, _ := jsonparser.GetFloat(value, "keiganjil_5050umum")
		keitengah_5050umum, _ := jsonparser.GetFloat(value, "keitengah_5050umum")
		keitepi_5050umum, _ := jsonparser.GetFloat(value, "keitepi_5050umum")
		discbesar_5050umum, _ := jsonparser.GetFloat(value, "discbesar_5050umum")
		disckecil_5050umum, _ := jsonparser.GetFloat(value, "disckecil_5050umum")
		discgenap_5050umum, _ := jsonparser.GetFloat(value, "discgenap_5050umum")
		discganjil_5050umum, _ := jsonparser.GetFloat(value, "discganjil_5050umum")
		disctengah_5050umum, _ := jsonparser.GetFloat(value, "disctengah_5050umum")
		disctepi_5050umum, _ := jsonparser.GetFloat(value, "disctepi_5050umum")
		limitglobal_5050umum, _ := jsonparser.GetFloat(value, "limitglobal_5050umum")
		limittotal_5050umum, _ := jsonparser.GetFloat(value, "limittotal_5050umum")
		minbet_5050special, _ := jsonparser.GetFloat(value, "minbet_5050special")
		maxbet_5050special, _ := jsonparser.GetFloat(value, "maxbet_5050special")
		maxbuy_5050special, _ := jsonparser.GetFloat(value, "maxbuy_5050special")
		keiasganjil_5050special, _ := jsonparser.GetFloat(value, "keiasganjil_5050special")
		keiasgenap_5050special, _ := jsonparser.GetFloat(value, "keiasgenap_5050special")
		keiasbesar_5050special, _ := jsonparser.GetFloat(value, "keiasbesar_5050special")
		keiaskecil_5050special, _ := jsonparser.GetFloat(value, "keiaskecil_5050special")
		keikopganjil_5050special, _ := jsonparser.GetFloat(value, "keikopganjil_5050special")
		keikopgenap_5050special, _ := jsonparser.GetFloat(value, "keikopgenap_5050special")
		keikopbesar_5050special, _ := jsonparser.GetFloat(value, "keikopbesar_5050special")
		keikopkecil_5050special, _ := jsonparser.GetFloat(value, "keikopkecil_5050special")
		keikepalaganjil_5050special, _ := jsonparser.GetFloat(value, "keikepalaganjil_5050special")
		keikepalagenap_5050special, _ := jsonparser.GetFloat(value, "keikepalagenap_5050special")
		keikepalabesar_5050special, _ := jsonparser.GetFloat(value, "keikepalabesar_5050special")
		keikepalakecil_5050special, _ := jsonparser.GetFloat(value, "keikepalakecil_5050special")
		keiekorganjil_5050special, _ := jsonparser.GetFloat(value, "keiekorganjil_5050special")
		keiekorgenap_5050special, _ := jsonparser.GetFloat(value, "keiekorgenap_5050special")
		keiekorbesar_5050special, _ := jsonparser.GetFloat(value, "keiekorbesar_5050special")
		keiekorkecil_5050special, _ := jsonparser.GetFloat(value, "keiekorkecil_5050special")
		discasganjil_5050special, _ := jsonparser.GetFloat(value, "discasganjil_5050special")
		discasgenap_5050special, _ := jsonparser.GetFloat(value, "discasgenap_5050special")
		discasbesar_5050special, _ := jsonparser.GetFloat(value, "discasbesar_5050special")
		discaskecil_5050special, _ := jsonparser.GetFloat(value, "discaskecil_5050special")
		disckopganjil_5050special, _ := jsonparser.GetFloat(value, "disckopganjil_5050special")
		disckopgenap_5050special, _ := jsonparser.GetFloat(value, "disckopgenap_5050special")
		disckopbesar_5050special, _ := jsonparser.GetFloat(value, "disckopbesar_5050special")
		disckopkecil_5050special, _ := jsonparser.GetFloat(value, "disckopkecil_5050special")
		disckepalaganjil_5050special, _ := jsonparser.GetFloat(value, "disckepalaganjil_5050special")
		disckepalagenap_5050special, _ := jsonparser.GetFloat(value, "disckepalagenap_5050special")
		disckepalabesar_5050special, _ := jsonparser.GetFloat(value, "disckepalabesar_5050special")
		disckepalakecil_5050special, _ := jsonparser.GetFloat(value, "disckepalakecil_5050special")
		discekorganjil_5050special, _ := jsonparser.GetFloat(value, "discekorganjil_5050special")
		discekorgenap_5050special, _ := jsonparser.GetFloat(value, "discekorgenap_5050special")
		discekorbesar_5050special, _ := jsonparser.GetFloat(value, "discekorbesar_5050special")
		discekorkecil_5050special, _ := jsonparser.GetFloat(value, "discekorkecil_5050special")
		limitglobal_5050special, _ := jsonparser.GetFloat(value, "limitglobal_5050special")
		limittotal_5050special, _ := jsonparser.GetFloat(value, "limittotal_5050special")
		minbet_5050kombinasi, _ := jsonparser.GetFloat(value, "minbet_5050kombinasi")
		maxbet_5050kombinasi, _ := jsonparser.GetFloat(value, "maxbet_5050kombinasi")
		maxbuy_5050kombinasi, _ := jsonparser.GetFloat(value, "maxbuy_5050kombinasi")
		belakangkeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeimono_5050kombinasi")
		belakangkeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeistereo_5050kombinasi")
		belakangkeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikembang_5050kombinasi")
		belakangkeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikempis_5050kombinasi")
		belakangkeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangkeikembar_5050kombinasi")
		tengahkeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeimono_5050kombinasi")
		tengahkeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeistereo_5050kombinasi")
		tengahkeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikembang_5050kombinasi")
		tengahkeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikempis_5050kombinasi")
		tengahkeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahkeikembar_5050kombinasi")
		depankeimono_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeimono_5050kombinasi")
		depankeistereo_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeistereo_5050kombinasi")
		depankeikembang_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikembang_5050kombinasi")
		depankeikempis_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikempis_5050kombinasi")
		depankeikembar_5050kombinasi, _ := jsonparser.GetFloat(value, "depankeikembar_5050kombinasi")
		belakangdiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdiscmono_5050kombinasi")
		belakangdiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdiscstereo_5050kombinasi")
		belakangdisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckembang_5050kombinasi")
		belakangdisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckempis_5050kombinasi")
		belakangdisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "belakangdisckembar_5050kombinasi")
		tengahdiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdiscmono_5050kombinasi")
		tengahdiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdiscstereo_5050kombinasi")
		tengahdisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckembang_5050kombinasi")
		tengahdisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckempis_5050kombinasi")
		tengahdisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "tengahdisckembar_5050kombinasi")
		depandiscmono_5050kombinasi, _ := jsonparser.GetFloat(value, "depandiscmono_5050kombinasi")
		depandiscstereo_5050kombinasi, _ := jsonparser.GetFloat(value, "depandiscstereo_5050kombinasi")
		depandisckembang_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckembang_5050kombinasi")
		depandisckempis_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckempis_5050kombinasi")
		depandisckembar_5050kombinasi, _ := jsonparser.GetFloat(value, "depandisckembar_5050kombinasi")
		limitglobal_5050kombinasi, _ := jsonparser.GetFloat(value, "limitglobal_5050kombinasi")
		limittotal_5050kombinasi, _ := jsonparser.GetFloat(value, "limittotal_5050kombinasi")
		minbet_kombinasi, _ := jsonparser.GetFloat(value, "minbet_kombinasi")
		maxbet_kombinasi, _ := jsonparser.GetFloat(value, "maxbet_kombinasi")
		maxbuy_kombinasi, _ := jsonparser.GetFloat(value, "maxbuy_kombinasi")
		win_kombinasi, _ := jsonparser.GetFloat(value, "win_kombinasi")
		disc_kombinasi, _ := jsonparser.GetFloat(value, "disc_kombinasi")
		limitglobal_kombinasi, _ := jsonparser.GetFloat(value, "limitglobal_kombinasi")
		limittotal_kombinasi, _ := jsonparser.GetFloat(value, "limittotal_kombinasi")
		minbet_dasar, _ := jsonparser.GetFloat(value, "minbet_dasar")
		maxbet_dasar, _ := jsonparser.GetFloat(value, "maxbet_dasar")
		maxbuy_dasar, _ := jsonparser.GetFloat(value, "maxbuy_dasar")
		keibesar_dasar, _ := jsonparser.GetFloat(value, "keibesar_dasar")
		keikecil_dasar, _ := jsonparser.GetFloat(value, "keikecil_dasar")
		keigenap_dasar, _ := jsonparser.GetFloat(value, "keigenap_dasar")
		keiganjil_dasar, _ := jsonparser.GetFloat(value, "keiganjil_dasar")
		discbesar_dasar, _ := jsonparser.GetFloat(value, "discbesar_dasar")
		disckecil_dasar, _ := jsonparser.GetFloat(value, "disckecil_dasar")
		discgenap_dasar, _ := jsonparser.GetFloat(value, "discgenap_dasar")
		discganjil_dasar, _ := jsonparser.GetFloat(value, "discganjil_dasar")
		limitglobal_dasar, _ := jsonparser.GetFloat(value, "limitglobal_dasar")
		limittotal_dasar, _ := jsonparser.GetFloat(value, "limittotal_dasar")
		minbet_shio, _ := jsonparser.GetFloat(value, "minbet_shio")
		maxbet_shio, _ := jsonparser.GetFloat(value, "maxbet_shio")
		maxbuy_shio, _ := jsonparser.GetFloat(value, "maxbuy_shio")
		win_shio, _ := jsonparser.GetFloat(value, "win_shio")
		disc_shio, _ := jsonparser.GetFloat(value, "disc_shio")
		shioyear_shio, _ := jsonparser.GetString(value, "shioyear_shio")
		limitglobal_shio, _ := jsonparser.GetFloat(value, "limitglobal_shio")
		limittotal_shio, _ := jsonparser.GetFloat(value, "limittotal_shio")

		obj.Limitline4d = int(limitline_4d)
		obj.Limitline3d = int(limitline_3d)
		obj.Limitline3dd = int(limitline_3dd)
		obj.Limitline2d = int(limitline_2d)
		obj.Limitline2dd = int(limitline_2dd)
		obj.Limitline2dt = int(limitline_2dt)
		obj.Bbfs = int(bbfs)
		obj.Minbet_432d = float32(minbet_432d)
		obj.Maxbet4d_432d = float32(maxbet4d_432d)
		obj.Maxbet3d_432d = float32(maxbet3d_432d)
		obj.Maxbet3dd_432d = float32(maxbet3dd_432d)
		obj.Maxbet2d_432d = float32(maxbet2d_432d)
		obj.Maxbet2dd_432d = float32(maxbet2dd_432d)
		obj.Maxbet2dt_432d = float32(maxbet2dt_432d)
		obj.Maxbet4dfullbb_432d = float32(maxbet4dfullbb_432d)
		obj.Maxbet3dfullbb_432d = float32(maxbet3dfullbb_432d)
		obj.Maxbet3ddfullbb_432d = float32(maxbet3ddfullbb_432d)
		obj.Maxbet2dfullbb_432d = float32(maxbet2dfullbb_432d)
		obj.Maxbet2ddfullbb_432d = float32(maxbet2ddfullbb_432d)
		obj.Maxbet2dtfullbb_432d = float32(maxbet2dtfullbb_432d)
		obj.Maxbuy4d_432d = float32(Maxbuy4d_432d)
		obj.Maxbuy3d_432d = float32(Maxbuy3d_432d)
		obj.Maxbuy3dd_432d = float32(Maxbuy3dd_432d)
		obj.Maxbuy2d_432d = float32(Maxbuy2d_432d)
		obj.Maxbuy2dd_432d = float32(Maxbuy2dd_432d)
		obj.Maxbuy2dt_432d = float32(Maxbuy2dt_432d)
		obj.Limitotal4d_432d = float32(limitotal4d_432d)
		obj.Limitotal3d_432d = float32(limitotal3d_432d)
		obj.Limitotal3dd_432d = float32(limitotal3dd_432d)
		obj.Limitotal2d_432d = float32(limitotal2d_432d)
		obj.Limitotal2dd_432d = float32(limitotal2dd_432d)
		obj.Limitotal2dt_432d = float32(limitotal2dt_432d)
		obj.Limitglobal4d_432d = float32(limitglobal4d_432d)
		obj.Limitglobal3d_432d = float32(limitglobal3d_432d)
		obj.Limitglobal3dd_432d = float32(limitglobal3dd_432d)
		obj.Limitglobal2d_432d = float32(limitglobal2d_432d)
		obj.Limitglobal2dd_432d = float32(limitglobal2dd_432d)
		obj.Limitglobal2dt_432d = float32(limitglobal2dt_432d)
		obj.Limitglobal4d_fullbb_432d = float32(limitotal4d_fullbb_432d)
		obj.Limitotal3d_fullbb_432d = float32(limitotal3d_fullbb_432d)
		obj.Limitotal3dd_fullbb_432d = float32(limitotal3dd_fullbb_432d)
		obj.Limitotal2d_fullbb_432d = float32(limitotal2d_fullbb_432d)
		obj.Limitotal2dd_fullbb_432d = float32(limitotal2dd_fullbb_432d)
		obj.Limitotal2dt_fullbb_432d = float32(limitotal2dt_fullbb_432d)
		obj.Limitglobal4d_fullbb_432d = float32(limitglobal4d_fullbb_432d)
		obj.Limitglobal3d_fullbb_432d = float32(limitglobal3d_fullbb_432d)
		obj.Limitglobal3dd_fullbb_432d = float32(limitglobal3dd_fullbb_432d)
		obj.Limitglobal2d_fullbb_432d = float32(limitglobal2d_fullbb_432d)
		obj.Limitglobal2dd_fullbb_432d = float32(limitglobal2dd_fullbb_432d)
		obj.Limitglobal2dt_fullbb_432d = float32(limitglobal2dt_fullbb_432d)
		obj.Disc4d_432d = float32(disc4d_432d)
		obj.Disc3d_432d = float32(disc3d_432d)
		obj.Disc3dd_432d = float32(disc3dd_432d)
		obj.Disc2d_432d = float32(disc2d_432d)
		obj.Disc2dd_432d = float32(disc2dd_432d)
		obj.Disc2dt_432d = float32(disc2dt_432d)
		obj.Win4d_432d = float32(win4d_432d)
		obj.Win3d_432d = float32(win3d_432d)
		obj.Win3dd_432d = float32(win3dd_432d)
		obj.Win2d_432d = float32(win2d_432d)
		obj.Win2dd_432d = float32(win2dd_432d)
		obj.Win2dt_432d = float32(win2dt_432d)
		obj.Win4dnodisc_432d = float32(win4dnodisc_432d)
		obj.Win3dnodisc_432d = float32(win3dnodisc_432d)
		obj.Win3ddnodisc_432d = float32(win3ddnodisc_432d)
		obj.Win2dnodisc_432d = float32(win2dnodisc_432d)
		obj.Win2ddnodisc_432d = float32(win2ddnodisc_432d)
		obj.Win2dtnodisc_432d = float32(win2dtnodisc_432d)
		obj.Win4dbb_kena_432d = float32(win4dbb_kena_432d)
		obj.Win3dbb_kena_432d = float32(win3dbb_kena_432d)
		obj.Win3ddbb_kena_432d = float32(win3ddbb_kena_432d)
		obj.Win2dbb_kena_432d = float32(win2dbb_kena_432d)
		obj.Win2ddbb_kena_432d = float32(win2ddbb_kena_432d)
		obj.Win2dtbb_kena_432d = float32(win2dtbb_kena_432d)
		obj.Win4dbb_432d = float32(win4dbb_432d)
		obj.Win3dbb_432d = float32(win3dbb_432d)
		obj.Win3ddbb_432d = float32(win3ddbb_432d)
		obj.Win2dbb_432d = float32(win2dbb_432d)
		obj.Win2ddbb_432d = float32(win2ddbb_432d)
		obj.Win2dtbb_432d = float32(win2dtbb_432d)
		obj.Minbet_cbebas = float32(minbet_cbebas)
		obj.Maxbet_cbebas = float32(maxbet_cbebas)
		obj.Maxbuy_cbebas = float32(maxbuy_cbebas)
		obj.Win_cbebas = float32(win_cbebas)
		obj.Disc_cbebas = float32(disc_cbebas)
		obj.Limitglobal_cbebas = float32(limitglobal_cbebas)
		obj.Limittotal_cbebas = float32(limittotal_cbebas)
		obj.Minbet_cmacau = float32(minbet_cmacau)
		obj.Maxbet_cmacau = float32(maxbet_cmacau)
		obj.Maxbuy_cmacau = float32(maxbuy_cmacau)
		obj.Win2d_cmacau = float32(win2d_cmacau)
		obj.Win3d_cmacau = float32(win3d_cmacau)
		obj.Win4d_cmacau = float32(win4d_cmacau)
		obj.Disc_cmacau = float32(disc_cmacau)
		obj.Limitglobal_cmacau = float32(limitglobal_cmacau)
		obj.Limitotal_cmacau = float32(limitotal_cmacau)
		obj.Minbet_cnaga = float32(minbet_cnaga)
		obj.Maxbet_cnaga = float32(maxbet_cnaga)
		obj.Maxbuy_cnaga = float32(maxbuy_cnaga)
		obj.Win3_cnaga = float32(win3_cnaga)
		obj.Win4_cnaga = float32(win4_cnaga)
		obj.Disc_cnaga = float32(disc_cnaga)
		obj.Limitglobal_cnaga = float32(limitglobal_cnaga)
		obj.Limittotal_cnaga = float32(limittotal_cnaga)
		obj.Minbet_cjitu = float32(minbet_cjitu)
		obj.Maxbet_cjitu = float32(maxbet_cjitu)
		obj.Maxbuy_cjitu = float32(maxbuy_cjitu)
		obj.Winas_cjitu = float32(winas_cjitu)
		obj.Winkop_cjitu = float32(winkop_cjitu)
		obj.Winkepala_cjitu = float32(winkepala_cjitu)
		obj.Winekor_cjitu = float32(winekor_cjitu)
		obj.Desc_cjitu = float32(desc_cjitu)
		obj.Limitglobal_cjitu = float32(limitglobal_cjitu)
		obj.Limittotal_cjitu = float32(limittotal_cjitu)
		obj.Minbet_5050umum = float32(minbet_5050umum)
		obj.Maxbet_5050umum = float32(maxbet_5050umum)
		obj.Maxbuy_5050umum = float32(maxbuy_5050umum)
		obj.Keibesar_5050umum = float32(keibesar_5050umum)
		obj.Keikecil_5050umum = float32(keikecil_5050umum)
		obj.Keigenap_5050umum = float32(Keigenap_5050umum)
		obj.Keiganjil_5050umum = float32(keiganjil_5050umum)
		obj.Keitengah_5050umum = float32(keitengah_5050umum)
		obj.Keitepi_5050umum = float32(keitepi_5050umum)
		obj.Discbesar_5050umum = float32(discbesar_5050umum)
		obj.Disckecil_5050umum = float32(disckecil_5050umum)
		obj.Discgenap_5050umum = float32(discgenap_5050umum)
		obj.Discganjil_5050umum = float32(discganjil_5050umum)
		obj.Disctengah_5050umum = float32(disctengah_5050umum)
		obj.Disctepi_5050umum = float32(disctepi_5050umum)
		obj.Limitglobal_5050umum = float32(limitglobal_5050umum)
		obj.Limittotal_5050umum = float32(limittotal_5050umum)
		obj.Minbet_5050special = float32(minbet_5050special)
		obj.Maxbet_5050special = float32(maxbet_5050special)
		obj.Maxbuy_5050special = float32(maxbuy_5050special)
		obj.Keiasganjil_5050special = float32(keiasganjil_5050special)
		obj.Keiasgenap_5050special = float32(keiasgenap_5050special)
		obj.Keiasbesar_5050special = float32(keiasbesar_5050special)
		obj.Keiaskecil_5050special = float32(keiaskecil_5050special)
		obj.Keikopganjil_5050special = float32(keikopganjil_5050special)
		obj.Keikopgenap_5050special = float32(keikopgenap_5050special)
		obj.Keikopbesar_5050special = float32(keikopbesar_5050special)
		obj.Keikopkecil_5050special = float32(keikopkecil_5050special)
		obj.Keikepalaganjil_5050special = float32(keikepalaganjil_5050special)
		obj.Keikepalagenap_5050special = float32(keikepalagenap_5050special)
		obj.Keikepalabesar_5050special = float32(keikepalabesar_5050special)
		obj.Keikepalakecil_5050special = float32(keikepalakecil_5050special)
		obj.Keiekorganjil_5050special = float32(keiekorganjil_5050special)
		obj.Keiekorgenap_5050special = float32(keiekorgenap_5050special)
		obj.Keiekorbesar_5050special = float32(keiekorbesar_5050special)
		obj.Keiekorkecil_5050special = float32(keiekorkecil_5050special)
		obj.Discasganjil_5050special = float32(discasganjil_5050special)
		obj.Discasgenap_5050special = float32(discasgenap_5050special)
		obj.Discasbesar_5050special = float32(discasbesar_5050special)
		obj.Discaskecil_5050special = float32(discaskecil_5050special)
		obj.Disckopganjil_5050special = float32(disckopganjil_5050special)
		obj.Disckopgenap_5050special = float32(disckopgenap_5050special)
		obj.Disckopbesar_5050special = float32(disckopbesar_5050special)
		obj.Disckopkecil_5050special = float32(disckopkecil_5050special)
		obj.Disckepalaganjil_5050special = float32(disckepalaganjil_5050special)
		obj.Disckepalagenap_5050special = float32(disckepalagenap_5050special)
		obj.Disckepalabesar_5050special = float32(disckepalabesar_5050special)
		obj.Disckepalakecil_5050special = float32(disckepalakecil_5050special)
		obj.Discekorganjil_5050special = float32(discekorganjil_5050special)
		obj.Discekorgenap_5050special = float32(discekorgenap_5050special)
		obj.Discekorbesar_5050special = float32(discekorbesar_5050special)
		obj.Discekorkecil_5050special = float32(discekorkecil_5050special)
		obj.Limitglobal_5050special = float32(limitglobal_5050special)
		obj.Limittotal_5050special = float32(limittotal_5050special)
		obj.Minbet_5050kombinasi = float32(minbet_5050kombinasi)
		obj.Maxbet_5050kombinasi = float32(maxbet_5050kombinasi)
		obj.Maxbuy_5050kombinasi = float32(maxbuy_5050kombinasi)
		obj.Belakangkeimono_5050kombinasi = float32(belakangkeimono_5050kombinasi)
		obj.Belakangkeistereo_5050kombinasi = float32(belakangkeistereo_5050kombinasi)
		obj.Belakangkeikembang_5050kombinasi = float32(belakangkeikembang_5050kombinasi)
		obj.Belakangkeikempis_5050kombinasi = float32(belakangkeikempis_5050kombinasi)
		obj.Belakangkeikembar_5050kombinasi = float32(belakangkeikembar_5050kombinasi)
		obj.Tengahkeimono_5050kombinasi = float32(tengahkeimono_5050kombinasi)
		obj.Tengahkeistereo_5050kombinasi = float32(tengahkeistereo_5050kombinasi)
		obj.Tengahkeikembang_5050kombinasi = float32(tengahkeikembang_5050kombinasi)
		obj.Tengahkeikempis_5050kombinasi = float32(tengahkeikempis_5050kombinasi)
		obj.Tengahkeikembar_5050kombinasi = float32(tengahkeikembar_5050kombinasi)
		obj.Depankeimono_5050kombinasi = float32(depankeimono_5050kombinasi)
		obj.Depankeistereo_5050kombinasi = float32(depankeistereo_5050kombinasi)
		obj.Depankeikembang_5050kombinasi = float32(depankeikembang_5050kombinasi)
		obj.Depankeikempis_5050kombinasi = float32(depankeikempis_5050kombinasi)
		obj.Depankeikembar_5050kombinasi = float32(depankeikembar_5050kombinasi)
		obj.Belakangdiscmono_5050kombinasi = float32(belakangdiscmono_5050kombinasi)
		obj.Belakangdiscstereo_5050kombinasi = float32(belakangdiscstereo_5050kombinasi)
		obj.Belakangdisckembang_5050kombinasi = float32(belakangdisckembang_5050kombinasi)
		obj.Belakangdisckempis_5050kombinasi = float32(belakangdisckempis_5050kombinasi)
		obj.Belakangdisckembar_5050kombinasi = float32(belakangdisckembar_5050kombinasi)
		obj.Tengahdiscmono_5050kombinasi = float32(tengahdiscmono_5050kombinasi)
		obj.Tengahdiscstereo_5050kombinasi = float32(tengahdiscstereo_5050kombinasi)
		obj.Tengahdisckembang_5050kombinasi = float32(tengahdisckembang_5050kombinasi)
		obj.Tengahdisckempis_5050kombinasi = float32(tengahdisckempis_5050kombinasi)
		obj.Tengahdisckembar_5050kombinasi = float32(tengahdisckembar_5050kombinasi)
		obj.Depandiscmono_5050kombinasi = float32(depandiscmono_5050kombinasi)
		obj.Depandiscstereo_5050kombinasi = float32(depandiscstereo_5050kombinasi)
		obj.Depandisckembang_5050kombinasi = float32(depandisckembang_5050kombinasi)
		obj.Depandisckempis_5050kombinasi = float32(depandisckempis_5050kombinasi)
		obj.Depandisckembar_5050kombinasi = float32(depandisckembar_5050kombinasi)
		obj.Limitglobal_5050kombinasi = float32(limitglobal_5050kombinasi)
		obj.Limittotal_5050kombinasi = float32(limittotal_5050kombinasi)
		obj.Minbet_kombinasi = float32(minbet_kombinasi)
		obj.Maxbet_kombinasi = float32(maxbet_kombinasi)
		obj.Maxbuy_kombinasi = float32(maxbuy_kombinasi)
		obj.Win_kombinasi = float32(win_kombinasi)
		obj.Disc_kombinasi = float32(disc_kombinasi)
		obj.Limitglobal_kombinasi = float32(limitglobal_kombinasi)
		obj.Limittotal_kombinasi = float32(limittotal_kombinasi)
		obj.Minbet_dasar = float32(minbet_dasar)
		obj.Maxbet_dasar = float32(maxbet_dasar)
		obj.Maxbuy_dasar = float32(maxbuy_dasar)
		obj.Keibesar_dasar = float32(keibesar_dasar)
		obj.Keikecil_dasar = float32(keikecil_dasar)
		obj.Keigenap_dasar = float32(keigenap_dasar)
		obj.Keiganjil_dasar = float32(keiganjil_dasar)
		obj.Discbesar_dasar = float32(discbesar_dasar)
		obj.Disckecil_dasar = float32(disckecil_dasar)
		obj.Discgenap_dasar = float32(discgenap_dasar)
		obj.Discganjil_dasar = float32(discganjil_dasar)
		obj.Limitglobal_dasar = float32(limitglobal_dasar)
		obj.Limittotal_dasar = float32(limittotal_dasar)
		obj.Minbet_shio = float32(minbet_shio)
		obj.Maxbet_shio = float32(maxbet_shio)
		obj.Maxbuy_shio = float32(maxbuy_shio)
		obj.Win_shio = float32(win_shio)
		obj.Disc_shio = float32(disc_shio)
		obj.Shioyear_shio = string(shioyear_shio)
		obj.Limitglobal_shio = float32(limitglobal_shio)
		obj.Limittotal_shio = float32(limittotal_shio)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaranDetailConf(client.Pasarancode)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(FieldpasarandetailCONF_home_redis+"_"+client.Pasarancode, result, 30*time.Minute)
		log.Println("PASARAN CONF MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN CONF CACHE")

		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PasaranSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransave)
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

	result, err := models.Save_pasaran(client.Sdata, client.Master, client.Idrecord, client.Name, client.Tipe, client.Url, client.Diundi, client.Jamtutup, client.Jamjadwal, client.Jamopen)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveLimitLine(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransavelimitline)
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
		client.Pasaran_limitline4d, client.Pasaran_limitline3d, client.Pasaran_limitline3dd, client.Pasaran_limitline2d,
		client.Pasaran_limitline2dd, client.Pasaran_limitline2dt, client.Pasaran_bbfs)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConf432d(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconf432d)
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
		client.Pasaran_minbet_432d, client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet3dd_432d, client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
		client.Pasaran_maxbet4dfullbb_432d, client.Pasaran_maxbet3dfullbb_432d, client.Pasaran_maxbet3ddfullbb_432d,
		client.Pasaran_maxbet2dfullbb_432d, client.Pasaran_maxbet2ddfullbb_432d, client.Pasaran_maxbet2dtfullbb_432d,
		client.Pasaran_maxbuy4d_432d, client.Pasaran_maxbuy3d_432d, client.Pasaran_maxbuy3dd_432d,
		client.Pasaran_maxbuy2d_432d, client.Pasaran_maxbuy2dd_432d, client.Pasaran_maxbuy2dt_432d,
		client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win3dd_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
		client.Pasaran_win4dnodisc_432d, client.Pasaran_win3dnodisc_432d, client.Pasaran_win3ddnodisc_432d, client.Pasaran_win2dnodisc_432d, client.Pasaran_win2ddnodisc_432d, client.Pasaran_win2dtnodisc_432d,
		client.Pasaran_win4dbb_kena_432d, client.Pasaran_win3dbb_kena_432d, client.Pasaran_win3ddbb_kena_432d, client.Pasaran_win2dbb_kena_432d, client.Pasaran_win2ddbb_kena_432d, client.Pasaran_win2dtbb_kena_432d,
		client.Pasaran_win4dbb_432d, client.Pasaran_win3dbb_432d, client.Pasaran_win3ddbb_432d, client.Pasaran_win2dbb_432d, client.Pasaran_win2ddbb_432d, client.Pasaran_win2dtbb_432d,
		client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc3dd_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
		client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal3dd_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
		client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal3dd_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d,
		client.Pasaran_limitglobal4d_fullbb_432d, client.Pasaran_limitglobal3d_fullbb_432d, client.Pasaran_limitglobal3dd_fullbb_432d, client.Pasaran_limitglobal2d_fullbb_432d, client.Pasaran_limitglobal2dd_fullbb_432d, client.Pasaran_limitglobal2dt_fullbb_432d,
		client.Pasaran_limitotal4d_fullbb_432d, client.Pasaran_limitotal3d_fullbb_432d, client.Pasaran_limitotal3dd_fullbb_432d, client.Pasaran_limitotal2d_fullbb_432d, client.Pasaran_limitotal2dd_fullbb_432d, client.Pasaran_limitotal2dt_fullbb_432d)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfColokBebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfcolokbebas)
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
		client.Pasaran_minbet_cbebas, client.Pasaran_maxbet_cbebas, client.Pasaran_maxbuy_cbebas,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfColokMacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfcolokmacau)
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
		client.Pasaran_minbet_cmacau, client.Pasaran_maxbet_cmacau, client.Pasaran_maxbuy_cmacau,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfColokNaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfcoloknaga)
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
		client.Pasaran_minbet_cnaga, client.Pasaran_maxbet_cnaga, client.Pasaran_maxbuy_cnaga,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfColokJitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfcolokjitu)
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
		client.Pasaran_minbet_cjitu, client.Pasaran_maxbet_cjitu, client.Pasaran_maxbuy_cjitu,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConf5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconf5050umum)
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
		client.Pasaran_minbet_5050umum, client.Pasaran_maxbet_5050umum, client.Pasaran_maxbuy_5050umum,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConf5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconf5050special)
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
		client.Pasaran_minbet_5050special, client.Pasaran_maxbet_5050special, client.Pasaran_maxbuy_5050special,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConf5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconf5050kombinasi)
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
		client.Pasaran_minbet_5050kombinasi, client.Pasaran_maxbet_5050kombinasi, client.Pasaran_maxbuy_5050kombinasi,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfmacaukombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfmacaukombinasi)
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
		client.Pasaran_minbet_kombinasi, client.Pasaran_maxbet_kombinasi, client.Pasaran_maxbuy_kombinasi,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfdasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfdasar)
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
		client.Pasaran_minbet_dasar, client.Pasaran_maxbet_dasar, client.Pasaran_maxbuy_dasar,
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
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}
func PasaranSaveConfshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaransaveconfshio)
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
		client.Pasaran_minbet_shio, client.Pasaran_maxbet_shio, client.Pasaran_maxbuy_shio,
		client.Pasaran_win_shio, client.Pasaran_disc_shio, client.Pasaran_limitglobal_shio, client.Pasaran_limittotal_shio)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_pasaran(client.Idrecord)
	return c.JSON(result)
}

func _deleteredis_pasaran(idrecord string) {
	_deleteredis_log()
	//MASTER
	val_master := helpers.DeleteRedis(Fieldpasaran_home_redis)
	log.Printf("Redis Delete MASTER PASARAN : %d", val_master)
	val_master_detail := helpers.DeleteRedis(Fieldpasarandetail_home_redis + "_" + idrecord)
	log.Printf("Redis Delete MASTER DETAIL PASARAN : %d", val_master_detail)
	val_master_conf := helpers.DeleteRedis(FieldpasarandetailCONF_home_redis + "_" + idrecord)
	log.Printf("Redis Delete MASTER DETAIL CONF PASARAN : %d", val_master_conf)

}
