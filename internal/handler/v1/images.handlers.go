package handlersV1

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"my-diet-server/database"
	"my-diet-server/internal/models"
	"my-diet-server/internal/utils"
	"my-diet-server/storage"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var img models.ImageModel

	img.Url = "placeholder"

	database.DB.Create(&img)

	fileName := fmt.Sprintf("%s%s", img.ID, filepath.Ext(file.Filename))

	dir, _ := utils.GetCurrentDir()

	reader, err2 := os.Open(filepath.Join(dir, "storage", "images", "733ce04c-a436-4961-98cd-4a6775196d7a.jpg"))

	if err2 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err2.Error())
	}

	defer reader.Close()

	im, _, err3 := image.DecodeConfig(reader)

	if err3 != nil {
		return c.Status(fiber.StatusAccepted).JSON(err3.Error())
	}

	img.Width = im.Width
	img.Height = im.Height

	err = storage.LocalStorage.PutFileAs("images", file, fileName)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	img.Url = storage.LocalStorage.URL(fmt.Sprintf("images/%s", fileName))

	database.DB.Save(&img)

	return c.JSON(img.ToResponse())
}
