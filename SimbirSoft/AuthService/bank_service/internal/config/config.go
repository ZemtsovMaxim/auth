package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Storage struct {
	Driver string `yaml:"driver" env-required:"true"`
	Info   string `yaml:"info" env-required:"true"`
	URL    string `yaml:"url" env-required:"true`
}

type GRPCConfig struct {
	Network string        `yaml:"network" env-required:"true"`
	Address string        `yaml:"address" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type Config struct {
	MigrationPath string     `yaml:"migration_path" env-required:"true"`
	Storage       Storage    `yaml:"storage" env-required:"true"`
	MyGRPC        GRPCConfig `yaml:"my_grpc" env-required:"true"`
}

func MustLoad() *Config {
	configPath := "config/default.yaml"

	// check file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return &cfg
}
