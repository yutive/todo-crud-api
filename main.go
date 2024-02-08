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

	// Versioning of API
	v1 := r.Group("/api/v1")
	{
		//get Todos
		v1.GET("/todos", GetTodos)
	}

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		// return json if route is not defined
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// listen and serve on localhost:8080
	err := r.Run()
	if err != nil {
		return
	}
}
