package main

import (
	"strconv"
	"errors"

	"github.com/gin-gonic/gin"
	)

type input_struct struct{
	//we use pointers to differentiate default zero-valued from user input.
	Title	  *string	`json:"title"`
	Completed *bool		`json:"completed"`
}

func create_todo_handler(c *gin.Context){
	//get the data from json body.
	var data input_struct
	err := c.ShouldBindJSON(&data)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid json"})
		return
	}
	//create new todo with the input data.
	todo_ptr, err := create_todo(data)
	//check if something went wrong.
	switch{
	case err == nil:
		c.JSON(200, *todo_ptr)
	case errors.Is(err, ErrEmptyTitle):
		c.JSON(400, gin.H{"error": ErrEmptyTitle})
	default:
		c.JSON(500, gin.H{"error": "Error with creating todo"})
	}
}

func get_all_todos_handler(c *gin.Context){
	//call get all function.
	todos_ptr, err := get_all_todos()
	if err != nil{
		c.JSON(500, gin.H{"error": "Error with fetching items"})
		return
	}
	//if the process succeeded, display all items.
	c.JSON(200, *todos_ptr)
}

func get_todo_handler(c *gin.Context){
	//get id parameter and validate it.
	input_id := c.Param("id")
	//convert id to integer.
	id, err := strconv.Atoi(input_id)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid todo id"})
		return
	}
	//if id converted successfully call get function.
	todo_ptr, err := get_todo(id)
	//check if something went wrong.
	switch{
	case err == nil:
		c.JSON(200, *todo_ptr)
	case errors.Is(err, ErrNotFound):
		c.JSON(404, gin.H{"error": ErrNotFound})
	default:
		c.JSON(500, gin.H{"error": "Error in fetching item"})				
	}
}

func update_todo_handler(c *gin.Context){
	//first take the id param.
	input_id := c.Param("id")
	//convert it to integer.
	id, err := strconv.Atoi(input_id)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid todo id"})
		return
	}

	//then take the json body with struct.
	var data input_struct
	err = c.ShouldBindJSON(&data)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid json"})
		return
	}

	//finally, update required item with update function.
	todo_ptr, err := update_todo(id, data)
	//check if something went wrong.
	switch{
	case err == nil:
		c.JSON(200, *todo_ptr)
	case errors.Is(err, ErrEmptyTitle):
		c.JSON(400, gin.H{"error": ErrEmptyTitle})
	case errors.Is(err, ErrNotFound):
		c.JSON(404, gin.H{"error": ErrNotFound})
	default:
		c.JSON(500, gin.H{"error": "Error with updating todo"})
	}
}

func delete_todo_handler(c *gin.Context){
	//get id parameter and validate it.
	input_id := c.Param("id")
	//convert id to integer.
	id, err := strconv.Atoi(input_id)
	if err != nil{
		c.JSON(400, gin.H{"error": "Invalid todo id"})
		return
	}
	//delete todo by id.
	err = delete_todo(id)
	//check if something went wrong.
	switch{
	case err == nil:
		c.JSON(200, gin.H{"message": "Todo deleted successfully"})
	case errors.Is(err, ErrNotFound):
		c.JSON(404, gin.H{"error": "Todo not found"})
	default:
		c.JSON(500, gin.H{"error": "Error with deleting todo"})
	}
}