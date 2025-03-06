package main

import (
	"simple-wallet/internal/app"
)

func main() {
	// router := gin.Default()

	app.StartServer()

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// router.Run(":8585")
}
