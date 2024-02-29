package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": true,
		})
	})

	ContactRoutes(router)

	router.Run(":3033") // listen and serve on 0.0.0.0:3033
}