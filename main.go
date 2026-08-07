package main

import (
	"Permanent-Storage-Todo/config"
	"Permanent-Storage-Todo/repository"
	"Permanent-Storage-Todo/service"
	"Permanent-Storage-Todo/handler"
	"Permanent-Storage-Todo/router"
	"log"
	)

func main(){

	//first: get a connection to database.
	connection_data := "host=localhost user=postgres password=admin dbname=todo_app port=5432 sslmode=disable"
	database, err := database.Init_database(connection_data)
	//if something went wrong exit.
	if err != nil{
		log.Fatal(err)
	}

	//linking database to repository layer.
	repository := repository.Init_postgres_repo(database)

	//linking repository layer with service layer.
	service := service.Init_service(repository)

	//linking service layer to handler layer.
	handler := handler.Init_handler(service)

	//linking the handler layer to the router.
	router := router.Init_router(handler)

	//run the server at local port ":8080".
	router.Run(":8080")

}