package main

import (
	"fdms/infra/database"
	"fdms/server"
	"log"
)

func main(){
    pool, err := database.NewDatabase()

	if err != nil {
		log.Fatal("No hay conexion a la bd")
	}
		
	server.Run(pool)
}