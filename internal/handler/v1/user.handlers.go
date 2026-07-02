package handlersV1

import (
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponseDto)

	response := utils.MakeAppResponse()

	response.Status = "success"
	response.Data = user

	return c.Status(fiber.StatusOK).JSON(response)
}
