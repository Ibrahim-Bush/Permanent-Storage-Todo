package main

import "log"

func main(){

	//get a connection to database.
	err := get_db_connection()
	//if something went wrong exit.
	if err != nil{
		log.Fatal(err)
	}
}