package router

import (
	"Permanent-Storage-Todo/handler"
	"github.com/gin-gonic/gin"
)

func Init_router(handler *handler.Todo_handler) (*gin.Engine){
	//init router with default middleware.
	router := gin.Default()

	//register routes.
	router.GET("/todos", handler.Get_all_todos_handler)
	router.GET("/todos/search", handler.Get_todos_by_title_handler)
	router.GET("/todos/category/:category", handler.Get_todos_by_category_handler)
	router.GET("/todos/status/:status", handler.Get_todos_by_status_handler)
	router.GET("/todos/:id", handler.Get_todo_by_id_handler)
	router.POST("/todos", handler.Create_todo_handler)
	router.PUT("/todos/:id", handler.Update_todo_by_id_handler)
	router.PUT("/todos/category/:category", handler.Update_todos_by_category_handler)
	router.DELETE("/todos/:id", handler.Delete_todo_by_id_handler)
	router.DELETE("todos", handler.Delete_all_todos_handler)

	//return the address of the router.
	return router
}