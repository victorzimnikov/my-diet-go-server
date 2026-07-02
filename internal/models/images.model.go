package models

import "github.com/google/uuid"

/*
	|-------------------------------|
  |-----------ImageModel----------|
	|-------------------------------|
*/

type ImageModel struct {
	BaseModel

	Width  int
	Height int
	Url    string
}

func (model *ImageModel) TableName() string {
	return "images"
}

func (model *ImageModel) ToResponse() ImageResponseDto {
	return ImageResponseDto{
		ID:     model.ID,
		Width:  model.Width,
		Height: model.Height,
		Url:    model.Url,
	}
}

type ImageResponseDto struct {
	ID     uuid.UUID `json:"id"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Url    string    `json:"url"`
}
