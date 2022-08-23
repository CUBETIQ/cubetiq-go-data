package error

import "github.com/gofiber/fiber/v2"

func BaseErrorResponse(status int16, msg string, err error) *fiber.Map {
	return &fiber.Map{
		"status": status,
		"msg":    msg,
		"error":  err,
	}
}
