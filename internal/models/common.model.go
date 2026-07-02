package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
	|-------------------------------|
  |-----------BaseModel-----------|
	|-------------------------------|
*/

type BaseModel struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key;unique"`

	CreatedAt time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"default:null;index"`
}
