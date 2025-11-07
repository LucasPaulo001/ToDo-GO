package routes

import (
	"github.com/LucasPaulo001/ToDo-GO/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine){
	api := r.Group("/api")
	{
		api.GET("/todos", controllers.ListTodos)
		api.GET("/todo/:id", controllers.ListTodo)
		api.POST("/todo", controllers.CreateTodo)
		api.PUT("/todo/:id", controllers.UpdateTodo)
		api.DELETE("/todo/:id", controllers.DeleteTodo)
	}
}