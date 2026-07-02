package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

/*
	|-------------------------------|
  |----------RecipeModel----------|
	|-------------------------------|
*/

type RecipeModel struct {
	BaseModel

	FoodIntakes pq.StringArray `gorm:"type:food_intake_type[]"` // Время приема рецепта в пищу
	Name        string         // Название рецепта
	Description string         // Описание рецепта

	NutritionID *uuid.UUID
	Nutrition   *NutritionModel
	Steps       []RecipeStepModel       `gorm:"foreignKey:RecipeID"` // Шаги приготовления
	Ingredients []RecipeIngredientModel `gorm:"foreignKey:RecipeID"` // Ингредиенты

	UserId uuid.UUID
}

func (RecipeModel) TableName() string {
	return "recipes"
}

func (recipe *RecipeModel) ToResponse() RecipeResponseDto {
	var ingredientsDto []RecipeIngredientResponseDto

	for _, ingredient := range recipe.Ingredients {
		var ingredientDto RecipeIngredientResponseDto

		ingredientDto.Amount = ingredient.Amount
		ingredientDto.Description = ingredient.Description
		ingredientDto.ID = ingredient.ID
		ingredientDto.Product = ingredient.Product.GetProductResponse()
		ingredientDto.Unit = ingredient.Unit

		ingredientsDto = append(ingredientsDto, ingredientDto)
	}

	var stepsDto []RecipeStepResponseDto

	for _, step := range recipe.Steps {
		var stepDto RecipeStepResponseDto

		stepDto.Description = step.Description
		stepDto.ID = step.ID
		stepDto.Number = step.Number
		stepDto.Title = step.Title

		if len(step.Images) > 0 {
			var imagesDto []ImageResponseDto

			for _, image := range step.Images {
				imageDto := image.ToResponse()

				imagesDto = append(imagesDto, imageDto)
			}

			stepDto.Images = imagesDto
		}

		stepsDto = append(stepsDto, stepDto)
	}

	nutrition := new(NutritionResponseDto)

	if recipe.Nutrition != nil {
		nutrition.Carbohydrates = recipe.Nutrition.Carbohydrates
		nutrition.Energy = recipe.Nutrition.Energy
		nutrition.Fats = recipe.Nutrition.Fats
		nutrition.ID = recipe.Nutrition.ID
		nutrition.Minerals = recipe.Nutrition.Minerals
		nutrition.Proteins = recipe.Nutrition.Proteins
		nutrition.Vitamins = recipe.Nutrition.Vitamins
	} else {
		nutrition = nil
	}

	return RecipeResponseDto{
		ID:          recipe.ID,
		CreatedAt:   recipe.CreatedAt,
		UpdatedAt:   recipe.UpdatedAt,
		Name:        recipe.Name,
		Description: recipe.Description,
		Steps:       stepsDto,
		Ingredients: ingredientsDto,
		Nutrition:   nutrition,
	}
}

func (req *RecipeModel) ToCreate(userId uuid.UUID, body RecipeCreateBodyDto) {
	req.UserId = userId

	req.Name = body.Name
	req.Description = body.Description

	var steps []RecipeStepModel

	for _, stepDto := range body.Steps {
		var step RecipeStepModel

		step.ToCreate(userId, stepDto)

		steps = append(steps, step)
	}

	req.Steps = steps

	var ingredients []RecipeIngredientModel

	for _, ingredientDto := range body.Ingredients {
		var ingredient RecipeIngredientModel

		ingredient.Amount = ingredientDto.Amount
		ingredient.Description = ingredientDto.Description
		ingredient.ProductID = ingredientDto.ProductID
		ingredient.Unit = ingredientDto.Unit

		ingredients = append(ingredients, ingredient)
	}

	req.Ingredients = ingredients

	if req.Nutrition != nil {
		req.NutritionID = &req.Nutrition.ID
	}

	req.FoodIntakes = body.FoodIntakes
}

