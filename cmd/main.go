package main

import (
	"fdms/infra/config"
	"fdms/infra/database"
	"fdms/server"
	"log"
)

func main(){

	config.NewConfig()
    pool, err := database.NewDatabase()

	if err != nil {
		log.Fatal("No hay conexion a la bd")
	}

	server.Run(pool)

}