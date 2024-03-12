package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/muaramasad/ubersnap-challenge/handler"
	"github.com/stretchr/testify/assert"
)

func createMultipartFormDataResize(t *testing.T, fieldName, fileName string, width, height int) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)

	// Add image file part
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}

	// Add width field
	if fw, err = w.CreateFormField("width"); err != nil {
		t.Errorf("Error creating width field: %v", err)
	}
	if _, err = io.Copy(fw, bytes.NewBufferString(fmt.Sprintf("%d", width))); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}

	// Add height field
	if fw, err = w.CreateFormField("height"); err != nil {
		t.Errorf("Error creating height field: %v", err)
	}
	if _, err = io.Copy(fw, bytes.NewBufferString(fmt.Sprintf("%d", height))); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}

	w.Close()
	return b, w
}

func TestResizeRoute(t *testing.T) {
	server := fiber.New()
	server.Post("/api/v1/resize", handler.Resize)
	b, w := createMultipartFormDataResize(t, "image", "./image_test/image_test.png", 20, 10)

	req, err := http.NewRequest("POST", "/api/v1/resize", &b)
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
