package handler
import (
 "github.com/gofiber/fiber/v2"
 "github.com/muaramasad/ubersnap-challenge/service"
)

// Convert Image
func Convert(c *fiber.Ctx) error {
	return service.ProcessConvert(c)
}

// Resize Image
func Resize(c *fiber.Ctx) error {
	return service.ProcessResize(c)
}

// Compress Image
func Compress(c *fiber.Ctx) error {
	return service.ProcessCompress(c)
}

// View Image
func View(c *fiber.Ctx) error {
	return service.ProcessView(c)
}