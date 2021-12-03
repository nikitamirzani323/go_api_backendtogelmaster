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

type settinghome struct {
	Master string `json:"master" validate:"required"`
}
type settingsave struct {
	Maintenance_start string `json:"maintenance_start" validate:"required"`
	Maintenance_end   string `json:"maintenance_end" validate:"required"`
}
type redis_settinghome struct {
	StartMaintenance string `json:"maintenance_start"`
	EndMaintenance   string `json:"maintenance_end"`
}

func SettingHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(settinghome)
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
	field_redis := "LISTSETTING_MASTER"
	render_page := time.Now()
	var obj redis_settinghome
	var arraobj []redis_settinghome
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		maintenance_start, _ := jsonparser.GetString(value, "maintenance_start")
		maintenance_end, _ := jsonparser.GetString(value, "maintenance_end")

		obj.StartMaintenance = maintenance_start
		obj.EndMaintenance = maintenance_end
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_setting()
		helpers.SetRedis(field_redis, result, 0)
		log.Println("SETTING MYSQL")
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
		log.Println("SETTING CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func SettingSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(settingsave)
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

	result, err := models.Save_setting(client.Maintenance_start, client.Maintenance_end)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "LISTSETTING_MASTER"
	val_setting := helpers.DeleteRedis(field_redis)
	log.Printf("Redis Delete MASTER - SETTING status: %d", val_setting)
	return c.JSON(result)
}
