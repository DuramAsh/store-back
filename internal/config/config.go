package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultAppMode    = "dev"
	defaultAppPort    = "2004"
	defaultAppHost    = "http://localhost:2004"
	defaultAppPath    = "/"
	defaultAppTimeout = 60 * time.Second
)

type (
	Configs struct {
		APP      AppConfig
		POSTGRES StoreConfig
	}

	AppConfig struct {
		Mode    string `required:"true"`
		Port    string
		Host    string
		Path    string
		Timeout time.Duration
	}

	ClientConfig struct {
		URL      string
		Gateway  string
		Login    string
		Password string
	}

	StoreConfig struct {
		DSN string
	}
)

// New populates Configs struct with values from config file
// located at filepath and environment variables.
func New() (cfg Configs, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}

	if err = godotenv.Load(filepath.Join(root, ".env")); err != nil {
		return
	}

	cfg.APP = AppConfig{
		Mode:    defaultAppMode,
		Port:    defaultAppPort,
		Host:    defaultAppHost,
		Path:    defaultAppPath,
		Timeout: defaultAppTimeout,
	}

	if err = envconfig.Process("POSTGRES", &cfg.POSTGRES); err != nil {
		return
	}

	return
}
