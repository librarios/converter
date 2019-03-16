package minio

import (
	"net/http"
	"os"
)

// getFileContentType detects file content type.
// Returns a valid content-type by returning "application/octet-stream" if no others seemed to match.
func getFileContentType(file *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}

func getFilePathContentType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return getFileContentType(file)
}
