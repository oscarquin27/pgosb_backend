package main

import (
	"fdms/cmd/api"
	api_models "fdms/cmd/api/models"
	"fdms/src/infrastructure/config"
	"fdms/src/infrastructure/keycloak"
	"fdms/src/infrastructure/postgres"
	"fdms/src/models"
	"fmt"
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

func Test() {

	userJson := &api_models.UserJson{}

	userJson.Code = "1200000"
	userJson.Id = "45"

	var model models.AbstactModel[models.User, api_models.UserJson] = userJson

	fmt.Println(model)
}
