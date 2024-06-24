package database

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"fdms/infra/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

//var Pool *pgxpool.Pool
var pgxConfig *pgxpool.Config

func NewDatabase() (*pgxpool.Pool, error) {
	var err error

	path, err := os.Getwd()

	path = filepath.Join(path, "config", "settings.yml")
	
	config, err := config.NewConfig(path)

	if err != nil {
		log.Fatal("No se pudo obtener archivo de configuracion, verificar filepath")
		os.Exit(1)
	}

	pgxConfig, err = pgxpool.ParseConfig(config.Database.ConnectionString)

	if err != nil {
		log.Fatal("No se pudo parsear el archivo de configuracion, verificar el yaml")
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)

	if err != nil {
		log.Fatal("No se pudo generar el pool de conexiones con la configuracion actual, verificar")
		os.Exit(1)
	}


	return pool, nil
}