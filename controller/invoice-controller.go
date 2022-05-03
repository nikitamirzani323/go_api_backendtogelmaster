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

type redis_invoicehome struct {
	Idinvoice string `json:"invoice_id"`
	Company   string `json:"invoice_company"`
	Date      string `json:"invoice_date"`
	Name      string `json:"invoice_name"`
	Winlose   int    `json:"invoice_winlose"`
	Status    string `json:"invoice_status"`
	Statuscss string `json:"invoice_statuscss"`
}

const Fieldinvoice_home_redis = "LISTINVOICE_MASTER"

func InvoiceHome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoicehome)
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
	var obj redis_invoicehome
	var arraobj []redis_invoicehome
	resultredis, flag := helpers.GetRedis(Fieldinvoice_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		invoice_id, _ := jsonparser.GetString(value, "invoice_id")
		invoice_company, _ := jsonparser.GetString(value, "invoice_company")
		invoice_date, _ := jsonparser.GetString(value, "invoice_date")
		invoice_name, _ := jsonparser.GetString(value, "invoice_name")
		invoice_winlose, _ := jsonparser.GetInt(value, "invoice_winlose")
		invoice_status, _ := jsonparser.GetString(value, "invoice_status")
		invoice_statuscss, _ := jsonparser.GetString(value, "invoice_statuscss")

		obj.Idinvoice = invoice_id
		obj.Company = invoice_company
		obj.Date = invoice_date
		obj.Name = invoice_name
		obj.Winlose = int(invoice_winlose)
		obj.Status = invoice_status
		obj.Statuscss = invoice_statuscss
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_invoice()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldinvoice_home_redis, result, 60*time.Minute)
		log.Println("INVOICE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("INVOICE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func InvoiceDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoicedetail)
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
	var obj redis_invoicehome
	var arraobj []redis_invoicehome
	resultredis, flag := helpers.GetRedis(Fieldinvoice_home_redis + "_" + client.Invoice)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		invoice_id, _ := jsonparser.GetString(value, "invoice_id")
		invoice_company, _ := jsonparser.GetString(value, "invoice_company")
		invoice_date, _ := jsonparser.GetString(value, "invoice_date")
		invoice_name, _ := jsonparser.GetString(value, "invoice_name")
		invoice_winlose, _ := jsonparser.GetInt(value, "invoice_winlose")
		invoice_status, _ := jsonparser.GetString(value, "invoice_status")
		invoice_statuscss, _ := jsonparser.GetString(value, "invoice_statuscss")

		obj.Idinvoice = invoice_id
		obj.Company = invoice_company
		obj.Date = invoice_date
		obj.Name = invoice_name
		obj.Winlose = int(invoice_winlose)
		obj.Status = invoice_status
		obj.Statuscss = invoice_statuscss
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_invoicedetail(client.Invoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldinvoice_home_redis+"_"+client.Invoice, result, 60*time.Minute)
		log.Println("INVOICE DETAIL MYSQL")
		return c.JSON(result)
	} else {
		log.Println("INVOICE DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func InvoiceSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoicesave)
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

	result, err := models.Save_invoice(client.Sdata, client.Master, client.Periode)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_invoice()
	return c.JSON(result)
}
func InvoiceSavewinlosestatus(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoicesavestatus)
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

	result, err := models.Save_invoicestatus(client.Master, client.Invoice, client.Tipe)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	val_master := helpers.DeleteRedis(Fieldinvoice_home_redis)
	log.Printf("Redis Delete MASTER LISTINVOICE_MASTER : %d", val_master)
	return c.JSON(result)
}
func InvoiceSavePasaran(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_invoicesavepasaran)
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

	result, err := models.Save_company_listpasaran(client.Master, client.Invoice)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_invoice()
	return c.JSON(result)
}
func _deleteredis_invoice() {
	//MASTER
	val_master := helpers.DeleteRedis(Fieldinvoice_home_redis)
	log.Printf("Redis Delete MASTER INVOICE : %d", val_master)
}
