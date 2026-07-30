package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//make global variable for database connection.
var DB *gorm.DB

type Todo struct {
	//we use pointers to differentiate default zero-valued from user input.
	Title     *string `json:"title" gorm:"not null"`
	ID        *int    `json:"id" gorm:"primaryKey"`
	Completed *bool   `json:"completed" gorm:"default:false"`
}

func get_db_connection() error {

	//first get connection to the database.
	conn_data := "host:localhost user:postgres password:2006 dbname:todo_app port:5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn_data), &gorm.Config{})
	//check if something went wrong.
	if err != nil{
		return err
	}
	//create the table with AutoMigrate.
	err = db.AutoMigrate(&Todo{})
	if err != nil{
		return err
	}
	//assign the database connection to the global variable.
	DB = db
	return nil
}
