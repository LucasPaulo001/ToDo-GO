package controllers

import (
	"net/http"
	"strconv"

	"github.com/LucasPaulo001/ToDo-GO/src/model"
	"github.com/gin-gonic/gin"
)

var todos = []model.Todo {
	{ID: 1, Title: "Estudar GO", Done: false},
	{ID: 2, Title: "Estudar Docker", Done: false},
}

var newIndex = 3

// Listar tasks
func ListTodos(c *gin.Context){
	c.JSON(http.StatusOK, todos)
}

// Listar uma task
func ListTodo(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido."})
		return
	}

	for _, t := range todos {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada."})
}

// Criando task
func CreateTodo(c *gin.Context){
	var newTodo model.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido."})
		return
	}

	newTodo.ID = newIndex
	newIndex++
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// Atualizando uma task
func UpdateTodo(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido."})
		return
	}

	var updated model.Todo
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payload inválido"})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			updated.ID = id
			todos[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada."})
}

// Deletar uma task
func DeleteTodo(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id inválido."})
		return
	}

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada."})
}


