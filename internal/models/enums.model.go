package models

import "slices"

/*
	|-------------------------------|
  |-------ProductUnitEnum---------|
	|-------------------------------|
*/

type ProductUnitType string

const (
	Gram       ProductUnitType = "g"
	Kilogram   ProductUnitType = "kg"
	Milliliter ProductUnitType = "ml"
	Liter      ProductUnitType = "l"
)

func (value *ProductUnitType) IsProductUnit() bool {
	productUnits := []ProductUnitType{Gram, Kilogram, Milliliter, Liter}

	return slices.Contains(productUnits, *value)
}

/*
	|-------------------------------|
  |-----------MenuEnum------------|
	|-------------------------------|
*/

type MenuType string

const (
	Week    MenuType = "week"
	TwoWeek MenuType = "two_week"
	Month   MenuType = "month"
)

func (value *MenuType) IsMenuType() bool {
	menuTypes := []MenuType{Week, TwoWeek, Month}

	return slices.Contains(menuTypes, *value)
}

/*
	|-------------------------------|
  |------IngredientUnitEnum-------|
	|-------------------------------|
*/

type IngredientUnitType string

const (
	IngredientTablespoon  IngredientUnitType = "tablespoon"
	IngredientTeaspoon    IngredientUnitType = "teaspoon"
	IngredientDesertSpoon IngredientUnitType = "desert_spoon"
	IngredientGlass       IngredientUnitType = "glass"
	IngredientGram        IngredientUnitType = "g"
	IngredientKilogram    IngredientUnitType = "kg"
	IngredientMilliliter  IngredientUnitType = "ml"
	IngredientLiter       IngredientUnitType = "l"
	IngredientPiece       IngredientUnitType = "piece"
	IngredientTaste       IngredientUnitType = "taste"
)

func (value *IngredientUnitType) IsIngredientUnitType() bool {
	ingredientUnits := []IngredientUnitType{
		IngredientTablespoon,
		IngredientTeaspoon,
		IngredientDesertSpoon,
		IngredientGlass,
		IngredientGram,
		IngredientKilogram,
		IngredientMilliliter,
		IngredientLiter,
		IngredientPiece,
	}

	return slices.Contains(ingredientUnits, *value)
}

/*
	|-------------------------------|
  |--------FoodIntakeEnum---------|
	|-------------------------------|
*/

type FoodIntakeType string

const (
	Breakfast      FoodIntakeType = "breakfast"
	FirstSnack     FoodIntakeType = "first_snack"
	Launch         FoodIntakeType = "launch"
	AfternoonSnack FoodIntakeType = "afternoon_snack"
	Diner          FoodIntakeType = "dinner"
)

func (value *FoodIntakeType) IsFoodIntakeType() bool {
	foodIntakeTypes := []FoodIntakeType{Breakfast, FirstSnack, Launch, AfternoonSnack, Diner}

	return slices.Contains(foodIntakeTypes, *value)
}

/*
	|-------------------------------|
  |-----------UserEnum------------|
	|-------------------------------|
*/

type UserType string

const (
	User      UserType = "user"
	Admin     UserType = "admin"
	Moderator UserType = "moderator"
)

func (value *UserType) IsUserType() bool {
	userTypes := []UserType{User, Admin}

	return slices.Contains(userTypes, *value)
}

/*
	|-------------------------------|
  |---------WeekdayEnum-----------|
	|-------------------------------|
*/

type WeekDayType string

const (
	Monday    WeekDayType = "mon"
	Tuesday   WeekDayType = "tue"
	Wednesday WeekDayType = "wed"
	Thursday  WeekDayType = "thu"
	Friday    WeekDayType = "fri"
	Saturday  WeekDayType = "sat"
	Sunday    WeekDayType = "sun"
)

func (value *WeekDayType) IsWeekDayType() bool {
	weekDayTypes := []WeekDayType{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}

	return slices.Contains(weekDayTypes, *value)
}

/*
	|-------------------------------|
  |---------TokenTypeEnum---------|
	|-------------------------------|
*/

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

var tokenTypeMap = map[string]TokenType{
	"access_token":  AccessToken,
	"refresh_token": RefreshToken,
}

func (value *TokenType) IsTokenType() bool {
	tokenTypes := []TokenType{AccessToken, RefreshToken}

	return slices.Contains(tokenTypes, *value)
}

func IsAccessToken(tokenType string) bool {
	return tokenTypeMap[tokenType] == AccessToken
}

func IsRefreshToken(tokenType string) bool {
	return tokenTypeMap[tokenType] == RefreshToken
}

/*
	|-------------------------------|
  |------ResponseStatusEnum-------|
	|-------------------------------|
*/

type ResponseStatusType string

const (
	Success ResponseStatusType = "success"
	Error   ResponseStatusType = "error"
)
