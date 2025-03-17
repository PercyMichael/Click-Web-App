package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a simple route
	r.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{"message": "Hello, World!"})
		c.Data(200, "text/html; charset=utf-8", []byte("<h1>Hello, World!</h1>"))
	})

	
	// Start the server on port 8080
	r.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
