package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/nNottp33/maximumsoft-test/src/configs"
	"github.com/nNottp33/maximumsoft-test/src/models"
	"golang.org/x/crypto/bcrypt"
)

type EmployeesBody struct {
	Id           int8      `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	FullName     string    `json:"full_name"`
	Position     string    `json:"position"`
	SessionToken string    `json:"session_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" `
	IsActive     bool      `json:"is_active"`
}

func GetEmployees(c *fiber.Ctx) error {
	emp := []models.Employees{}

	result := configs.Db.Find(&emp)
	if result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data is empty!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully",
		"data":    emp,
	})
}

func NewEmployees(c *fiber.Ctx) error {
	body := new(EmployeesBody)

	if ok := c.BodyParser(body); ok != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": ok})
	}

	emp := models.Employees{}
	if errCopy := copier.Copy(&emp, &body); errCopy != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errCopy,
		})
	}

	salt, errHash := bcrypt.GenerateFromPassword([]byte(body.Password), 11)
	if errHash != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errHash,
		})
	}

	emp.Password = string(salt)
	result := configs.Db.Create(&emp)
	if result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully",
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	body := new(EmployeesBody)

	if ok := c.BodyParser(body); ok != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": ok})
	}

	emp := models.Employees{}
	if errCopy := copier.Copy(&emp, &body); errCopy != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{

			"errors": errCopy,
		})
	}

	salt, errHash := bcrypt.GenerateFromPassword([]byte(body.Password), 11)
	if errHash != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errHash,
		})
	}

	emp.Password = string(salt)
	if result := configs.Db.Save(&emp); result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Updated Employee success",
	})
}

func RemoveEmployee(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	emp := models.Employees{}
	if result := configs.Db.Where("id = ?", id).Delete(&emp); result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted Employee success",
	})
}
