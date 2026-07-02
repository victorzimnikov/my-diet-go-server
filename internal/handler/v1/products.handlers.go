package handlersV1

import (
	"errors"
	"my-diet-server/database"
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func getListQueries() *models.AppListRequest {
	var listQueries = new(models.AppListRequest)

	if listQueries.Page == 0 {
		listQueries.Page = 1
	}

	if listQueries.PerPage == 0 {
		listQueries.PerPage = 10
	}

	return listQueries
}

func GetProductsList(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	listQueries := getListQueries()

	err := c.QueryParser(listQueries)

	offset := (listQueries.Page - 1) * listQueries.PerPage

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Review your input"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var total int64
	var products []models.ProductModel

	userId := utils.GetContextUser(c).ID

	query := database.DB.Limit(listQueries.PerPage).Offset(offset).Find(&products, "user_id = ?", userId)
	query.Limit(-1).Offset(-1).Count(&total)

	paginatorInfo := utils.GetPaginatorInfo(listQueries.Page, listQueries.PerPage, total, offset)

	var newProducts []models.ProductResponseDto

	for _, value := range products {
		newProducts = append(newProducts, value.GetProductResponse())
	}

	response.Status = "success"
	response.Data = newProducts
	response.PaginatorInfo = &paginatorInfo

	return c.JSON(response)
}

func GetProduct(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	product := new(models.ProductModel)

	id := c.Params("productId")

	database.DB.Find(&product, "id = ?", id)

	if product.ID == uuid.Nil {
		err := errors.New("not found").Error()

		response.Status = "error"
		response.Error = err
		response.Message = "Продукт не найден"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Status = "success"
	response.Data = product.GetProductResponse()

	return c.JSON(response)
}

func CreateProduct(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	productBody := new(models.CreateProductDto)

	c.BodyParser(productBody)

	errors := productBody.Validate()

	if len(errors) > 0 {
		response.Status = "error"
		response.Error = errors
		response.Message = "Review your input"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product := new(models.ProductModel)

	userId := utils.GetContextUser(c).ID

	product.CreateProduct(userId, *productBody)

	err := database.DB.Create(&product).Error

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Could not create product"

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Status = "success"
	response.Data = product.GetProductResponse()

	return c.JSON(response)
}

func UpdateProduct(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	productBody := new(models.UpdateProductDto)

	c.BodyParser(productBody)

	errs := productBody.Validate()

	if len(errs) > 0 {
		response.Status = "error"
		response.Error = errs
		response.Message = "Review your input"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	id := c.Params("productId")

	product := new(models.ProductModel)

	database.DB.Find(&product, "id = ?", id)

	if product.ID == uuid.Nil {
		err := errors.New("not found").Error()

		response.Status = "error"
		response.Error = err
		response.Message = "Продукт не найден"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	product.UpdateProduct(*productBody)

	err := database.DB.Model(&product).Where("id = ?", id).Updates(product).Error

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Could not create product"

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Status = "success"
	response.Data = product.GetProductResponse()

	return c.JSON(response)
}

func DeleteProduct(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	product := new(models.ProductModel)

	id := c.Params("productId")

	err := database.DB.Delete(&product, "id = ?", id).Error

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Ошибка удаления продукта"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Status = "success"
	response.Data = product.GetProductResponse()

	return c.JSON(response)
}
