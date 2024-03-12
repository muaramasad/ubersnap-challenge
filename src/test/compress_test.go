package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/muaramasad/ubersnap-challenge/handler"
	"github.com/stretchr/testify/assert"
)

func TestCompressRoute(t *testing.T) {
	server := fiber.New()
	server.Post("/api/v1/compress", handler.Compress)
	b, w := createMultipartFormData(t, "image", "./image_test/image_test.png")

	req, err := http.NewRequest("POST", "/api/v1/compress", &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	resp, _ := server.Test(req, -1)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(resp_body))
		return
	}

	// fmt.Println(resp_body)
	// fmt.Println(string(resp_body))

	assert.Equal(t, 200, resp.StatusCode)
}
