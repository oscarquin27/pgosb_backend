package postgres

import (
	"context"
	"fdms/src/infrastructure/config"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool() (*pgxpool.Pool, error) {

	pgxConfig, err := pgxpool.ParseConfig(config.Configuration.Database.ConnectionString)

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