package repository

import (
	"my-diet-server/database"
	"my-diet-server/internal/models"

	"github.com/google/uuid"
)

func CreateNutrition(userId uuid.UUID, body models.NutritionCreateBodyDto) *models.NutritionModel {
	var nutrition models.NutritionModel

	nutrition.ToCreate(userId, body)

	database.DB.Create(&nutrition)

	return &nutrition
}
