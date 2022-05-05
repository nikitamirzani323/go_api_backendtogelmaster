package controller

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/entities"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/helpers"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/models"
)

func CheckLogin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Login)
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

	result, err := models.Login_Model(client.Username, client.Password)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	if !result {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Username or Password Not Found",
			})

	} else {
		_deleteredis_log()
		dataclient := client.Username
		dataclient_encr, keymap := helpers.Encryption(dataclient)
		dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"token":  t,
		})

	}
}
func CheckLoginOtherWebsite(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Loginwebsite)
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

	log.Printf("PASSWORD : %s", client.Password)
	dataclient := client.Password
	dataclient_encr, keymap := helpers.Encryption(dataclient)
	dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
	log.Printf("HASH : %s", dataclient_encr_final)
	t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"token":  t,
	})
}

func Home(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Home)
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

	log.Println(client.Page)
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_username := helpers.Parsing_Decry(temp_decp, "==")
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Welcome " + client_username,
		"record":  nil,
	})

}
