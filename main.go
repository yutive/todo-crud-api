package main

import (
	"fmt"
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

func main() {
	//starter text
	fmt.Println("todo crud api")

	//start gin server
	r := gin.Default()

	//	db
	todo = append(todo, models.Todo{ID: "0", Text: "wash dishes"})
	todo = append(todo, models.Todo{ID: "1", Text: "watch football"})

	//Home Handler
	r.GET("/", HomeHandler)

	//get Todos
	r.GET("/todos", GetTodos)

	// listen and serve on localhost:8080
	r.Run()
}
