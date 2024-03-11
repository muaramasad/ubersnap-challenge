package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Process Converting Image
func ProcessConvert(c *fiber.Ctx) error {
	newExt := "jpeg"

	port := os.Getenv("APP_PORT")

	// parse incomming image file

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	outputImage := fmt.Sprintf("%s.%s", filename, newExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	inputFilePath := fmt.Sprintf("./images/%s", image)
	outputFilePath := fmt.Sprintf("./images/%s", outputImage)

	// Run FFmpeg command
	cmd := exec.Command("ffmpeg", "-i", inputFilePath, outputFilePath)
	err = cmd.Run()
	if err != nil {
		log.Println("Error converting:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// generate image url to serve to client using CDN
	newImageUrl := fmt.Sprintf("http://localhost:"+port+"/api/v1/view/%s", outputImage)

	// create meta data and send to client

	data := map[string]interface{}{
		"imageUrl": newImageUrl,
	}

	return c.JSON(fiber.Map{"status": 200, "message": "Image converted successfully", "data": data})
}
