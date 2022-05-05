package controller

import (
	"log"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
	"github.com/nleeper/goment"
)

const Fielddashboard_home_redis = "LISTDASHBOARDWINLOSE_MASTER"

func Dashboardhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_dashboard)
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

	var obj entities.Model_dashboardwinlose_parent
	var arraobj []entities.Model_dashboardwinlose_parent
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddashboard_home_redis + "_" + client.Year)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dashboardwinlose_nmcompany, _ := jsonparser.GetString(value, "dashboardwinlose_nmcompany")
		child_RD, _, _, _ := jsonparser.Get(value, "dashboardwinlose_detail")

		var obj_child entities.Model_dashboardwinlose_child
		var arraobj_child []entities.Model_dashboardwinlose_child
		jsonparser.ArrayEach(child_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			dashboardwinlose_winlose, _ := jsonparser.GetInt(value, "dashboardwinlose_winlose")
			obj_child.Dashboardwinlose_winlose = int(dashboardwinlose_winlose)
			arraobj_child = append(arraobj_child, obj_child)
		})

		obj.Dashboardwinlose_nmcompany = dashboardwinlose_nmcompany
		obj.Dashboardwinlose_detail = arraobj_child
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_dashboardwinlose(client.Year)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielddashboard_home_redis+"_"+client.Year, result, 20*time.Minute)
		log.Println("DASHBOARD WINLOSE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("DASHBOARD WINLOSE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func DashboardCompanyPasaranWinlose(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_dashboardcompanypasaran)
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

	var obj entities.Model_companypasaran_parent
	var arraobj []entities.Model_companypasaran_parent
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddashboard_home_redis + "_" + client.Company + "_" + client.Year)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		dashboardcompanypasaran_nmpasaran, _ := jsonparser.GetString(value, "dashboardcompanypasaran_nmpasaran")
		child_RD, _, _, _ := jsonparser.Get(value, "dashboardcompanypasaran_detail")

		var obj_child entities.Model_companypasaran_child
		var arraobj_child []entities.Model_companypasaran_child
		jsonparser.ArrayEach(child_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			dashboardcompanypasaran_winlose, _ := jsonparser.GetInt(value, "dashboardcompanypasaran_winlose")
			obj_child.Dashboardcompanypasaran_winlose = int(dashboardcompanypasaran_winlose)
			arraobj_child = append(arraobj_child, obj_child)
		})

		obj.Dashboardcompanypasaran_nmpasaran = dashboardcompanypasaran_nmpasaran
		obj.Dashboardcompanypasaran_detail = arraobj_child
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_dashboardwinlosebycompany(client.Company, client.Year)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielddashboard_home_redis+"_"+client.Company+"_"+client.Year, result, 20*time.Minute)
		log.Println("DASHBOARD WINLOSE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("DASHBOARD WINLOSE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}

func _deleteredis_dashboard() {
	tglnow, _ := goment.New()
	year := tglnow.Format("YYYY")
	year_int, _ := strconv.Atoi(year)
	year_1 := int(year_int) - 1
	//MASTER
	val_master := helpers.DeleteRedis(Fielddashboard_home_redis + "_" + year)
	log.Printf("Redis Delete MASTER DASHBOARDWINLOSE : %d", val_master)
	val_master2 := helpers.DeleteRedis(Fielddashboard_home_redis + "_" + strconv.Itoa(year_1))
	log.Printf("Redis Delete MASTER DASHBOARDWINLOSE : %d", val_master2)

	_deleteredis_log()
}
