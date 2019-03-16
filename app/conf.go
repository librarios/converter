package app

import (
	"fmt"
	"github.com/lechuckroh/appconfig"
	"github.com/librarios/librc/minio"
	"os"
)

type Librarios struct {
	Bucket string
}

type Config struct {
	Minio     minio.Conf
	Librarios Librarios
}

func defaultConfig() *Config {
	return &Config{
		Minio: minio.Conf{
			EndPoint:  "localhost:9000",
			AccessKey: "",
			SecretKey: "",
			Location:  "us-east-1",
			UseSSL:    false,
		},
		Librarios: Librarios{
			Bucket: "",
		},
	}
}

func loadConfig() (*Config, error) {
	config := defaultConfig()

	appconfig.ActiveProfileEnvName = "librc.profiles.active"
	appconfig.ConfigFilenamePrefix = "librc"
	configFilename := os.Getenv("LIBRC_CONFIG_FILE")
	loadFilenames, err := appconfig.LoadConfig(configFilename, config)
	for idx, filename := range loadFilenames {
		fmt.Printf("Configuration loaded: [%d] %s\n", idx+1, filename)
	}
	if len(loadFilenames) == 0 {
		if len(configFilename) > 0 {
			fmt.Printf("Config file not found: %s", configFilename)
		}

		fmt.Println("Default configuration loaded.")
	}
	return config, err
}
