package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/msg", func(c *gin.Context){
		c.JSON(200, gin.H{
			"Massage": "Welcome to Gin",
		})
	})

	r.Run(":4000") // localhost:4000/msg
}
