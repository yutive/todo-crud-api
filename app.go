package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yutive/todo-crud-api/controllers"
)

func main() {
	//starter text
	fmt.Println("todo crud api")

	//start gin server
	r := gin.Default()

	//Home Handler
	r.GET("/", controllers.HomeHandler)

	// Versioning of API
	v1 := r.Group("/api/v1")
	{
		//get Todos
		v1.GET("/todos", controllers.GetTodos)
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
