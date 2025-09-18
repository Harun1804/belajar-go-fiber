package media

import (
	"belajar-go-fiber/configs"
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var MINIO_BUCKET_NAME string = configs.GetEnv("MINIO_BUCKET_NAME", "golang-fiber")

func InitMinio() error {
	endpointUrl := configs.GetEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKeyID := configs.GetEnv("MINIO_ACCESS_KEY", "minioadmin")
	secretAccessKey := configs.GetEnv("MINIO_SECRET_KEY", "minioadmin")
	useSSL := configs.GetEnv("MINIO_USE_SSL", "false") == "true"

	client, err := minio.New(endpointUrl, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return err
	}

	MinioClient = client
	return nil
}

// GetFileFromMinio downloads an object from MinIO and returns its reader
func GetFileFromMinio(objectName string) (io.ReadCloser, error) {
	ctx := context.Background()
	obj, err := MinioClient.GetObject(ctx, MINIO_BUCKET_NAME, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	// Optionally check if object exists by reading a byte
	_, err = obj.Stat()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// GeneratePresignedURL generates a presigned URL for accessing an object
func GeneratePresignedURL(objectName string, expirySeconds int64) (string, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	presignedURL, err := MinioClient.PresignedGetObject(ctx, MINIO_BUCKET_NAME, objectName, time.Duration(expirySeconds)*time.Second, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func CreateBucket() error {
	ctx := context.Background()
	exists, err := MinioClient.BucketExists(ctx, MINIO_BUCKET_NAME)
	if err != nil {
		return err
	}

	if !exists {
		err = MinioClient.MakeBucket(ctx, MINIO_BUCKET_NAME, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}

func SendFileToMinio(objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	err := CreateBucket()
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	_, err = MinioClient.PutObject(ctx, MINIO_BUCKET_NAME, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", MinioClient.EndpointURL(), MINIO_BUCKET_NAME, objectName)
	return url, nil
}

func DeleteFileFromMinio(objectName string) error {
	err := CreateBucket()
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = MinioClient.RemoveObject(ctx, MINIO_BUCKET_NAME, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
