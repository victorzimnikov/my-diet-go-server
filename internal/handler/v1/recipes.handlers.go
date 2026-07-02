package handlersV1

import (
	"errors"
	"my-diet-server/database"
	"my-diet-server/internal/models"
	"my-diet-server/internal/repository"
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateRecipe(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	recipeBody := new(models.RecipeCreateBodyDto)

	c.BodyParser(recipeBody)

	errors := recipeBody.Validate()

	if len(errors) > 0 {
		response.Status = "error"

		response.Error = errors
		response.Message = "Review your input"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	userId := utils.GetContextUser(c).ID

	recipe := new(models.RecipeModel)

	if recipeBody.Nutrition != nil {
		recipe.Nutrition = repository.CreateNutrition(userId, *recipeBody.Nutrition)
	}

	recipe.ToCreate(userId, *recipeBody)

	err := database.DB.Preload("Ingredients.Product").Preload("Steps.Images").Preload("Nutrition").Create(&recipe).Find(&recipe).Error

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Could not create recipe"

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Status = "success"
	response.Data = recipe.ToResponse()

	return c.JSON(response)
}

func GetRecipesList(c *fiber.Ctx) error {
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
	var recipes []models.RecipeModel

	userId := utils.GetContextUser(c).ID

	query := database.DB.Limit(listQueries.PerPage).Offset(offset).Preload("Ingredients.Product").Preload("Steps.Images").Preload("Nutrition").Find(&recipes, "user_id = ?", userId)
	query.Limit(-1).Offset(-1).Count(&total).Association("Ingredients")

	paginatorInfo := utils.GetPaginatorInfo(listQueries.Page, listQueries.PerPage, total, offset)

	var newRecipes []models.RecipeResponseDto
	var validationErrors []*models.ValidationError

	for _, value := range recipes {
		newRecipes = append(newRecipes, value.ToResponse())
	}

	if len(validationErrors) > 0 {
		response.Status = "error"
		response.Error = validationErrors

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Status = "success"
	response.Data = newRecipes
	response.PaginatorInfo = &paginatorInfo

	return c.JSON(response)
}

func GetRecipe(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	recipe := new(models.RecipeModel)

	id := c.Params("recipeId")

	database.DB.Preload("Ingredients.Product").Preload("Steps.Images").Preload("Nutrition").Find(&recipe, "id = ?", id)

	if recipe.ID == uuid.Nil {
		err := errors.New("not found").Error()

		response.Status = "error"
		response.Error = err
		response.Message = "Рецепт не найден"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Status = "success"
	response.Data = recipe.ToResponse()

	return c.JSON(response)
}

func DeleteRecipe(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	recipe := new(models.RecipeModel)

	id := c.Params("recipeId")

	err := database.DB.Delete(&recipe, "id = ?", id).Error

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Ошибка удаления продукта"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response.Status = "success"
	response.Data = fiber.Map{"id": id}

	return c.JSON(response)
}

func UpdateRecipe(c *fiber.Ctx) error {
	// recipeId := c.Params("recipeId")

	response := utils.MakeAppResponse()
	recipeBody := new(models.RecipeCreateBodyDto)

	c.BodyParser(recipeBody)

	response.Status = "success"
	response.Data = response

	return c.JSON(response)
}
