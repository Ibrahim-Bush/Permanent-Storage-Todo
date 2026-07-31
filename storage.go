package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"errors"
	"strings"
)

//make global variable for database connection.
var DB *gorm.DB

//define common errors to clarify error testing.
var (
	ErrNotFound = errors.New("Todo not found")
	ErrEmptyTitle = errors.New("Title cannot be empty")
)

type Todo struct {
	Title     string `json:"title" gorm:"not null"`
	ID        int    `json:"id" gorm:"primaryKey"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

func get_db_connection() error {

	//first get connection to the database.
	conn_data := "host=localhost user=postgres password=2006 dbname=todo_app port=5432 sslmode=disable"
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

func create_todo(data input_struct) (*Todo, error){
	//create a new struct with input data.
	var todo Todo
	if data.Title != nil{
		todo.Title = *(data.Title)
	}
	if data.Completed != nil{
		todo.Completed = *(data.Completed)
	}
	
	//validate title of todo.
	if strings.TrimSpace(todo.Title) == ""{
		return nil, ErrEmptyTitle
	}

	//insert the new todo into database.
	result := DB.Create(&todo)
	//check result of insertion.
	if result.Error != nil{
		return nil, result.Error
	}
	//if the creation succeeded the object will be updated to be as inserted.
	return &todo, nil
}

func get_all_todos() (*[]Todo, error){
	//create empty slice to get all records.
	var elements = make([]Todo, 0)
	//get all rows in the table.
	result := DB.Order("id asc").Find(&elements)
	//check result of getting todos.
	if result.Error != nil{
		return nil, result.Error
	}
	//after successfull process return a pointer to todos slice.
	return &elements, nil
}

func get_todo(id int) (*Todo, error){
	//create an object to store the data of required row.
	var todo Todo
	//get the todo with its id from database.
	result := DB.First(&todo, id)
	//check the result of searching.
	if result.Error != nil{
		//differentiate server error from not found element.
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return nil, ErrNotFound
		}
		//if the err is not nil or not found then it is a server error.
		return nil, result.Error
	}
	//if the todo found, it updates object to be as found todo.
	return &todo, nil
}

func update_todo(id int, data input_struct) (*Todo, error){
	//check if there is a new title.
	if data.Title != nil{
		//validate new title
		if strings.TrimSpace(*(data.Title)) == ""{
			return nil, ErrEmptyTitle
		}
	}

	//update target: first we get the target todo.
	todo_ptr, err := get_todo(id)
	if err != nil{
		return nil, err
	}
	//then update the todo.
	result := DB.Model(todo_ptr).Updates(data)
	//check result.
	if result.Error != nil{
		return nil, result.Error
	}
	return todo_ptr, nil
}

func delete_todo(id int) error{
	//delete the target todo with id.
	result := DB.Where("ID = ?", id).Delete(&Todo{})
	//check the result of deletion.
	if result.Error != nil{
		return result.Error
	}else if result.RowsAffected == 0{
		return ErrNotFound
	}
	return nil
}