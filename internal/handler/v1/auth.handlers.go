package handlersV1

import (
	"errors"
	"fmt"
	"my-diet-server/config"
	"my-diet-server/database"
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	var payload *models.RegisterUserDto

	if err := c.BodyParser(&payload); err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errs := payload.Validate()

	if errs != nil {
		response.Status = "error"
		response.Error = errs
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if payload.Password != payload.PasswordConfirm {
		response.Status = "error"
		response.Error = errors.New("passwords do not match").Error()
		response.Message = "Пароли не совпадают"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newUser := new(models.UserModel)

	newUser.CreateUser(hashedPassword, *payload)

	result := database.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		response.Status = "error"
		response.Error = result.Error.Error()
		response.Message = "User with that email already exists"

		return c.Status(fiber.StatusConflict).JSON(response)
	} else if result.Error != nil {
		response.Status = "error"
		response.Error = result.Error.Error()
		response.Message = "Something bad happened"

		return c.Status(fiber.StatusBadGateway).JSON(response)
	}

	response.Status = "success"
	response.Data = newUser.GetUserResponse()

	return c.Status(fiber.StatusCreated).JSON(response)
}

type JwtAccessData struct {
	JwtMaxAge    int
	JwtSecret    string
	JwtExpiredIn time.Duration
}

type JwtRefreshData struct {
	JwtMaxAge    int
	JwtSecret    string
	JwtExpiredIn time.Duration
}

type JwtData struct {
	Access  JwtAccessData
	Refresh JwtRefreshData
}

func getJwtAccessData() (JwtAccessData, error) {
	var data JwtAccessData

	jwtExpiredIn, err := time.ParseDuration(config.Config("JWT_ACCESS_EXPIRED_IN"))

	if err != nil {
		return JwtAccessData{}, err
	}

	jwtMaxAge, err := strconv.Atoi(config.Config("JWT_ACCESS_MAXAGE"))

	jwtSecret := config.Config("JWT_ACCESS_SECRET")

	if jwtSecret == "" {
		return JwtAccessData{}, errors.New("invalid jwt secret")
	}

	if err != nil {
		return JwtAccessData{}, err
	}

	data.JwtExpiredIn = jwtExpiredIn
	data.JwtMaxAge = jwtMaxAge
	data.JwtSecret = jwtSecret

	return data, nil
}

func getJwtRefreshData() (JwtRefreshData, error) {
	var data JwtRefreshData

	jwtExpiredIn, err := time.ParseDuration(config.Config("JWT_REFRESH_EXPIRED_IN"))

	if err != nil {
		return JwtRefreshData{}, err
	}

	jwtMaxAge, err := strconv.Atoi(config.Config("JWT_REFRESH_MAXAGE"))

	jwtSecret := config.Config("JWT_REFRESH_SECRET")

	if jwtSecret == "" {
		return JwtRefreshData{}, errors.New("invalid jwt secret")
	}

	if err != nil {
		return JwtRefreshData{}, err
	}

	data.JwtExpiredIn = jwtExpiredIn
	data.JwtMaxAge = jwtMaxAge
	data.JwtSecret = jwtSecret

	return data, nil
}

func getJwtData() (JwtData, error) {
	var data JwtData

	jwtAccessData, accessErr := getJwtAccessData()
	jwtRefreshData, refreshErr := getJwtRefreshData()

	if accessErr != nil {
		return JwtData{}, accessErr
	}

	if refreshErr != nil {
		return JwtData{}, refreshErr
	}

	data.Access = jwtAccessData
	data.Refresh = jwtRefreshData

	return data, nil
}

func getTokenString(userId uuid.UUID, jwtSecret string, jwtExpiredIn time.Duration, tokenType models.TokenType) (string, error) {
	tokenByte := jwt.New(jwt.SigningMethodHS256)

	claims := tokenByte.Claims.(jwt.MapClaims)

	now := time.Now().UTC()

	claims["type"] = tokenType
	claims["sub"] = userId
	claims["exp"] = now.Add(jwtExpiredIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	return tokenByte.SignedString([]byte(jwtSecret))
}

func getTokensResponse(c *fiber.Ctx, userId uuid.UUID, response *models.NewAppResponse[any]) error {
	jwtData, err := getJwtData()

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "JwtData failed"

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	accessTokenString, accessTokenErr := getTokenString(userId, jwtData.Access.JwtSecret, jwtData.Access.JwtExpiredIn, models.AccessToken)
	refreshTokenString, refreshTokenErr := getTokenString(userId, jwtData.Refresh.JwtSecret, jwtData.Refresh.JwtExpiredIn, models.RefreshToken)

	if accessTokenErr != nil {
		response.Status = "error"
		response.Error = accessTokenErr.Error()
		response.Message = "Generating JWT Token failed"

		return c.Status(fiber.StatusBadGateway).JSON(response)
	}

	if refreshTokenErr != nil {
		response.Status = "error"
		response.Error = refreshTokenErr.Error()
		response.Message = "Generating JWT Token failed"

		return c.Status(fiber.StatusBadGateway).JSON(response)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    accessTokenString,
		Path:     "/",
		MaxAge:   jwtData.Access.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	response.Status = "success"
	response.Data = fiber.Map{"accessToken": accessTokenString, "refreshToken": refreshTokenString}

	return c.Status(fiber.StatusOK).JSON(response)
}

func Login(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	var payload *models.LoginUserDto

	if err := c.BodyParser(&payload); err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errs := payload.Validate()

	if errs != nil {
		response.Status = "error"
		response.Error = errs
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var user models.UserModel

	result := database.DB.First(&user, "email = ?", strings.ToLower(payload.Email))

	if result.Error != nil {
		response.Status = "error"
		response.Error = result.Error.Error()
		response.Message = "Invalid email or Password"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Invalid email or Password"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return getTokensResponse(c, user.ID, response)
}

func LogoutUser(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()

	response.Status = "success"

	return c.Status(fiber.StatusOK).JSON(response)
}

func RefreshAccessToken(c *fiber.Ctx) error {
	response := utils.MakeAppResponse()
	var payload *models.RefreshTokenDto

	if err := c.BodyParser(&payload); err != nil {
		response.Status = "error"
		response.Error = err.Error()
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	errs := payload.Validate()

	if errs != nil {
		response.Status = "error"
		response.Error = errs
		response.Message = "Bad request"

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	jwtSecret := config.Config("JWT_REFRESH_SECRET")

	tokenByte, err := jwt.Parse(payload.RefreshToken, func(jwtToken *jwt.Token) (any, error) {
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

	if !ok || !tokenByte.Valid {
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

	return getTokensResponse(c, user.ID, response)
}
