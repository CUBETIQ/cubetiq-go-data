package body

import "github.com/gofiber/fiber/v2"

func BaseBodyResponse(data interface{}, status int16, msg string) *fiber.Map {
	return &fiber.Map{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}
