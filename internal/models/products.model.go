package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

/*
	|-------------------------------|
  |---------ProductModel----------|
	|-------------------------------|
*/

type ProductModel struct {
	BaseModel

	Unit      ProductUnitType `json:"unit" gorm:"type:product_unit_type"` // Единица измерения продукта
	Name      string          `json:"name"`                               // Название продукта
	Amount    float32         `json:"amount"`                             // Количество продукта
	ExpiresIn time.Time       `json:"expiresIn"`                          // Срок годности продукта
	IsActive  bool            `json:"isActive"`                           // Активность (наличие) товара

	UserId uuid.UUID `json:"userId"`
}

func (ProductModel) TableName() string {
	return "products"
}

func (product *ProductModel) GetProductResponse() ProductResponseDto {
	return ProductResponseDto{
		ID:        product.ID,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Unit:      product.Unit,
		Amount:    product.Amount,
		ExpiresIn: product.ExpiresIn,
		IsActive:  product.IsActive,
	}
}

func (req *ProductModel) UpdateProduct(body UpdateProductDto) {
	req.UpdatedAt = time.Now().UTC()

	if body.Unit != nil {
		req.Unit = *body.Unit
	}

	if body.Name != nil {
		req.Name = *body.Name
	}

	if body.Amount != nil {
		req.Amount = *body.Amount
	}

	if body.IsActive != nil {
		req.IsActive = *body.IsActive
	}

	if body.ExpiresIn != nil {
		req.ExpiresIn = body.ExpiresIn.UTC()
	}
}

func (req *ProductModel) CreateProduct(userId uuid.UUID, body CreateProductDto) {
	req.ID = uuid.New()
	req.UserId = userId

	req.CreatedAt = time.Now().UTC()
	req.UpdatedAt = time.Now().UTC()

	req.Unit = body.Unit
	req.Name = body.Name
	req.Amount = body.Amount
	req.IsActive = body.IsActive
	req.ExpiresIn = body.ExpiresIn.UTC()
}

/*
	|-------------------------------|
	|-------CreateProductDto--------|
	|-------------------------------|
*/

type CreateProductDto struct {
	Unit      ProductUnitType `json:"unit"`
	Name      string          `json:"name"`
	Amount    float32         `json:"amount"`
	ExpiresIn time.Time       `json:"expiresIn"`
	IsActive  bool            `json:"isActive"`
}

func (req *CreateProductDto) Validate() []*ValidationError {
	var errors []*ValidationError

	if req.Amount <= 0 {
		var element = new(ValidationError)

		element.Field = "amount"
		element.Message = "Поле 'amount' обязательное и должно быть числом"

		errors = append(errors, element)
	}

	if !req.Unit.IsProductUnit() {
		var element = new(ValidationError)

		element.Field = "unit"
		element.Message = "Поле 'unit' должно быть типом ProductUnite"

		errors = append(errors, element)
	}

	if len(req.Name) == 0 || len(req.Name) > 25 {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'name' должна быть от 1 до 25 символов"

		errors = append(errors, element)
	}

	if req.ExpiresIn.IsZero() {
		var element = new(ValidationError)

		element.Field = "expiresIn"
		element.Message = "Неверная формат даты"

		errors = append(errors, element)
	}

	if !req.ExpiresIn.IsZero() && req.ExpiresIn.Before(time.Now().AddDate(0, 0, 1)) {
		var element = new(ValidationError)

		element.Field = "expiresIn"
		element.Message = fmt.Sprintf("Дата не может быть меньше %s", time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05"))

		errors = append(errors, element)
	}

	return errors
}

/*
	|-------------------------------|
	|-------UpdateProductDto--------|
	|-------------------------------|
*/

type UpdateProductDto struct {
	Unit      *ProductUnitType `json:"unit"`
	Name      *string          `json:"name"`
	Amount    *float32         `json:"amount"`
	ExpiresIn *time.Time       `json:"expiredIn"`
	IsActive  *bool            `json:"isActive"`
}

func (req *UpdateProductDto) Validate() []*ValidationError {
	var errors []*ValidationError

	if req.Unit != nil && !req.Unit.IsProductUnit() {
		var element = new(ValidationError)

		element.Field = "unite"
		element.Message = "'unite' must be ProductUnite type"

		errors = append(errors, element)
	}

	if req.Name != nil && (len(*req.Name) == 0 || len(*req.Name) > 25) {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'name' должна быть от 1 до 25 символов"

		errors = append(errors, element)
	}

	if req.Amount != nil && *req.Amount <= 0 {
		var element = new(ValidationError)

		element.Field = "amount"
		element.Message = "Поле 'amount' обязательное и должно быть числом"

		errors = append(errors, element)
	}

	if req.ExpiresIn != nil && req.ExpiresIn.Before(time.Now().AddDate(0, 0, 1)) {
		var element = new(ValidationError)

		element.Field = "expiresIn"
		element.Message = fmt.Sprintf("Дата не может быть меньше %s", time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05"))

		errors = append(errors, element)
	}

	return errors
}

/*
	|-------------------------------|
	|------ProductResponseDto-------|
	|-------------------------------|
*/

type ProductResponseDto struct {
	ID        uuid.UUID       `json:"id"`
	Unit      ProductUnitType `json:"unit"`
	Name      string          `json:"name"`
	Amount    float32         `json:"amount"`
	ExpiresIn time.Time       `json:"expiresIn"`
	IsActive  bool            `json:"isActive"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
