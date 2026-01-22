package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	// set cookie
	r.GET("/set-cookei", func(c *gin.Context){
		c.SetCookie("user", "Akhil", 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{"message":"Cookei set"})
	})

	// get cookie
	r.GET("/get-cookei", func(c *gin.Context){
		user, err := c.Cookie("user")
		if err != nil{
			c.JSON(400, gin.H{"error":"Cookei not found"})
			return
		}

		c.JSON(200, gin.H{"user": user})
	})

	// delete cookei
	r.GET("/delete-cookie", func(c *gin.Context){
		c.SetCookie("user", "", -1, "/", "localhost", false, true)
		c.JSON(200, gin.H{"message": "cookie delected"})
	})

	r.Run(":8080")
}