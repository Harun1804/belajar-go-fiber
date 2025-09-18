package media

import (
	"crypto/rand"
	"encoding/hex"
	"mime/multipart"
	"path/filepath"
)

// generateRandomFilename generates a random filename with the same extension as the original
func GenerateRandomFilename(originalName string) string {
	ext := filepath.Ext(originalName)
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return originalName // fallback to original name
	}
	return hex.EncodeToString(b) + ext
}

// ExtractFileData returns objectName, fileSize, contentType from *multipart.FileHeader
func ExtractFileData(fileHeader *multipart.FileHeader) (objectName string, fileSize int64, contentType string, err error) {
	objectName = GenerateRandomFilename(fileHeader.Filename)
	fileSize = fileHeader.Size
	contentType = fileHeader.Header.Get("Content-Type")
	return objectName, fileSize, contentType, nil
}
