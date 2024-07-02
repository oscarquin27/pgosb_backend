package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		ConnectionString string `yaml:"connection_string"`
	}
	Keycloak struct {
		ClientId string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
	}
	Http struct {
		Port string `yaml:"port"`
	}
}

var Configuration *Config

func NewConfig() (*Config, error) {
	
	path, err := os.Getwd()


	path = filepath.Join(path, "bin", "config", "settings.yml")


	file, err := os.Open(path)

	if err != nil {
		log.Fatal("No se encontro el archivo de configuracion")
		os.Exit(1)
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&Configuration); err != nil {
		return nil, err
	}

	return Configuration, nil
}