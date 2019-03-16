package app

import (
	"github.com/librarios/librc/minio"
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

func (app *App) Upload(filePath, objectName, bucketName string) {
	if len(bucketName) == 0 {
		bucketName = app.conf.Librarios.Bucket
	}
	if len(objectName) == 0 {
		objectName = filepath.Base(filePath)
	}

	size, err := minio.Upload(&app.conf.Minio, bucketName, filePath, objectName)

	if err != nil {
		log.Fatalf("Failed to upload '%s': %s", filePath, err.Error())
	}
	log.Printf("Uploaded: %s --> %s/%s [%d bytes]\n", filePath, bucketName, objectName, size)
}
