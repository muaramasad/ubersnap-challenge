package service

import (
	"fmt"
	"io/ioutil"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// Process Converting Image
func ProcessView(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := filepath.Join("./images/", filename)
	data, err := ioutil.ReadFile(filePath)
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType != "" {
		c.Set("Content-Type", contentType)
	} else {
		c.Set("Content-Type", "application/octet-stream")
	}
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(fmt.Sprintf("File not found sss: %s", err.Error()))
	}
	return c.Status(200).Send(data)
}
