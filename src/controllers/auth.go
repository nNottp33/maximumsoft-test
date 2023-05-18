package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nNottp33/maximumsoft-test/src/configs"
	"github.com/nNottp33/maximumsoft-test/src/models"
	"golang.org/x/crypto/bcrypt"
)

type (
	AuthBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Profile struct {
		Id           int8   `json:"id"`
		Username     string `json:"username"`
		Password     string `json:"-"`
		FullName     string `json:"full_name"`
		SessionToken string `json:"session_token"`
	}
)

func SignIn(c *fiber.Ctx) error {
	profile := Profile{}
	body := new(AuthBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err.Error()})
	}

	emp := models.Employees{}
	if checkUser := configs.Db.Model(&emp).Where("username = ?", body.Username).First(&profile); checkUser.Error != nil {
		fmt.Print(checkUser.Error.Error())
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": "User not found"})
	}

	hash := []byte(profile.Password)
	password := []byte(body.Password)
	if isMatch := bcrypt.CompareHashAndPassword(hash, password); isMatch != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Password incorrect!"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        profile.Id,
		"username":  profile.Username,
		"full_name": profile.FullName,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, errJwt := token.SignedString([]byte(configs.JWT_SECRET))
	if errJwt != nil {
		return c.Status(fiber.StatusFailedDependency).JSON(fiber.Map{"errors": errJwt})
	}

	profile.SessionToken = tokenString
	profile.Password = tokenString
	isUpdate := configs.Db.Model(&emp).Where("id = ?", profile.Id).Update("session_token", profile.SessionToken)
	if isUpdate.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": isUpdate})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login success", "data": fiber.Map{"profile": profile, "is_login": true}})
}
