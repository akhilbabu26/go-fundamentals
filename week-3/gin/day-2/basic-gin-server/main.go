package main

import "github.com/gin-gonic/gin"

// basic server
// func main(){
// 	r := gin.Default()

// 	r.GET("user", func(c *gin.Context){ 
// 		c.JSON(200, gin.H{
// 			"name": "akhil",
// 			"age": "24",
// 		})
// 	})

// 	r.Run(":8080")
// }

// multiple endpoints
func main(){
	r := gin.Default()

	r.GET("/api/userinfo", userInfo)
	r.GET("/api/hello", msg)
	r.GET("/api/data", data)

	r.Run(":8080")
}

func userInfo(c *gin.Context){
	c.JSON(200, gin.H{
		"name":"alhil",
		"age":"24",
	})
}

func msg(c *gin.Context){
	c.JSON(200, gin.H{
		"message":"Hello, World!!",
	})
}

func data(c *gin.Context){
	c.JSON(200, gin.H{
		"data":"123456",
	})
}