package repository

import (
	"my-diet-server/internal/models"

	"github.com/google/uuid"
)

func CreateStep(userId uuid.UUID, body models.RecipeStepCreateBodyDto) *models.RecipeStepModel {
	var step models.RecipeStepModel

	return &step
}

func CreateSteps(userId uuid.UUID, body []models.RecipeStepCreateBodyDto) *[]models.RecipeStepModel {
	var steps []models.RecipeStepModel

	for _, step := range body {
		steps = append(steps, *CreateStep(userId, step))
	}

	return &steps
}