func (model *RecipeModel) ToUpdate(body RecipeUpdateBodyDto) {
	if body.Name != nil {
		model.Name = *body.Name
	}

	if body.Description != nil {
		model.Description = *body.Description
	}
}

/*
	|-------------------------------|
	|-------RecipeResponseDto-------|
	|-------------------------------|
*/

type RecipeResponseDto struct {
	ID          uuid.UUID                     `json:"id"`
	FoodIntakes FoodIntakeType                `json:"foodIntakes"`
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	CreatedAt   time.Time                     `json:"createdAt"`
	UpdatedAt   time.Time                     `json:"updatedAt"`
	Ingredients []RecipeIngredientResponseDto `json:"ingredients"`
	Steps       []RecipeStepResponseDto       `json:"steps"`
	Nutrition   *NutritionResponseDto         `json:"nutrition,omitempty"`
}

/*
	|-------------------------------|
	|---------RecipeBodyDto---------|
	|-------------------------------|
*/

type RecipeCreateBodyDto struct {
	FoodIntakes pq.StringArray
	Name        string
	Description string
	Steps       []RecipeStepCreateBodyDto
	Ingredients []RecipeIngredientBodyDto
	Nutrition   *NutritionCreateBodyDto
}

func (req *RecipeCreateBodyDto) Validate() []*ValidationError {
	return MakeValidationErrors(
		new(ValidationError).New(len(req.FoodIntakes) <= 0, "FoodIntakes", "Поле 'FoodIntakes' обязательное", nil),
		new(ValidationError).New(len(req.Name) == 0 || len(req.Name) > 25, "name", "Длина поля 'name' должна быть от 1 до 25 символов", nil),
		new(ValidationError).New(len(req.Description) == 0 || len(req.Description) > 255, "description", "Длина поля 'description' должна быть от 1 до 255 символов", nil),
		new(ValidationError).New(len(req.Ingredients) <= 0, "ingredients", "Поле 'ingredients' обязательное", nil),
		new(ValidationError).New(len(req.Steps) <= 0, "steps", "Поле 'steps' обязательное", nil),
	)
}

type RecipeUpdateBodyDto struct {
	FoodIntakes *pq.StringArray
	Name        *string
	Description *string
	Steps       *[]RecipeStepCreateBodyDto
	Ingredients *[]RecipeIngredientBodyDto
	Nutrition   *NutritionCreateBodyDto
}

func (body *RecipeUpdateBodyDto) Validate() []*ValidationError {
	isFoodIntakeValid := true
	isNameValid := true
	isDescriptionValid := true
	isIngredientsValid := true
	isStepsValid := true

	if body.FoodIntakes != nil {
		isFoodIntakeValid = len(*body.FoodIntakes) <= 0
	}

	if body.Name != nil {
		isNameValid = len(*body.Name) == 0 || len(*body.Name) > 25
	}

	if body.Description != nil {
		isDescriptionValid = len(*body.Description) == 0 || len(*body.Description) > 255
	}

	if body.Ingredients != nil {
		isIngredientsValid = len(*body.Ingredients) <= 0
	}

	if body.Steps != nil {
		isStepsValid = len(*body.Steps) <= 0
	}

	return MakeValidationErrors(
		new(ValidationError).New(
			isFoodIntakeValid,
			"FoodIntakes",
			"Поле 'FoodIntakes' обязательное",
			nil,
		),
		new(ValidationError).New(
			isNameValid,
			"name",
			"Длина поля 'name' должна быть от 1 до 25 символов",
			nil,
		),
		new(ValidationError).New(
			isDescriptionValid,
			"description",
			"Длина поля 'description' должна быть от 1 до 255 символов",
			nil,
		),
		new(ValidationError).New(
			isIngredientsValid,
			"ingredients",
			"Поле 'ingredients' обязательное",
			nil,
		),
		new(ValidationError).New(
			isStepsValid,
			"steps",
			"Поле 'steps' обязательное",
			nil,
		),
	)
}

