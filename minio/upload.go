package minio

import (
	"github.com/minio/minio-go"
	"log"
)

// newClient creates a new minio client.
func newClient(conf *Conf) (*minio.Client, error) {
	return minio.New(conf.EndPoint, conf.AccessKey, conf.SecretKey, conf.UseSSL)
}

// createBucket creates a bucket if not exists.
func createBucket(client *minio.Client, bucketName, location string) error {
	exists, err := client.BucketExists(bucketName)
	if err != nil {
		return err
	}

	if !exists {
		return client.MakeBucket(bucketName, location)
	}

	return nil
}

// upload uploads file.
// returns uploaded file size.
func Upload(conf *Conf, bucketName, filePath, objectName string) (int64, error) {
	client, err := newClient(conf)
	if err != nil {
		return -1, err
	}
	log.Printf("Connected to %s\n", conf.EndPoint)

	// create bucket
	if err := createBucket(client, bucketName, conf.Location); err != nil {
		return -1, err
	}
	log.Printf("Ensure bucked created: %s\n", bucketName)

	// detect file content-type
	contentType, err := getFilePathContentType(filePath)
	if err != nil {
		return -1, err
	}

	// upload
	size, err := client.FPutObject(
		bucketName,
		objectName,
		filePath,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return -1, err
	}

	return size, nil
}
