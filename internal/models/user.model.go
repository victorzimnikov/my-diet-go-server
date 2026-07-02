package models

import (
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
)

/*
	|-------------------------------|
  |----------UserModel------------|
	|-------------------------------|
*/

type UserModel struct {
	BaseModel

	Name            string     `gorm:"type:varchar(50)"`
	Email           string     `gorm:"type:varchar(320);uniqueIndex"`
	Password        string     `gorm:"type:varchar(100)"`
	Role            string     `gorm:"type:varchar(50);default:'user'"`
	IsEmailSend     bool       `gorm:"default:false"`
	IsEmailVerified bool       `gorm:"default:false"`
	EmailSentAt     *time.Time `gorm:"default:null"`
}

func (UserModel) TableName() string {
	return "users"
}

func (user *UserModel) GetUserResponse() UserResponseDto {
	return UserResponseDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user *UserModel) CreateUser(hashedPassword []byte, body RegisterUserDto) {
	user.Name = body.Name
	user.Email = strings.ToLower(body.Email)
	user.Password = string(hashedPassword)
	user.Role = "user"
	user.IsEmailSend = false
	user.IsEmailVerified = false
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
}

/*
	|-------------------------------|
  |--------RegisterUserDto--------|
	|-------------------------------|
*/

type RegisterUserDto struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (req *RegisterUserDto) Validate() []*ValidationError {
	var errors []*ValidationError

	if len(req.Name) < 5 || len(req.Name) > 50 {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'name' должна быть от 5 до 50 символов"

		errors = append(errors, element)
	}

	if len(req.Password) < 8 || len(req.Password) > 50 {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'password' должна быть от 8 до 100 символов"

		errors = append(errors, element)
	}

	if len(req.PasswordConfirm) < 8 || len(req.PasswordConfirm) > 50 {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'passwordConfirm' должна быть от 8 до 100 символов"

		errors = append(errors, element)
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		var element = new(ValidationError)

		element.Field = "email"
		element.Message = "Неверный формат поля 'email'"

		errors = append(errors, element)
	}

	return errors
}

/*
	|-------------------------------|
  |---------LoginUserDto----------|
	|-------------------------------|
*/

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginUserDto) Validate() []*ValidationError {
	var errors []*ValidationError

	if len(req.Password) < 8 || len(req.Password) > 50 {
		var element = new(ValidationError)

		element.Field = "name"
		element.Message = "Длина поля 'password' должна быть от 8 до 100 символов"

		errors = append(errors, element)
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		var element = new(ValidationError)

		element.Field = "email"
		element.Message = "Неверный формат поля 'email'"

		errors = append(errors, element)
	}

	return errors
}

/*
	|-------------------------------|
  |--------RefreshTokenDto--------|
	|-------------------------------|
*/

type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}

func (req *RefreshTokenDto) Validate() []*ValidationError {
	var errors []*ValidationError

	if len(req.RefreshToken) == 0 {
		var element = new(ValidationError)

		element.Field = "refreshToken"
		element.Message = "Поля 'refreshToken' обязательное"

		errors = append(errors, element)
	}

	return errors
}

/*
	|-------------------------------|
  |--------UserResponseDto--------|
	|-------------------------------|
*/

type UserResponseDto struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
