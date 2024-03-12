package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestResizeEndpoint(t *testing.T) {

	// Create a multipart form body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a file part
	filePart, err := writer.CreateFormFile("file", "./image_test/image_test.png")
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}

	// Open and copy file content into file part
	file, err := os.Open("./image_test/image_test.png")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(filePart, file)
	if err != nil {
		t.Fatalf("failed to copy file content: %v", err)
	}

	// Add width and height parameters
	err = writer.WriteField("width", "100")
	if err != nil {
		t.Fatalf("failed to write width field: %v", err)
	}

	err = writer.WriteField("height", "100")
	if err != nil {
		t.Fatalf("failed to write height field: %v", err)
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		t.Fatalf("failed to close multipart writer: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "/api/v1/resize", body)
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a recorder to record the response
	rr := httptest.NewRecorder()

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Image resized successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
