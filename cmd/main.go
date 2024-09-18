package main

import (
	"fdms/cmd/api"
	"fdms/src/infrastructure/config"
	"fdms/src/infrastructure/keycloak"
	"fdms/src/infrastructure/postgres"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

func main() {

	pool, err := postgres.CreatePool()

	keycloakGoClient := gocloak.NewClient(config.Get().Keycloak.Address)

	keycloakService := keycloak.NewService(keycloakGoClient)

	if err != nil {
		log.Fatal("No hay conexion a la bd")
	}

	api.Run(pool, keycloakService)

}
