package handlersV1

import (
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func CheckServerStatus(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()

	response.Status = "success"

	return c.Status(fiber.StatusOK).JSON(response)
}
