package test

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/muaramasad/ubersnap-challenge/handler"
	"github.com/stretchr/testify/assert"
)

func TestViewRoute(t *testing.T) {
	imageFormatExpected := "image/png"
	server := fiber.New()
	server.Get("/api/v1/view/:filename", handler.View)

	req, _ := http.NewRequest("GET", "/api/v1/view/7f6bd20ebb8d4601b22b6d74ad38e0c3_compressed.png", nil)
	resp, _ := server.Test(req, -1)

	assert.Equal(t, imageFormatExpected, resp.Header.Get("Content-Type"))
}
