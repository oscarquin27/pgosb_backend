package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"

	"gopkg.in/yaml.v3"
)

var threadProfile = pprof.Lookup("threadcreate")
var goRoutineProfile = pprof.Lookup("goroutine")

var Configuration *Config

func Get() *Config {
	return Configuration
}

func GetNumbersOfThreads() int {
	return threadProfile.Count()
}
func GetNumberOfGoRoutines() int {
	return goRoutineProfile.Count()
}

func LoadConfig() *Config {

	path, err := os.Getwd()

	path = filepath.Join(path, "config", "settings.yml")

	file, err := os.Open(path)

	if err != nil {
		panic(fmt.Sprintf("no se pudo leer el archivo de configuracion %v", err))
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&Configuration); err != nil {
		panic(fmt.Sprintf("no se pudo parsear el archivo de configuracion %v", err))
	}

	return Configuration
}

func init() {
	fmt.Println("Inicio Config Package")
	LoadConfig()

}
