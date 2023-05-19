package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type QuizBody struct {
	A []string `json:"array_a"`
	B []string `json:"array_b"`
}

func SecondQuiz(c *fiber.Ctx) error {
	body := new(QuizBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": err.Error()})
	}

	merge, unique := FindAnswers(body.A, body.B)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success", "result": fiber.Map{"merge_array": merge, "unique_array": unique}})
}

func FindAnswers(a []string, b []string) ([]string, []string) {
	uniqueMap := make(map[string]bool)

	for _, item := range a {
		uniqueMap[item] = true
	}

	for _, item := range b {
		uniqueMap[item] = true
	}

	x := []string{}
	for key := range uniqueMap {
		x = append(x, key)
	}

	y := []string{}
	for _, item := range a {
		if !contains(b, item) {
			y = append(y, item)
		}
	}

	for _, item := range b {
		if !contains(a, item) {
			y = append(y, item)
		}
	}

	return x, y
}

func contains(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}
