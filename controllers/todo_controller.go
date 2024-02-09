package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yutive/todo-crud-api/models"
	"net/http"
)

var todo []models.Todo

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "todo crud api",
	})
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo)
}

func NewTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todo = append(todo, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
