package middleware

import (
	"errors"
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminAccess(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponseDto)

	if user.Role == "admin" {
		return c.Next()
	}

	response := utils.MakeAppResponse()

	err := errors.New("no access").Error()

	response.Status = "error"
	response.Error = err
	response.Message = "No access"

	return c.Status(fiber.StatusForbidden).JSON(response)
}

func ModeratorAccess(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponseDto)

	if user.Role == "admin" || user.Role == "moderator" {
		return c.Next()
	}

	response := utils.MakeAppResponse()

	err := errors.New("no access").Error()

	response.Status = "error"
	response.Error = err
	response.Message = "No access"

	return c.Status(fiber.StatusForbidden).JSON(response)
}
