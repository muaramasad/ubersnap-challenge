package router
import (
 "github.com/muaramasad/ubersnap-challenge/handler"
 "github.com/gofiber/fiber/v2"
)
// SetupRoutes func
func SetupRoutes(app *fiber.App) {
 // grouping
 api := app.Group("/api")
 v1 := api.Group("/v1")
 // routes
 v1.Post("/convert", handler.Convert)
 v1.Post("/resize", handler.Resize)
 v1.Post("/compress", handler.Compress)
 v1.Get("/view/:filename", handler.View)
}