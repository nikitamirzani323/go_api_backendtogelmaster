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

const Fieldlistpasaran_home_redis = "LISTPASARANWAJIB_MASTER"
const Fieldlistprediksiwajib_home_redis = "LISTPREDIKSIWAJIB_MASTER"

func Listpasaranwajib(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_pasaranlist)
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

	var obj entities.Model_pasaranlist
	var arraobj []entities.Model_pasaranlist
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistpasaran_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaranlist_idpasarantogel, _ := jsonparser.GetString(value, "pasaranlist_idpasarantogel")
		pasaranlist_nmpasarantogel, _ := jsonparser.GetString(value, "pasaranlist_nmpasarantogel")

		obj.Pasaranlist_idpasarantogel = pasaranlist_idpasarantogel
		obj.Pasaranlist_nmpasarantogel = pasaranlist_nmpasarantogel
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_listpasaranwajib()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistpasaran_home_redis, result, 24*time.Hour)
		log.Println("LIST PASARAN WAJIB MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LIST PASARAN WAJIB CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}

func Prediksiwajib(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_prediksi)
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

	var obj entities.Model_prediksiwajib
	var arraobj []entities.Model_prediksiwajib
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistprediksiwajib_home_redis + "_" + client.Idpasarantogel + "_" + client.Nomorprediksi)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		prediksi_idcompany, _ := jsonparser.GetString(value, "prediksi_idcompany")
		prediksi_nmcompany, _ := jsonparser.GetString(value, "prediksi_nmcompany")
		prediksi_totalbet, _ := jsonparser.GetInt(value, "prediksi_totalbet")
		prediksi_subtotal, _ := jsonparser.GetInt(value, "prediksi_subtotal")
		prediksi_subtotalwin, _ := jsonparser.GetInt(value, "prediksi_subtotalwin")
		prediksi_result, _, _, _ := jsonparser.Get(value, "prediksi_result")

		var objdetail entities.Model_listPrediksi
		var arraobjdetail []entities.Model_listPrediksi
		jsonparser.ArrayEach(prediksi_result, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			prediksi_tanggal, _ := jsonparser.GetString(value, "prediksi_tanggal")
			prediksi_username, _ := jsonparser.GetString(value, "prediksi_username")
			prediksi_permainan, _ := jsonparser.GetString(value, "prediksi_permainan")
			prediksi_nomor, _ := jsonparser.GetString(value, "prediksi_nomor")
			prediksi_bet, _ := jsonparser.GetInt(value, "prediksi_bet")

			objdetail.Prediksi_tanggal = prediksi_tanggal
			objdetail.Prediksi_username = prediksi_username
			objdetail.Prediksi_permainan = prediksi_permainan
			objdetail.Prediksi_nomor = prediksi_nomor
			objdetail.Prediksi_bet = int(prediksi_bet)
			arraobjdetail = append(arraobjdetail, objdetail)
		})

		obj.Prediksi_idcompany = prediksi_idcompany
		obj.Prediksi_nmcompany = prediksi_nmcompany
		obj.Prediksi_result = arraobjdetail
		obj.Prediksi_totalbet = int(prediksi_totalbet)
		obj.Prediksi_subtotal = int(prediksi_subtotal)
		obj.Prediksi_subtotalwin = int(prediksi_subtotalwin)
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_prediksipasaranwajib(client.Idpasarantogel, client.Nomorprediksi)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistprediksiwajib_home_redis+"_"+client.Idpasarantogel+"_"+client.Nomorprediksi, result, 30*time.Minute)
		log.Println("RESULT PREDIKSI WAJIB MYSQL")
		return c.JSON(result)
	} else {
		log.Println("RESULT PREDIKSI WAJIB CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
