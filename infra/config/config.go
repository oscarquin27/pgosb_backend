package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		ConnectionString string `yaml:"connection_string"`
	}
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	
	file, err := os.Open(configPath)
	
	if err != nil {
		log.Fatal("No se encontro el archivo de configuracion")
		os.Exit(1)
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}