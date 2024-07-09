package main

import (
	"context"
	"fdms/infra/config"
	"fdms/infra/database"
	authentication "fdms/infra/keycloak"
	"fdms/server"
	"fmt"
	"log"

	"flag"

	"github.com/Nerzal/gocloak/v13"
)

var usernamePtr = flag.String("username", "orlando", "Username for Keycloak")
var passwordPtr = flag.String("password", "12345", "Password for Keycloak")
var emailPtr = flag.String("email", "", "Email for Keycloak")

var command = flag.String("command", "loging", "Method to test")

func main() {

	pool, err := database.NewDatabase()

	if err != nil {
		log.Fatal("No hay conexion a la bd")
	}

	server.Run(pool)

}

func LoginUser() *gocloak.JWT {
	fmt.Println("Inicio Login")

	keyClocakClient := gocloak.NewClient(config.Configuration.Keycloak.Address)

	keycloakService := authentication.NewService(keyClocakClient)

	ctx := context.Background()

	jwt, err := keycloakService.LoginUser(ctx, *usernamePtr, *passwordPtr)

	if err != nil {
		fmt.Println("SE CAGO", err)
		return nil
	}

	fmt.Println("EXITO", jwt)
	return jwt
}

func CreateUser() {
	fmt.Println("Inicio Create")

	keyClocakClient := gocloak.NewClient(config.Configuration.Keycloak.Address)

	keycloakService := authentication.NewService(keyClocakClient)

	ctx := context.Background()

	_, err := keycloakService.CreateUser(ctx, *usernamePtr, *emailPtr, "USERID125544MMFD", *passwordPtr)

	if err != nil {
		fmt.Println("SE CAGO", err)
		return
	}
	fmt.Println("EXITO")
}
