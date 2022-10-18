package body

import (
	"github.com/cubetiq/cubetiq-data-go/model/page"
	"github.com/gofiber/fiber/v2"
)

func BaseBodyResponse(data interface{}, status int16, msg string) *fiber.Map {
	return &fiber.Map{
		"data":   data,
		"status": status,
		"msg":    msg,
	}
}

func BasePageBodyResponse(data interface{}, page page.Page, status int16, msg string) *fiber.Map {
	return &fiber.Map{
		"data":   data,
		"page":   page,
		"status": status,
		"msg":    msg,
	}
}
