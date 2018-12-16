package main

import "github.com/gin-gonic/gin"

func main() {
	println("Start of catbook")
	router := gin.Default()

	router.GET("/", MainHandler)

	// Listen and Serve on localhost:8080
	router.Run(":8000")
}

func MainHandler(c *gin.Context) {
	c.String(200, "Inside Main Handler")
}
