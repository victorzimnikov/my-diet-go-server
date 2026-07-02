package utils

import (
	"my-diet-server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetPaginatorInfo(page int, perPage int, total int64, offset int) models.AppPaginatorInfo {
	var paginatorInfo = new(models.AppPaginatorInfo)

	paginatorInfo.Page = page
	paginatorInfo.PerPage = perPage
	paginatorInfo.Total = total
	paginatorInfo.HasMorePages = total > int64(offset+perPage)

	return *paginatorInfo
}

func GetContextUser(c *fiber.Ctx) models.UserResponseDto {
	user := c.Locals("user")

	return user.(models.UserResponseDto)
}

func MakeAppResponse() *models.NewAppResponse[any] {
	return new(models.NewAppResponse[any])
}
