package controller

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
)

const Fieldcompany_home_redis = "LISTCOMPANY_MASTER"
const Fieldcompanydetail_home_redis = "LISTCOMPANYDETAIL_MASTER"
const Fieldcompanylistadmin_home_redis = "LISTCOMPANYLISTADMIN_MASTER"
const Fieldcompanylistpasaran_home_redis = "LISTCOMPANYLISTPASARAN_MASTER"
const Fieldcompanylistpasaranonline_home_redis = "LISTCOMPANYLISTPASARANONLINE_MASTER"
const Fieldcompanylistpasaranconf_home_redis = "LISTCOMPANYLISTPASARANCONF_MASTER"
const Fieldcompanylistpasarankeluaran_home_redis = "LISTCOMPANYLISTPASARANKELUARAN_MASTER"
const Fieldcompanyinvoicelist_home_redis = "LISTCOMPANYINVOICELISTMEMBER_MASTER"
const Fieldcompanyinvoicelisttemp_home_redis = "LISTCOMPANYINVOICELISTTEMP_MASTER"
const Fieldcompanyinvoicelistpermainan_home_redis = "LISTCOMPANYINVOICELISTPERMAINAN_MASTER"

func CompanyHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_company)
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
	var obj entities.Model_company
	var arraobj []entities.Model_company
	resultredis, flag := helpers.GetRedis(Fieldcompany_home_redis)
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

		obj.Company_no = int(company_no)
		obj.Company_idcompany = company_idcompany
		obj.Company_startjoin = company_startjoin
		obj.Company_endjoin = company_endjoin
		obj.Company_curr = company_curr
		obj.Company_name = company_name
		obj.Company_periode = company_periode
		obj.Company_winlose = int(company_winlose)
		obj.Company_winlosetemp = int(company_winlosetemp)
		obj.Company_status = company_status
		obj.Company_statuscss = company_statuscss
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompany_home_redis, result, 60*time.Minute)
		log.Println("COMPANY MYSQL")
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
	client := new(entities.Controller_companydetail)
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
	var obj entities.Model_companydetail
	var arraobj []entities.Model_companydetail
	resultredis, flag := helpers.GetRedis(Fieldcompanydetail_home_redis + "_" + client.Company)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_name, _ := jsonparser.GetString(value, "company_name")
		company_url, _ := jsonparser.GetString(value, "company_url")
		company_status, _ := jsonparser.GetString(value, "company_status")
		company_create, _ := jsonparser.GetString(value, "company_create")
		company_update, _ := jsonparser.GetString(value, "company_update")

		obj.Company_name = company_name
		obj.Company_url = company_url
		obj.Company_status = company_status
		obj.Company_create = company_create
		obj.Company_update = company_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_companyDetail(client.Company)
		helpers.SetRedis(Fieldcompanydetail_home_redis+"_"+client.Company, result, 20*time.Minute)
		log.Println("COMPANY DETAIL MYSQL " + client.Company)
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
		log.Println("COMPANY DETAIL CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyDetailListAdmin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companydetail)
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
	var obj entities.Model_companylistadmin
	var arraobj []entities.Model_companylistadmin
	resultredis, flag := helpers.GetRedis(Fieldcompanylistadmin_home_redis + "_" + client.Company)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_admin_username, _ := jsonparser.GetString(value, "company_admin_username")
		company_admin_typeadmin, _ := jsonparser.GetString(value, "company_admin_typeadmin")
		company_admin_nama, _ := jsonparser.GetString(value, "company_admin_nama")
		company_admin_status, _ := jsonparser.GetString(value, "company_admin_status")
		company_admin_status_css, _ := jsonparser.GetString(value, "company_admin_status_css")
		company_admin_lastlogin, _ := jsonparser.GetString(value, "company_admin_lastlogin")
		Company_admin_lastippadress, _ := jsonparser.GetString(value, "Company_admin_lastippadress")
		company_admin_create, _ := jsonparser.GetString(value, "company_admin_create")
		company_admin_update, _ := jsonparser.GetString(value, "company_admin_update")

		obj.Company_admin_username = company_admin_username
		obj.Company_admin_typeadmin = company_admin_typeadmin
		obj.Company_admin_name = company_admin_nama
		obj.Company_admin_status = company_admin_status
		obj.Company_admin_statuscss = company_admin_status_css
		obj.Company_admin_lastlogin = company_admin_lastlogin
		obj.Company_admin_lastippadress = Company_admin_lastippadress
		obj.Company_admin_create = company_admin_create
		obj.Company_admin_update = company_admin_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_listadmin(client.Company)
		helpers.SetRedis(Fieldcompanylistadmin_home_redis+"_"+client.Company, result, 20*time.Minute)
		log.Println("COMPANY LISTADMIN MYSQL " + client.Company)
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
		log.Println("COMPANY LISTADMIN CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyDetailListPasaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companydetail)
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
	var obj entities.Model_companylistpasaran
	var arraobj []entities.Model_companylistpasaran
	resultredis, flag := helpers.GetRedis(Fieldcompanylistpasaran_home_redis + "_" + client.Company)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_pasaran_idcomppasaran, _ := jsonparser.GetInt(value, "company_pasaran_idcomppasaran")
		company_pasaran_idpasarantogel, _ := jsonparser.GetString(value, "company_pasaran_idpasarantogel")
		company_pasaran_nmpasarantogel, _ := jsonparser.GetString(value, "company_pasaran_nmpasarantogel")
		company_pasaran_periode, _ := jsonparser.GetString(value, "company_pasaran_periode")
		company_pasaran_winlose, _ := jsonparser.GetInt(value, "company_pasaran_winlose")
		company_pasaran_displaypasaran, _ := jsonparser.GetInt(value, "company_pasaran_displaypasaran")
		company_pasaran_status, _ := jsonparser.GetString(value, "company_pasaran_status")
		company_pasaran_statuscss, _ := jsonparser.GetString(value, "company_pasaran_statuscss")
		company_pasaran_statuspasaranactive, _ := jsonparser.GetString(value, "company_pasaran_statuspasaranactive")
		company_pasaran_statuspasaranactivecss, _ := jsonparser.GetString(value, "company_pasaran_statuspasaranactivecss")

		obj.Company_pasaran_idcomppasaran = int(company_pasaran_idcomppasaran)
		obj.Company_pasaran_idpasarantogel = company_pasaran_idpasarantogel
		obj.Company_pasaran_nmpasarantogel = company_pasaran_nmpasarantogel
		obj.Company_pasaran_periode = company_pasaran_periode
		obj.Company_pasaran_winlose = int(company_pasaran_winlose)
		obj.Company_pasaran_displaypasaran = int(company_pasaran_displaypasaran)
		obj.Company_pasaran_status = company_pasaran_status
		obj.Company_pasaran_statuscss = company_pasaran_statuscss
		obj.Company_pasaran_statusactive = company_pasaran_statuspasaranactive
		obj.Company_pasaran_statusactivecss = company_pasaran_statuspasaranactivecss
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_listpasaran(client.Company)
		helpers.SetRedis(Fieldcompanylistpasaran_home_redis+"_"+client.Company, result, 20*time.Minute)
		log.Println("COMPANY LISTPASARAN MYSQL " + client.Company)
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
		log.Println("COMPANY LISTPASARAN CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyDetailListPasaranOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaranconf)
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
	var obj entities.Model_companylistpasaranonline
	var arraobj []entities.Model_companylistpasaranonline
	resultredis, flag := helpers.GetRedis(Fieldcompanylistpasaranonline_home_redis + "_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_pasaranonline_id, _ := jsonparser.GetInt(value, "company_pasaranonline_id")
		company_pasaranonline_hari, _ := jsonparser.GetString(value, "company_pasaranonline_hari")

		obj.Company_pasaran_onlineid = int(company_pasaranonline_id)
		obj.Company_pasaran_harian = company_pasaranonline_hari
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_listpasaranonline(client.Company, client.Companypasaran_id)
		helpers.SetRedis(Fieldcompanylistpasaranonline_home_redis+"_"+client.Company+"_"+strconv.Itoa(client.Companypasaran_id), result, 20*time.Minute)
		log.Println("COMPANY LISTPASARAN ONLINE MYSQL " + client.Company)
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
		log.Println("COMPANY LISTPASARAN ONLINE CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyDetailListPasaranConf(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaranconf)
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
	var obj entities.Model_companypasaranconf
	var arraobj []entities.Model_companypasaranconf
	resultredis, flag := helpers.GetRedis(Fieldcompanylistpasaranconf_home_redis + "_" + client.Company + "_" + strconv.Itoa(client.Companypasaran_id))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_diundi, _ := jsonparser.GetString(value, "pasaran_diundi")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")
		pasaran_jamtutup, _ := jsonparser.GetString(value, "pasaran_jamtutup")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_jamopen, _ := jsonparser.GetString(value, "pasaran_jamopen")
		pasaran_statusactive, _ := jsonparser.GetString(value, "pasaran_statusactive")
		royaltyfee, _ := jsonparser.GetFloat(value, "royaltyfee")
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
		maxbet4d_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet4d_fullbb_432d")
		maxbet3d_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet3d_fullbb_432d")
		maxbet3dd_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet3dd_fullbb_432d")
		maxbet2d_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2d_fullbb_432d")
		maxbet2dd_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2dd_fullbb_432d")
		maxbet2dt_fullbb_432d, _ := jsonparser.GetFloat(value, "maxbet2dt_fullbb_432d")
		maxbuy4d_432d, _ := jsonparser.GetFloat(value, "maxbuy4d_432d")
		maxbuy3d_432d, _ := jsonparser.GetFloat(value, "maxbuy3d_432d")
		maxbuy3dd_432d, _ := jsonparser.GetFloat(value, "maxbuy3dd_432d")
		maxbuy2d_432d, _ := jsonparser.GetFloat(value, "maxbuy2d_432d")
		maxbuy2dd_432d, _ := jsonparser.GetFloat(value, "maxbuy2dd_432d")
		maxbuy2dt_432d, _ := jsonparser.GetFloat(value, "maxbuy2dt_432d")
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
		keigenap_5050umum, _ := jsonparser.GetFloat(value, "keigenap_5050umum")
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

		obj.Company_Pasaran_diundi = pasaran_diundi
		obj.Company_Pasaran_url = pasaran_url
		obj.Company_Pasaran_jamtutup = pasaran_jamtutup
		obj.Company_Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Company_Pasaran_jamopen = pasaran_jamopen
		obj.Company_Pasaran_statusactive = pasaran_statusactive
		obj.Company_Royaltyfee = float32(royaltyfee)
		obj.Company_Limitline4d = int(limitline_4d)
		obj.Company_Limitline3d = int(limitline_3d)
		obj.Company_Limitline3dd = int(limitline_3dd)
		obj.Company_Limitline2d = int(limitline_2d)
		obj.Company_Limitline2dd = int(limitline_2dd)
		obj.Company_Limitline2dt = int(limitline_2dt)
		obj.Company_Bbfs = int(bbfs)
		obj.Company_Minbet_432d = float32(minbet_432d)
		obj.Company_Maxbet4d_432d = float32(maxbet4d_432d)
		obj.Company_Maxbet3d_432d = float32(maxbet3d_432d)
		obj.Company_Maxbet3dd_432d = float32(maxbet3dd_432d)
		obj.Company_Maxbet2d_432d = float32(maxbet2d_432d)
		obj.Company_Maxbet2dd_432d = float32(maxbet2dd_432d)
		obj.Company_Maxbet2dt_432d = float32(maxbet2dt_432d)
		obj.Company_Maxbet4d_fullbb_432d = float32(maxbet4d_fullbb_432d)
		obj.Company_Maxbet3d_fullbb_432d = float32(maxbet3d_fullbb_432d)
		obj.Company_Maxbet3dd_fullbb_432d = float32(maxbet3dd_fullbb_432d)
		obj.Company_Maxbet2d_fullbb_432d = float32(maxbet2d_fullbb_432d)
		obj.Company_Maxbet2dd_fullbb_432d = float32(maxbet2dd_fullbb_432d)
		obj.Company_Maxbet2dt_fullbb_432d = float32(maxbet2dt_fullbb_432d)
		obj.Company_Maxbuy4d_432d = float32(maxbuy4d_432d)
		obj.Company_Maxbuy3d_432d = float32(maxbuy3d_432d)
		obj.Company_Maxbuy3dd_432d = float32(maxbuy3dd_432d)
		obj.Company_Maxbuy2d_432d = float32(maxbuy2d_432d)
		obj.Company_Maxbuy2dd_432d = float32(maxbuy2dd_432d)
		obj.Company_Maxbuy2dt_432d = float32(maxbuy2dt_432d)
		obj.Company_Limitotal4d_432d = float32(limitotal4d_432d)
		obj.Company_Limitotal3d_432d = float32(limitotal3d_432d)
		obj.Company_Limitotal3dd_432d = float32(limitotal3dd_432d)
		obj.Company_Limitotal2d_432d = float32(limitotal2d_432d)
		obj.Company_Limitotal2dd_432d = float32(limitotal2dd_432d)
		obj.Company_Limitotal2dt_432d = float32(limitotal2dt_432d)
		obj.Company_Limitglobal4d_432d = float32(limitglobal4d_432d)
		obj.Company_Limitglobal3d_432d = float32(limitglobal3d_432d)
		obj.Company_Limitglobal3dd_432d = float32(limitglobal3dd_432d)
		obj.Company_Limitglobal2d_432d = float32(limitglobal2d_432d)
		obj.Company_Limitglobal2dd_432d = float32(limitglobal2dd_432d)
		obj.Company_Limitglobal2dt_432d = float32(limitglobal2dt_432d)
		obj.Company_Limitotal4d_fullbb_432d = float32(limitotal4d_fullbb_432d)
		obj.Company_Limitotal3d_fullbb_432d = float32(limitotal3d_fullbb_432d)
		obj.Company_Limitotal3dd_fullbb_432d = float32(limitotal3dd_fullbb_432d)
		obj.Company_Limitotal2d_fullbb_432d = float32(limitotal2d_fullbb_432d)
		obj.Company_Limitotal2dd_fullbb_432d = float32(limitotal2dd_fullbb_432d)
		obj.Company_Limitotal2dt_fullbb_432d = float32(limitotal2dt_fullbb_432d)
		obj.Company_Limitglobal4d_fullbb_432d = float32(limitglobal4d_fullbb_432d)
		obj.Company_Limitglobal3d_fullbb_432d = float32(limitglobal3d_fullbb_432d)
		obj.Company_Limitglobal3dd_fullbb_432d = float32(limitglobal3dd_fullbb_432d)
		obj.Company_Limitglobal2d_fullbb_432d = float32(limitglobal2d_fullbb_432d)
		obj.Company_Limitglobal2dd_fullbb_432d = float32(limitglobal2dd_fullbb_432d)
		obj.Company_Limitglobal2dt_fullbb_432d = float32(limitglobal2dt_fullbb_432d)
		obj.Company_Disc4d_432d = float32(disc4d_432d)
		obj.Company_Disc3d_432d = float32(disc3d_432d)
		obj.Company_Disc3dd_432d = float32(disc3dd_432d)
		obj.Company_Disc2d_432d = float32(disc2d_432d)
		obj.Company_Disc2dd_432d = float32(disc2dd_432d)
		obj.Company_Disc2dt_432d = float32(disc2dt_432d)
		obj.Company_Win4d_432d = float32(win4d_432d)
		obj.Company_Win3d_432d = float32(win3d_432d)
		obj.Company_Win3dd_432d = float32(win3dd_432d)
		obj.Company_Win2d_432d = float32(win2d_432d)
		obj.Company_Win2dd_432d = float32(win2dd_432d)
		obj.Company_Win2dt_432d = float32(win2dt_432d)
		obj.Company_Win4dnodisc_432d = float32(win4dnodisc_432d)
		obj.Company_Win3dnodisc_432d = float32(win3dnodisc_432d)
		obj.Company_Win3ddnodisc_432d = float32(win3ddnodisc_432d)
		obj.Company_Win2dnodisc_432d = float32(win2dnodisc_432d)
		obj.Company_Win2ddnodisc_432d = float32(win2ddnodisc_432d)
		obj.Company_Win2dtnodisc_432d = float32(win2dtnodisc_432d)
		obj.Company_Win4dbb_kena_432d = float32(win4dbb_kena_432d)
		obj.Company_Win3dbb_kena_432d = float32(win3dbb_kena_432d)
		obj.Company_Win3ddbb_kena_432d = float32(win3ddbb_kena_432d)
		obj.Company_Win2dbb_kena_432d = float32(win2dbb_kena_432d)
		obj.Company_Win2ddbb_kena_432d = float32(win2ddbb_kena_432d)
		obj.Company_Win2dtbb_kena_432d = float32(win2dtbb_kena_432d)
		obj.Company_Win4dbb_432d = float32(win4dbb_432d)
		obj.Company_Win3dbb_432d = float32(win3dbb_432d)
		obj.Company_Win3ddbb_432d = float32(win3ddbb_432d)
		obj.Company_Win2dbb_432d = float32(win2dbb_432d)
		obj.Company_Win2ddbb_432d = float32(win2ddbb_432d)
		obj.Company_Win2dtbb_432d = float32(win2dtbb_432d)
		obj.Company_Minbet_cbebas = float32(minbet_cbebas)
		obj.Company_Maxbet_cbebas = float32(maxbet_cbebas)
		obj.Company_Maxbuy_cbebas = float32(maxbuy_cbebas)
		obj.Company_Win_cbebas = float32(win_cbebas)
		obj.Company_Disc_cbebas = float32(disc_cbebas)
		obj.Company_Limitglobal_cbebas = float32(limitglobal_cbebas)
		obj.Company_Limittotal_cbebas = float32(limittotal_cbebas)
		obj.Company_Minbet_cmacau = float32(minbet_cmacau)
		obj.Company_Maxbet_cmacau = float32(maxbet_cmacau)
		obj.Company_Maxbuy_cmacau = float32(maxbuy_cmacau)
		obj.Company_Win2d_cmacau = float32(win2d_cmacau)
		obj.Company_Win3d_cmacau = float32(win3d_cmacau)
		obj.Company_Win4d_cmacau = float32(win4d_cmacau)
		obj.Company_Disc_cmacau = float32(disc_cmacau)
		obj.Company_Limitglobal_cmacau = float32(limitglobal_cmacau)
		obj.Company_Limitotal_cmacau = float32(limitotal_cmacau)
		obj.Company_Minbet_cnaga = float32(minbet_cnaga)
		obj.Company_Maxbet_cnaga = float32(maxbet_cnaga)
		obj.Company_Maxbuy_cnaga = float32(maxbuy_cnaga)
		obj.Company_Win3_cnaga = float32(win3_cnaga)
		obj.Company_Win4_cnaga = float32(win4_cnaga)
		obj.Company_Disc_cnaga = float32(disc_cnaga)
		obj.Company_Limitglobal_cnaga = float32(limitglobal_cnaga)
		obj.Company_Limittotal_cnaga = float32(limittotal_cnaga)
		obj.Company_Minbet_cjitu = float32(minbet_cjitu)
		obj.Company_Maxbet_cjitu = float32(maxbet_cjitu)
		obj.Company_Maxbuy_cjitu = float32(maxbuy_cjitu)
		obj.Company_Winas_cjitu = float32(winas_cjitu)
		obj.Company_Winkop_cjitu = float32(winkop_cjitu)
		obj.Company_Winkepala_cjitu = float32(winkepala_cjitu)
		obj.Company_Winekor_cjitu = float32(winekor_cjitu)
		obj.Company_Desc_cjitu = float32(desc_cjitu)
		obj.Company_Limitglobal_cjitu = float32(limitglobal_cjitu)
		obj.Company_Limittotal_cjitu = float32(limittotal_cjitu)
		obj.Company_Minbet_5050umum = float32(minbet_5050umum)
		obj.Company_Maxbet_5050umum = float32(maxbet_5050umum)
		obj.Company_Maxbuy_5050umum = float32(maxbuy_5050umum)
		obj.Company_Keibesar_5050umum = float32(keibesar_5050umum)
		obj.Company_Keikecil_5050umum = float32(keikecil_5050umum)
		obj.Company_Keigenap_5050umum = float32(keigenap_5050umum)
		obj.Company_Keiganjil_5050umum = float32(keiganjil_5050umum)
		obj.Company_Keitengah_5050umum = float32(keitengah_5050umum)
		obj.Company_Keitepi_5050umum = float32(keitepi_5050umum)
		obj.Company_Discbesar_5050umum = float32(discbesar_5050umum)
		obj.Company_Disckecil_5050umum = float32(disckecil_5050umum)
		obj.Company_Discgenap_5050umum = float32(discgenap_5050umum)
		obj.Company_Discganjil_5050umum = float32(discganjil_5050umum)
		obj.Company_Disctengah_5050umum = float32(disctengah_5050umum)
		obj.Company_Disctepi_5050umum = float32(disctepi_5050umum)
		obj.Company_Limitglobal_5050umum = float32(limitglobal_5050umum)
		obj.Company_Limittotal_5050umum = float32(limittotal_5050umum)
		obj.Company_Minbet_5050special = float32(minbet_5050special)
		obj.Company_Maxbet_5050special = float32(maxbet_5050special)
		obj.Company_Maxbuy_5050special = float32(maxbuy_5050special)
		obj.Company_Keiasganjil_5050special = float32(keiasganjil_5050special)
		obj.Company_Keiasgenap_5050special = float32(keiasgenap_5050special)
		obj.Company_Keiasbesar_5050special = float32(keiasbesar_5050special)
		obj.Company_Keiaskecil_5050special = float32(keiaskecil_5050special)
		obj.Company_Keikopganjil_5050special = float32(keikopganjil_5050special)
		obj.Company_Keikopgenap_5050special = float32(keikopgenap_5050special)
		obj.Company_Keikopbesar_5050special = float32(keikopbesar_5050special)
		obj.Company_Keikopkecil_5050special = float32(keikopkecil_5050special)
		obj.Company_Keikepalaganjil_5050special = float32(keikepalaganjil_5050special)
		obj.Company_Keikepalagenap_5050special = float32(keikepalagenap_5050special)
		obj.Company_Keikepalabesar_5050special = float32(keikepalabesar_5050special)
		obj.Company_Keikepalakecil_5050special = float32(keikepalakecil_5050special)
		obj.Company_Keiekorganjil_5050special = float32(keiekorganjil_5050special)
		obj.Company_Keiekorgenap_5050special = float32(keiekorgenap_5050special)
		obj.Company_Keiekorbesar_5050special = float32(keiekorbesar_5050special)
		obj.Company_Keiekorkecil_5050special = float32(keiekorkecil_5050special)
		obj.Company_Discasganjil_5050special = float32(discasganjil_5050special)
		obj.Company_Discasgenap_5050special = float32(discasgenap_5050special)
		obj.Company_Discasbesar_5050special = float32(discasbesar_5050special)
		obj.Company_Discaskecil_5050special = float32(discaskecil_5050special)
		obj.Company_Disckopganjil_5050special = float32(disckopganjil_5050special)
		obj.Company_Disckopgenap_5050special = float32(disckopgenap_5050special)
		obj.Company_Disckopbesar_5050special = float32(disckopbesar_5050special)
		obj.Company_Disckopkecil_5050special = float32(disckopkecil_5050special)
		obj.Company_Disckepalaganjil_5050special = float32(disckepalaganjil_5050special)
		obj.Company_Disckepalagenap_5050special = float32(disckepalagenap_5050special)
		obj.Company_Disckepalabesar_5050special = float32(disckepalabesar_5050special)
		obj.Company_Disckepalakecil_5050special = float32(disckepalakecil_5050special)
		obj.Company_Discekorganjil_5050special = float32(discekorganjil_5050special)
		obj.Company_Discekorgenap_5050special = float32(discekorgenap_5050special)
		obj.Company_Discekorbesar_5050special = float32(discekorbesar_5050special)
		obj.Company_Discekorkecil_5050special = float32(discekorkecil_5050special)
		obj.Company_Limitglobal_5050special = float32(limitglobal_5050special)
		obj.Company_Limittotal_5050special = float32(limittotal_5050special)
		obj.Company_Minbet_5050kombinasi = float32(minbet_5050kombinasi)
		obj.Company_Maxbet_5050kombinasi = float32(maxbet_5050kombinasi)
		obj.Company_Maxbuy_5050kombinasi = float32(maxbuy_5050kombinasi)
		obj.Company_Belakangkeimono_5050kombinasi = float32(belakangkeimono_5050kombinasi)
		obj.Company_Belakangkeistereo_5050kombinasi = float32(belakangkeistereo_5050kombinasi)
		obj.Company_Belakangkeikembang_5050kombinasi = float32(belakangkeikembang_5050kombinasi)
		obj.Company_Belakangkeikempis_5050kombinasi = float32(belakangkeikempis_5050kombinasi)
		obj.Company_Belakangkeikembar_5050kombinasi = float32(belakangkeikembar_5050kombinasi)
		obj.Company_Tengahkeimono_5050kombinasi = float32(tengahkeimono_5050kombinasi)
		obj.Company_Tengahkeistereo_5050kombinasi = float32(tengahkeistereo_5050kombinasi)
		obj.Company_Tengahkeikembang_5050kombinasi = float32(tengahkeikembang_5050kombinasi)
		obj.Company_Tengahkeikempis_5050kombinasi = float32(tengahkeikempis_5050kombinasi)
		obj.Company_Tengahkeikembar_5050kombinasi = float32(tengahkeikembar_5050kombinasi)
		obj.Company_Depankeimono_5050kombinasi = float32(depankeimono_5050kombinasi)
		obj.Company_Depankeistereo_5050kombinasi = float32(depankeistereo_5050kombinasi)
		obj.Company_Depankeikembang_5050kombinasi = float32(depankeikembang_5050kombinasi)
		obj.Company_Depankeikempis_5050kombinasi = float32(depankeikempis_5050kombinasi)
		obj.Company_Depankeikembar_5050kombinasi = float32(depankeikembar_5050kombinasi)
		obj.Company_Belakangdiscmono_5050kombinasi = float32(belakangdiscmono_5050kombinasi)
		obj.Company_Belakangdiscstereo_5050kombinasi = float32(belakangdiscstereo_5050kombinasi)
		obj.Company_Belakangdisckembang_5050kombinasi = float32(belakangdisckembang_5050kombinasi)
		obj.Company_Belakangdisckempis_5050kombinasi = float32(belakangdisckempis_5050kombinasi)
		obj.Company_Belakangdisckembar_5050kombinasi = float32(belakangdisckembar_5050kombinasi)
		obj.Company_Tengahdiscmono_5050kombinasi = float32(tengahdiscmono_5050kombinasi)
		obj.Company_Tengahdiscstereo_5050kombinasi = float32(tengahdiscstereo_5050kombinasi)
		obj.Company_Tengahdisckembang_5050kombinasi = float32(tengahdisckembang_5050kombinasi)
		obj.Company_Tengahdisckempis_5050kombinasi = float32(tengahdisckempis_5050kombinasi)
		obj.Company_Tengahdisckembar_5050kombinasi = float32(tengahdisckembar_5050kombinasi)
		obj.Company_Depandiscmono_5050kombinasi = float32(depandiscmono_5050kombinasi)
		obj.Company_Depandiscstereo_5050kombinasi = float32(depandiscstereo_5050kombinasi)
		obj.Company_Depandisckembang_5050kombinasi = float32(depandisckembang_5050kombinasi)
		obj.Company_Depandisckempis_5050kombinasi = float32(depandisckempis_5050kombinasi)
		obj.Company_Depandisckembar_5050kombinasi = float32(depandisckembar_5050kombinasi)
		obj.Company_Limitglobal_5050kombinasi = float32(limitglobal_5050kombinasi)
		obj.Company_Limittotal_5050kombinasi = float32(limittotal_5050kombinasi)
		obj.Company_Minbet_kombinasi = float32(minbet_kombinasi)
		obj.Company_Maxbet_kombinasi = float32(maxbet_kombinasi)
		obj.Company_Maxbuy_kombinasi = float32(maxbuy_kombinasi)
		obj.Company_Win_kombinasi = float32(win_kombinasi)
		obj.Company_Disc_kombinasi = float32(disc_kombinasi)
		obj.Company_Limitglobal_kombinasi = float32(limitglobal_kombinasi)
		obj.Company_Limittotal_kombinasi = float32(limittotal_kombinasi)
		obj.Company_Minbet_dasar = float32(minbet_dasar)
		obj.Company_Maxbet_dasar = float32(maxbet_dasar)
		obj.Company_Maxbuy_dasar = float32(maxbuy_dasar)
		obj.Company_Keibesar_dasar = float32(keibesar_dasar)
		obj.Company_Keikecil_dasar = float32(keikecil_dasar)
		obj.Company_Keigenap_dasar = float32(keigenap_dasar)
		obj.Company_Keiganjil_dasar = float32(keiganjil_dasar)
		obj.Company_Discbesar_dasar = float32(discbesar_dasar)
		obj.Company_Disckecil_dasar = float32(disckecil_dasar)
		obj.Company_Discgenap_dasar = float32(discgenap_dasar)
		obj.Company_Discganjil_dasar = float32(discganjil_dasar)
		obj.Company_Limitglobal_dasar = float32(limitglobal_dasar)
		obj.Company_Limittotal_dasar = float32(limittotal_dasar)
		obj.Company_Minbet_shio = float32(minbet_shio)
		obj.Company_Maxbet_shio = float32(maxbet_shio)
		obj.Company_Maxbuy_shio = float32(maxbuy_shio)
		obj.Company_Win_shio = float32(win_shio)
		obj.Company_Disc_shio = float32(disc_shio)
		obj.Company_Shioyear_shio = shioyear_shio
		obj.Company_Limitglobal_shio = float32(limitglobal_shio)
		obj.Company_Limittotal_shio = float32(limittotal_shio)

		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_listpasaranConf(client.Company, client.Companypasaran_id)
		helpers.SetRedis(Fieldcompanylistpasaranconf_home_redis+"_"+client.Company+"_"+strconv.Itoa(client.Companypasaran_id), result, 20*time.Minute)
		log.Println("COMPANY CONF PASARAN MYSQL " + client.Company)
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
		log.Println("COMPANY CONF PASARAN CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyListKeluaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companylistkeluaran)
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
	var obj entities.Model_companylistkeluaran
	var arraobj []entities.Model_companylistkeluaran
	resultredis, flag := helpers.GetRedis(Fieldcompanylistpasarankeluaran_home_redis + "_" + client.Company + "_" + client.Periode + "_" + client.Year + "_" + strconv.Itoa(client.Pasaran))
	jsonredis := []byte(resultredis)
	totalwinlose_RD, _ := jsonparser.GetInt(jsonredis, "totalwinlose")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_pasaran_no, _ := jsonparser.GetInt(value, "company_pasaran_no")
		company_pasaran_invoice, _ := jsonparser.GetInt(value, "company_pasaran_invoice")
		company_pasaran_idcompp, _ := jsonparser.GetInt(value, "company_pasaran_idcompp")
		company_pasaran_code, _ := jsonparser.GetString(value, "company_pasaran_code")
		company_pasaran_periode, _ := jsonparser.GetString(value, "company_pasaran_periode")
		company_pasaran_name, _ := jsonparser.GetString(value, "company_pasaran_name")
		company_pasaran_tanggal, _ := jsonparser.GetString(value, "company_pasaran_tanggal")
		company_pasaran_keluaran, _ := jsonparser.GetString(value, "company_pasaran_keluaran")
		company_pasaran_status, _ := jsonparser.GetString(value, "company_pasaran_status")
		company_pasaran_status_css, _ := jsonparser.GetString(value, "company_pasaran_status_css")
		company_pasaran_totalmember, _ := jsonparser.GetFloat(value, "company_pasaran_totalmember")
		company_pasaran_totalbet, _ := jsonparser.GetFloat(value, "company_pasaran_totalbet")
		company_pasaran_totaloutstanding, _ := jsonparser.GetFloat(value, "company_pasaran_totaloutstanding")
		company_pasaran_totalcancelbet, _ := jsonparser.GetFloat(value, "company_pasaran_totalcancelbet")
		company_pasaran_winlose, _ := jsonparser.GetFloat(value, "company_pasaran_winlose")
		company_pasaran_winlosetemp, _ := jsonparser.GetInt(value, "company_pasaran_winlosetemp")
		company_pasaran_revisi, _ := jsonparser.GetInt(value, "company_pasaran_revisi")
		company_pasaran_msgrevisi, _ := jsonparser.GetString(value, "company_pasaran_msgrevisi")

		obj.Company_Pasaran_no = int(company_pasaran_no)
		obj.Company_Pasaran_idtrxkeluaran = int(company_pasaran_invoice)
		obj.Company_Pasaran_idcomppasaran = int(company_pasaran_idcompp)
		obj.Company_Pasaran_pasarancode = company_pasaran_code
		obj.Company_Pasaran_keluaranperiode = company_pasaran_periode
		obj.Company_Pasaran_nmpasaran = company_pasaran_name
		obj.Company_Pasaran_tanggalperiode = company_pasaran_tanggal
		obj.Company_Pasaran_keluarantogel = company_pasaran_keluaran
		obj.Company_Pasaran_status = company_pasaran_status
		obj.Company_Pasaran_status_css = company_pasaran_status_css
		obj.Company_Pasaran_total_Member = float32(company_pasaran_totalmember)
		obj.Company_Pasaran_total_bet = float32(company_pasaran_totalbet)
		obj.Company_Pasaran_total_outstanding = float32(company_pasaran_totaloutstanding)
		obj.Company_Pasaran_total_cancelbet = float32(company_pasaran_totalcancelbet)
		obj.Company_Pasaran_winlose = float32(company_pasaran_winlose)
		obj.Company_Pasaran_winlosetemp = int(company_pasaran_winlosetemp)
		obj.Company_Pasaran_revisi = int(company_pasaran_revisi)
		obj.Company_Pasaran_msgrevisi = company_pasaran_msgrevisi
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_listkeluaran(client.Company, client.Periode, client.Year, client.Pasaran)
		helpers.SetRedis(Fieldcompanylistpasarankeluaran_home_redis+"_"+client.Company+"_"+client.Periode+"_"+client.Year+"_"+strconv.Itoa(client.Pasaran), result, 20*time.Minute)
		log.Println("COMPANY LISTPASARAN KELUARAN MYSQL " + client.Company)
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
		log.Println("COMPANY LISTPASARAN KELUARAN CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"message":      "Success",
			"record":       arraobj,
			"totalwinlose": totalwinlose_RD,
			"time":         time.Since(render_page).String(),
		})
	}
}
func CompanyInvoiceMember(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyinvoice)
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
	var obj entities.Model_invoicelistMember
	var arraobj []entities.Model_invoicelistMember
	resultredis, flag := helpers.GetRedis(Fieldcompanyinvoicelist_home_redis + "_" + client.Company + "_" + strconv.Itoa(client.Invoice))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		member, _ := jsonparser.GetString(value, "member")
		totalbet, _ := jsonparser.GetInt(value, "totalbet")
		totalbayar, _ := jsonparser.GetInt(value, "totalbayar")
		totalcancelbet, _ := jsonparser.GetInt(value, "totalcancelbet")
		totalwin, _ := jsonparser.GetInt(value, "totalwin")

		obj.Member = member
		obj.Totalbet = int(totalbet)
		obj.Totalbayar = int(totalbayar)
		obj.Totalcancelbet = int(totalcancelbet)
		obj.Totalwin = int(totalwin)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_invoice_member(client.Company, client.Invoice)
		helpers.SetRedis(Fieldcompanyinvoicelist_home_redis+"_"+client.Company+"_"+strconv.Itoa(client.Invoice), result, 20*time.Minute)
		log.Println("COMPANY INVOICE LIST MEMBER MYSQL " + client.Company)
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
		log.Println("COMPANY INVOICE LIST MEMBER CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyInvoiceMemberTemp(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyinvoice)
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
	var obj entities.Model_invoicelistMember
	var arraobj []entities.Model_invoicelistMember
	resultredis, flag := helpers.GetRedis(Fieldcompanyinvoicelisttemp_home_redis + "_" + client.Company + "_" + strconv.Itoa(client.Invoice))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		member, _ := jsonparser.GetString(value, "member")
		totalbet, _ := jsonparser.GetInt(value, "totalbet")
		totalbayar, _ := jsonparser.GetInt(value, "totalbayar")
		totalcancelbet, _ := jsonparser.GetInt(value, "totalcancelbet")
		totalwin, _ := jsonparser.GetInt(value, "totalwin")

		obj.Member = member
		obj.Totalbet = int(totalbet)
		obj.Totalbayar = int(totalbayar)
		obj.Totalcancelbet = int(totalcancelbet)
		obj.Totalwin = int(totalwin)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_invoice_membertemp(client.Company, client.Invoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyinvoicelisttemp_home_redis+"_"+client.Company+"_"+strconv.Itoa(client.Invoice), result, 20*time.Minute)
		log.Println("COMPANY INVOICE LIST MEMBER TEMP MYSQL " + client.Company)
		return c.JSON(result)
	} else {
		log.Println("COMPANY INVOICE LIST MEMBER TEMP CACHE " + client.Company)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyInvoiceMemberSync(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyinvoice)
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
	_deleteredis_company(client.Company, client.Invoice, client.Pasaran, client.Month, client.Year)
	return c.JSON(result)
}
func CompanyInvoiceGroupPermainan(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyinvoice)
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
	client := new(entities.Controller_companyinvoicelistpermainan)
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
	var obj entities.Model_invoicelistpermainan
	var arraobj []entities.Model_invoicelistpermainan
	resultredis, flag := helpers.GetRedis(Fieldcompanyinvoicelistpermainan_home_redis + "_" + client.Company + "_" + strconv.Itoa(client.Invoice) + "_" + client.Permainan)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		bet_id, _ := jsonparser.GetInt(value, "bet_id")
		bet_datetime, _ := jsonparser.GetString(value, "bet_datetime")
		bet_ipaddress, _ := jsonparser.GetString(value, "bet_ipaddress")
		bet_device, _ := jsonparser.GetString(value, "bet_device")
		bet_timezone, _ := jsonparser.GetString(value, "bet_timezone")
		bet_username, _ := jsonparser.GetString(value, "bet_username")
		bet_typegame, _ := jsonparser.GetString(value, "bet_typegame")
		bet_posisitogel, _ := jsonparser.GetString(value, "bet_posisitogel")
		bet_nomortogel, _ := jsonparser.GetString(value, "bet_nomortogel")
		bet_bet, _ := jsonparser.GetInt(value, "bet_bet")
		bet_diskon, _ := jsonparser.GetInt(value, "bet_diskon")
		bet_diskonpercen, _ := jsonparser.GetInt(value, "bet_diskonpercen")
		bet_kei, _ := jsonparser.GetInt(value, "bet_kei")
		bet_keipercen, _ := jsonparser.GetInt(value, "bet_keipercen")
		bet_win, _ := jsonparser.GetFloat(value, "bet_win")
		bet_totalwin, _ := jsonparser.GetInt(value, "bet_totalwin")
		bet_bayar, _ := jsonparser.GetInt(value, "bet_bayar")
		bet_status, _ := jsonparser.GetString(value, "bet_status")
		bet_statuscss, _ := jsonparser.GetString(value, "bet_statuscss")
		bet_create, _ := jsonparser.GetString(value, "bet_create")
		bet_createdate, _ := jsonparser.GetString(value, "bet_createdate")
		bet_update, _ := jsonparser.GetString(value, "bet_update")
		bet_updatedate, _ := jsonparser.GetString(value, "bet_updatedate")

		obj.Bet_id = int(bet_id)
		obj.Bet_datetime = bet_datetime
		obj.Bet_ipaddress = bet_ipaddress
		obj.Bet_device = bet_device
		obj.Bet_timezone = bet_timezone
		obj.Bet_username = bet_username
		obj.Bet_typegame = bet_typegame
		obj.Bet_nomortogel = bet_nomortogel
		obj.Bet_posisitogel = bet_posisitogel
		obj.Bet_bet = int(bet_bet)
		obj.Bet_diskon = int(bet_diskon)
		obj.Bet_diskonpercen = int(bet_diskonpercen)
		obj.Bet_kei = int(bet_kei)
		obj.Bet_keipercen = int(bet_keipercen)
		obj.Bet_win = float32(bet_win)
		obj.Bet_totalwin = int(bet_totalwin)
		obj.Bet_bayar = int(bet_bayar)
		obj.Bet_status = bet_status
		obj.Bet_statuscss = bet_statuscss
		obj.Bet_create = bet_create
		obj.Bet_createDate = bet_createdate
		obj.Bet_update = bet_update
		obj.Bet_updateDate = bet_updatedate
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_company_invoice_listpermainan(client.Company, client.Permainan, client.Invoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyinvoicelistpermainan_home_redis+"_"+client.Company+"_"+strconv.Itoa(client.Invoice)+"_"+client.Permainan, result, 20*time.Minute)
		log.Println("LIST PERMAINAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LIST PERMAINAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CompanyInvoicelistpermainanbystatus(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyinvoicelistpermainanstatus)
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
	client := new(entities.Controller_companyinvoicelistpermainanusername)
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
	client := new(entities.Controller_companysave)
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
	_deleteredis_company(client.Company, 0, "", "", "")
	return c.JSON(result)
}
func CompanySaveNewAdmin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminsave)
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
	_deleteredis_company(client.Company, 0, "", "", "")
	return c.JSON(result)
}
func CompanySaveNewPasaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaransave)
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
	_deleteredis_company(client.Company, 0, "", "", "")

	return c.JSON(result)
}
func CompanySaveNewPasaranHariOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companydetailonlinesave)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, "", "", "")
	return c.JSON(result)
}
func CompanyDeletePasaranHariOnline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companydeletepasaranonline)
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
	log.Printf("%s-%s-%d-%d", client.Company, client.Master, client.Companypasaran_id, client.Companypasaran_idoff)
	result, err := models.Delete_companyPasaranharionline(client.Master, client.Company, client.Companypasaran_id, client.Companypasaran_idoff)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_company(client.Company, client.Companypasaran_id, "", "", "")
	return c.JSON(result)
}
func CompanyFetchPasaranlimitline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaran432(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasarancolokbebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasarancolokmacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasarancoloknaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasarancolokjitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaran5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaran5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaran5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaranmacaukombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasarandasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyFetchPasaranshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyfetchpasaran432)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdate(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaran)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, "", "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdateRoyaltyFee(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaranroyaltyfee)
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

	result, err := models.Save_companyUpdatePasaranRoyaltyFee(
		client.Master, client.Company, client.Royaltyfee, client.Companypasaran_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_company(client.Company, client.Companypasaran_id, "", "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdateLimitline(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaranline)
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
		client.Pasaran_limitline_3d, client.Pasaran_limitline_3dd, client.Pasaran_limitline_2d, client.Pasaran_limitline_2dd,
		client.Pasaran_limitline_2dt, client.Pasaran_bbfs)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdate432(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaran432)
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
		client.Pasaran_minbet_432d, client.Pasaran_maxbet4d_432d, client.Pasaran_maxbet3d_432d, client.Pasaran_maxbet3dd_432d,
		client.Pasaran_maxbet2d_432d, client.Pasaran_maxbet2dd_432d, client.Pasaran_maxbet2dt_432d,
		client.Pasaran_maxbet4d_fullbb_432d, client.Pasaran_maxbet3d_fullbb_432d, client.Pasaran_maxbet3dd_fullbb_432d,
		client.Pasaran_maxbet2d_fullbb_432d, client.Pasaran_maxbet2dd_fullbb_432d, client.Pasaran_maxbet2dt_fullbb_432d,
		client.Pasaran_maxbuy4d_432d, client.Pasaran_maxbuy3d_432d, client.Pasaran_maxbuy3dd_432d,
		client.Pasaran_maxbuy2d_432d, client.Pasaran_maxbuy2dd_432d, client.Pasaran_maxbuy2dt_432d,
		client.Pasaran_win4d_432d, client.Pasaran_win3d_432d, client.Pasaran_win3dd_432d, client.Pasaran_win2d_432d, client.Pasaran_win2dd_432d, client.Pasaran_win2dt_432d,
		client.Pasaran_win4dnodisc_432d, client.Pasaran_win3dnodisc_432d, client.Pasaran_win3ddnodisc_432d, client.Pasaran_win2dnodisc_432d, client.Pasaran_win2ddnodisc_432d, client.Pasaran_win2dtnodisc_432d,
		client.Pasaran_win4dbb_kena_432d, client.Pasaran_win3dbb_kena_432d, client.Pasaran_win3ddbb_kena_432d, client.Pasaran_win2dbb_kena_432d, client.Pasaran_win2ddbb_kena_432d, client.Pasaran_win2dtbb_kena_432d,
		client.Pasaran_win4dbb_432d, client.Pasaran_win3dbb_432d, client.Pasaran_win3ddbb_432d, client.Pasaran_win2dbb_432d, client.Pasaran_win2ddbb_432d, client.Pasaran_win2dtbb_432d,
		client.Pasaran_disc4d_432d, client.Pasaran_disc3d_432d, client.Pasaran_disc3dd_432d, client.Pasaran_disc2d_432d, client.Pasaran_disc2dd_432d, client.Pasaran_disc2dt_432d,
		client.Pasaran_limitglobal4d_432d, client.Pasaran_limitglobal3d_432d, client.Pasaran_limitglobal3dd_432d, client.Pasaran_limitglobal2d_432d, client.Pasaran_limitglobal2dd_432d, client.Pasaran_limitglobal2dt_432d,
		client.Pasaran_limitotal4d_432d, client.Pasaran_limitotal3d_432d, client.Pasaran_limitotal3dd_432d, client.Pasaran_limitotal2d_432d, client.Pasaran_limitotal2dd_432d, client.Pasaran_limitotal2dt_432d,
		client.Pasaran_limitglobal4d_fullbb_432d, client.Pasaran_limitglobal3d_fullbb_432d, client.Pasaran_limitglobal3dd_fullbb_432d, client.Pasaran_limitglobal2d_fullbb_432d, client.Pasaran_limitglobal2dd_fullbb_432d, client.Pasaran_limitglobal2dt_432d,
		client.Pasaran_limitotal4d_fullbb_432d, client.Pasaran_limitotal3d_fullbb_432d, client.Pasaran_limitotal3dd_fullbb_432d, client.Pasaran_limitotal2d_fullbb_432d, client.Pasaran_limitotal2dd_fullbb_432d, client.Pasaran_limitotal2dt_fullbb_432d)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokbebas(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarancolokbebas)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokmacau(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarancolokmacau)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatecoloknaga(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarancoloknaga)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatecolokjitu(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarancolokjitu)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdate5050umum(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaran5050umum)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdate5050special(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaran5050special)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdate5050kombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaran5050kombinasi)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatekombinasi(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarankombinasi)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdatedasar(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasarandasar)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func CompanyPasaranUpdateshio(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companypasaranshio)
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
	_deleteredis_company(client.Company, client.Companypasaran_id, client.Pasaran_id, "", "")
	return c.JSON(result)
}
func _deleteredis_company(company string, idcomppasaran int, pasaranid, periode, year string) {
	//MASTER
	val_company := helpers.DeleteRedis(Fieldcompany_home_redis)
	log.Printf("Redis Delete MASTER COMPANY : %d", val_company)
	val_company_listadmin := helpers.DeleteRedis(Fieldcompanylistadmin_home_redis + "_" + company)
	log.Printf("Redis Delete MASTER COMPANY LISTADMIN : %d", val_company_listadmin)
	val_master := helpers.DeleteRedis(Fieldcompanylistpasaran_home_redis + "_" + company)
	log.Printf("REDIS DELETE MASTER CONFIG : %d", val_master)

	val_master2 := helpers.DeleteRedis(Fieldcompanylistpasarankeluaran_home_redis + "_" + company + "_" + periode + "_" + year + "_" + pasaranid)
	log.Printf("REDIS DELETE MASTER2 CONFIG : %d", val_master2)

	if idcomppasaran > 0 {
		val_master_pasaranonline := helpers.DeleteRedis(Fieldcompanylistpasaranonline_home_redis + "_" + company + "_" + strconv.Itoa(idcomppasaran))
		log.Printf("REDIS DELETE MASTER: %d", val_master_pasaranonline)
		val_master := helpers.DeleteRedis(Fieldcompanylistpasaranconf_home_redis + "_" + company + "_" + strconv.Itoa(idcomppasaran))
		log.Printf("REDIS DELETE MASTER CONFIG : %d", val_master)
	}

	//AGEN
	val_agent_periode := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(company))
	log.Printf("REDIS DELETE AGENT PERIODE: %d", val_agent_periode)
	val_agent_pasaran := helpers.DeleteRedis("LISTPASARAN_AGENT_" + strings.ToLower(company))
	log.Printf("REDIS DELETE AGENT PASARAN: %d", val_agent_pasaran)

	if idcomppasaran > 0 {
		val_agent_pasaran_detail := helpers.DeleteRedis("LISTPASARAN_AGENT_" + strings.ToLower(company) + "_" + strconv.Itoa(idcomppasaran))
		log.Printf("REDIS DELETE AGENT PASARAN DETAIL: %d", val_agent_pasaran_detail)

	}

	//CLIENT
	val_frontend_pasaran_detail := helpers.DeleteRedis("listpasaran_" + strings.ToLower(company))
	log.Printf("REDIS DELETE FRONTEND PASARAN DETAIL: %d", val_frontend_pasaran_detail)

	if pasaranid != "" {
		val_conf_432 := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_4-3-2")
		log.Printf("REDIS DELETE FRONTEND CONFIG 4-3-2: %d", val_conf_432)
		val_colok := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_colok")
		log.Printf("REDIS DELETE FRONTEND CONFIG COLOK: %d", val_colok)
		val_5050 := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_5050")
		log.Printf("REDIS DELETE FRONTEND CONFIG 5050: %d", val_5050)
		val_macau := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_macaukombinasi")
		log.Printf("REDIS DELETE FRONTEND CONFIG MACAUKOMBINASI: %d", val_macau)
		val_dasar := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_dasar")
		log.Printf("REDIS DELETE FRONTEND CONFIG DASAR: %d", val_dasar)
		val_shio := helpers.DeleteRedis("config_" + strings.ToLower(company) + "_" + strings.ToLower(pasaranid) + "_shio")
		log.Printf("REDIS DELETE FRONTEND CONFIG SHIO: %d", val_shio)
	}
}
