package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	AppEnv             string             `yaml:"env" env-default:"local"` // Привязка к "env"
	DatabaseConnection DatabaseConnection `yaml:"database"`                // Остаётся без изменений
	HTTPConfig         HTTPServer         `yaml:"http_server"`             // Остаётся без изменений
	MigrationsPath     string             `yaml:"migrations_path" env-default:"../../migrations"`
}

type DatabaseConnection struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"secret"`
	Name     string `yaml:"db" env-default:"songs"`
}

type HTTPServer struct {
	Host               string        `yaml:"host" env-default:"0.0.0.0"`
	Port               string        `yaml:"port" env-default:"8080"`
	ReadTimeout        time.Duration `yaml:"read_timeout" env-default:"5s"`
	WriteTimeout       time.Duration `yaml:"write_timeout" env-default:"10s"`
	MaxHeaderMegabytes int           `yaml:"max_header_megabytes" env-default:"1"`
}

func MustLoad() *Config {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	exeDir := filepath.Dir(exePath)
	configABSPath := filepath.Join(exeDir, "../../config/config.yaml")

	if configABSPath == "" {
		log.Fatal("CONFIG_PATH not set")
	}
	if _, err := os.Stat(configABSPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist\n", configABSPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configABSPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", configABSPath)
	}
	cfg.MigrationsPath = filepath.Join(exeDir, cfg.MigrationsPath)
	return &cfg
}
