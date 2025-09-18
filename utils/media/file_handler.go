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

func UploadFileToMinio(fileForm *multipart.FileHeader, pathPrefix string) (string, error) {
	fileReader, err := fileForm.Open()
	if err != nil {
		return "", err
	}
	defer fileReader.Close()
	
	objectName, fileSize, contentType, err := ExtractFileData(fileForm)
	if err != nil {
		return "", err
	}

	objectName = pathPrefix + objectName
	_, err = SendFileToMinio(objectName, fileReader, fileSize, contentType)
	if err != nil {
		return "", err
	}

	return objectName, nil
}

func UpdateFileInMinio(oldObjectName string, newFileForm *multipart.FileHeader, pathPrefix string) (string, error) {
	if oldObjectName != "" {
		deleteFile(oldObjectName)
	}

	return UploadFileToMinio(newFileForm, pathPrefix)
}

func deleteFile(objectName string) error {
	err := DeleteFileFromMinio(objectName)
	if err != nil {
		return err
	}

	return nil
}
