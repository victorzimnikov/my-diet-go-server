package middleware

import (
	"errors"
	"fmt"
	"my-diet-server/config"
	"my-diet-server/database"
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	tokenString := utils.GetToken(c)

	if tokenString == "" {
		err := errors.New("not login").Error()

		response.Status = "error"
		response.Error = err
		response.Message = "You are not logged in"

		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	jwtSecret := config.Config("JWT_ACCESS_SECRET")

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (any, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Invalid token"

		return c.Status(fiber.StatusUnauthorized).JSON(response)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)

	if !ok || !tokenByte.Valid || !models.IsAccessToken(fmt.Sprint(claims["type"])) {
		response.Status = "error"
		response.Error = err
		response.Message = "Invalid token"

		return c.Status(fiber.StatusUnauthorized).JSON(response)

	}

	var user models.UserModel

	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID.String() != claims["sub"] {
		err := errors.New("not login").Error()

		response.Status = "error"
		response.Error = err
		response.Message = "The user belonging to this token no logger exists"

		return c.Status(fiber.StatusForbidden).JSON(response)
	}

	c.Locals("user", user.GetUserResponse())

	return c.Next()
}
