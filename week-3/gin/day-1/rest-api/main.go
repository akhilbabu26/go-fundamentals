package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/users", getUsers)
	r.POST("/users", createUser)

	r.Run(":9000") //http://localhost:9000/users
}

func getUsers(c *gin.Context){
	c.JSON(200, gin.H{"Massage": "Get all users"})
}

func createUser(c *gin.Context){
	c.JSON(201, gin.H{"Msg2": "user Created"})
}