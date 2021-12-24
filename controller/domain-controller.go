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

const Fielddomain_home_redis = "LISTDOMAIN_MASTER"

func Domainhome(c *fiber.Ctx) error {
	var obj entities.Model_domain
	var arraobj []entities.Model_domain
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fielddomain_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		domain_iddomain, _ := jsonparser.GetInt(value, "domain_iddomain")
		domain_name, _ := jsonparser.GetString(value, "domain_name")
		domain_status, _ := jsonparser.GetString(value, "domain_status")
		domain_create, _ := jsonparser.GetString(value, "domain_create")
		domain_update, _ := jsonparser.GetString(value, "domain_update")

		obj.Domain_iddomain = int(domain_iddomain)
		obj.Domain_name = domain_name
		obj.Domain_status = domain_status
		obj.Domain_create = domain_create
		obj.Domain_update = domain_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_domain()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fielddomain_home_redis, result, 60*time.Minute)
		log.Println("DOMAIN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("DOMAIN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func DomainSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_domainsave)
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

	result, err := models.Save_domain(
		client.Master,
		client.Domain_name, client.Domain_status, client.Sdata, client.Domain_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	val_master := helpers.DeleteRedis(Fielddomain_home_redis)
	log.Printf("Redis Delete MASTER DOMAIN : %d", val_master)
	return c.JSON(result)
}
