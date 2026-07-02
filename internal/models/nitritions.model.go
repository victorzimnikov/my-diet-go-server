package models

import (
	"github.com/google/uuid"
)

type NutritionModel struct {
	BaseModel

	Energy        float32
	Proteins      float32
	Carbohydrates float32
	Fats          float32
	Vitamins      float32
	Minerals      float32
	UserID        uuid.UUID
}

func (model *NutritionModel) TableName() string {
	return "nutrition"
}

func (model *NutritionModel) ToCreate(userId uuid.UUID, body NutritionCreateBodyDto) {
	model.UserID = userId
	model.Energy = body.Energy
	model.Proteins = body.Proteins
	model.Carbohydrates = body.Carbohydrates
	model.Fats = body.Fats
	model.Vitamins = body.Vitamins
	model.Minerals = body.Minerals
	model.UserID = body.UserID
}

func (model *NutritionModel) ToUpdate(userId uuid.UUID, body NutritionUpdateBodyDto) {
	if body.Carbohydrates != nil {
		model.Carbohydrates = *body.Carbohydrates
	}

	if body.Energy != nil {
		model.Energy = *body.Energy
	}

	if body.Proteins != nil {
		model.Proteins = *body.Proteins
	}

	if body.Fats != nil {
		model.Fats = *body.Fats
	}

	if body.Vitamins != nil {
		model.Vitamins = *body.Vitamins
	}

	if body.Minerals != nil {
		model.Minerals = *body.Minerals
	}
}

type NutritionResponseDto struct {
	ID            uuid.UUID `json:"id"`
	Energy        float32   `json:"energy"`
	Proteins      float32   `json:"proteins"`
	Carbohydrates float32   `json:"carbohydrates"`
	Fats          float32   `json:"fats"`
	Vitamins      float32   `json:"vitamins"`
	Minerals      float32   `json:"minerals"`
}

type NutritionCreateBodyDto struct {
	ID            *uuid.UUID
	Energy        float32
	Proteins      float32
	Carbohydrates float32
	Fats          float32
	Vitamins      float32
	Minerals      float32
	UserID        uuid.UUID
}

type NutritionUpdateBodyDto struct {
	ID            *uuid.UUID
	Energy        *float32
	Proteins      *float32
	Carbohydrates *float32
	Fats          *float32
	Vitamins      *float32
	Minerals      *float32
	UserID        *uuid.UUID
}