/*
	|-------------------------------|
	|-----RecipeIngredientModel-----|
	|-------------------------------|
*/

type RecipeIngredientModel struct {
	BaseModel

	Unit        IngredientUnitType
	Amount      float32
	Description string
	ProductID   uuid.UUID
	Product     ProductModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RecipeID    uuid.UUID
	UserID      uuid.UUID
}

func (RecipeIngredientModel) TableName() string {
	return "recipe_ingredients"
}

/*
	|-------------------------------|
	|------RecipeIngredientDto------|
	|-------------------------------|
*/

type RecipeIngredientResponseDto struct {
	ID          uuid.UUID          `json:"id"`
	Unit        IngredientUnitType `json:"unit"`
	Amount      float32            `json:"amount"`
	Description string             `json:"description"`
	Product     ProductResponseDto `json:"product"`
}

type RecipeIngredientBodyDto struct {
	Unit        IngredientUnitType
	Amount      float32
	Description string
	ProductID   uuid.UUID
}

func (value *RecipeIngredientBodyDto) Validate(index int) []*ValidationError {
	var errors []*ValidationError

	if !value.Unit.IsIngredientUnitType() {
		var element = new(ValidationError)

		element.Position = index
		element.Field = "unit"
		element.Message = "Поле 'unit' должно быть типом IngredientUnitType"

		errors = append(errors, element)
	}

	if value.Amount <= 0 {
		var element = new(ValidationError)

		element.Position = index
		element.Field = "amount"
		element.Message = "Поле 'amount' обязательно для заполнения и должно быть больше 0"

		errors = append(errors, element)
	}

	return errors
}

func (value *RecipeIngredientBodyDto) ToResponseDto() RecipeIngredientResponseDto {
	return RecipeIngredientResponseDto{
		Unit:        value.Unit,
		Amount:      value.Amount,
		Description: value.Description,
	}
}

/*
	|-------------------------------|
  |--------RecipeStepModel--------|
	|-------------------------------|
*/

type RecipeStepModel struct {
	BaseModel

	Title       string
	Description string
	Number      int
	Images      []ImageModel `gorm:"many2many:steps_images;"`
	UserID      uuid.UUID
	RecipeID    uuid.UUID
}

func (model *RecipeStepModel) TableName() string {
	return "recipe_steps"
}

func (model *RecipeStepModel) ToCreate(userId uuid.UUID, body RecipeStepCreateBodyDto) {
	model.UserID = userId

	model.Title = body.Title
	model.Description = body.Description
	model.Number = body.Number

	if len(body.ImageIds) > 0 {
		var images []ImageModel

		for _, value := range body.ImageIds {
			var image ImageModel

			image.ID = value

			images = append(images, image)
		}

		model.Images = images
	}
}

/*
	|-------------------------------|
	|---------RecipeStepDto---------|
	|-------------------------------|
*/

type RecipeStepResponseDto struct {
	ID          uuid.UUID          `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Number      int                `json:"number"`
	Images      []ImageResponseDto `json:"images"`
}

type RecipeStepCreateBodyDto struct {
	Title       string
	Description string
	Number      int
	ImageIds    []uuid.UUID
}

func (value *RecipeStepCreateBodyDto) Validate(index int) []*ValidationError {
	return []*ValidationError{
		new(ValidationError).New(len(value.Title) <= 0 || len(value.Title) > 25, "title", "Длина поля 'title' должна быть от 1 до 25 символов", &index),
		new(ValidationError).New(value.Number <= 0, "number", "Поле 'number' обязательно для заполнения и должно быть больше 0", &index),
	}
}

func (value *RecipeStepCreateBodyDto) ToResponseDto() RecipeStepResponseDto {
	return RecipeStepResponseDto{
		Title:       value.Title,
		Description: value.Description,
		Number:      value.Number,
	}
}

type RecipeStepUpdateBodyDto struct {
	Title       *string
	Description *string
	Number      *int
	ImageIds    *[]uuid.UUID
}
