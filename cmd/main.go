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

// func test() {

// 	userJson := &api_models.UserJson{}

// 	userJson.Code = "1200000"
// 	userJson.Id = "45"

// 	var model Abstra[models.User, api_models.UserJson] = userJson

// 	message, err := json.Marshal(userJson)

// 	if err != nil {
// 		panic(err)
// 	}

// 	message2 := model.ToModel()

// 	fmt.Println(string(message))

// 	json2, _ := json.Marshal(message2)

// 	fmt.Println(model)
// 	fmt.Println("***---**--**-*-****-----****------*****-----*****")

// 	fmt.Println(string(json2))
// }
