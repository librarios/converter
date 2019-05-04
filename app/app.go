package app

import (
	"log"
	"path/filepath"
)

type App struct {
	conf *Config
}

func NewApp() *App {
	return &App{}
}

func (app *App) Init() {
	// load config
	conf, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}
	app.conf = conf
}

func (app *App) Scan(directory string, opt *ScanOption) {
	scanCmd := NewScanCommand()
	if err := scanCmd.Scan(directory, opt); err != nil {
		log.Fatal(err)
	}
}

func (app *App) Upload(filePath, bucketName, objectName string) (int64, error) {
	if len(objectName) == 0 {
		objectName = filepath.Base(filePath)
	}

	upload := NewUpload(&UploadOption{
		MinioConf:         &app.conf.Minio,
		DefaultBucketName: app.conf.Librarios.Bucket,
	})

	size, err := upload.Upload(filePath, bucketName, objectName)

	if err != nil {
		log.Fatalf("Failed to upload '%s': %s", filePath, err.Error())
	}
	log.Printf("Uploaded: %s --> %s/%s [%d bytes]\n", filePath, bucketName, objectName, size)

	return size, err
}
