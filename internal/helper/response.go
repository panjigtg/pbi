package helper

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func logResponse(c *fiber.Ctx, statusCode int, payload interface{}) {
	event := log.Info()

	if statusCode >= 500 {
		event = log.Error()
	} else if statusCode >= 400 {
		event = log.Warn()
	}

	b, _ := json.Marshal(payload)

	event.
		Str("method", c.Method()).
		Str("path", c.Path()).
		Int("status", statusCode).
		Int("size", len(b)).
		Send()
}


func Success(c *fiber.Ctx, message string, data interface{}) error {
	if data == nil {
        data = ""
    }

	response := fiber.Map{
		"status":  true,
		"message": message,
		"errors":  nil,
		"data":    data,
	}

	logResponse(c, fiber.StatusOK, response)
	return c.Status(fiber.StatusOK).JSON(response)
}


func Error(c *fiber.Ctx, statusCode int, message string, errs ...string) error {
	response := fiber.Map{
		"status":  false,
		"message": message,
		"errors":  errs,
		"data":    nil,
	}

	logResponse(c, statusCode, response)
	return c.Status(statusCode).JSON(response)
}


func BadRequest(c *fiber.Ctx, message string, errs ...string) error {
	return Error(c, fiber.StatusBadRequest, message, errs...)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnauthorized, message, "Unauthorized")
}

func Forbidden(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusForbidden, message, "Forbidden")
}

func NotFound(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusNotFound, message, "Not Found")
}

func InternalServerError(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusInternalServerError, message, "Internal Server Error")
}
