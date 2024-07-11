package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var Configuration *Config

func Get() *Config {
	return Configuration
}

func LoadConfig() (*Config, error) {

	//path, err := os.Getwd()

	path := filepath.Join("./config", "settings.yml")

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

func init() {
	fmt.Println("Inicio Config Package")
	_, err := LoadConfig()
	if err != nil {
		panic(err)
	}
}
