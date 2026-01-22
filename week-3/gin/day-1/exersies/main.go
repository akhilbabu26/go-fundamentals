// 1️⃣ Create a simple Go server that handles GET and POST requests.

// 2️⃣ Implement an API that returns a JSON response (/api/hello -> { "message": "Hello, World!" }).

// 3️⃣ Create an endpoint /api/user that accepts POST requests with JSON data (name, age).

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

var userList []User

func main(){
	r := gin.Default()

	r.GET("/api/hello", hello)
	r.GET("/api/users", getUsers)
	r.POST("/api/users", users)

	r.Run(":8080")// http://localhost:8080
}

func hello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello, World!",
	})
}

func users(c *gin.Context){
	var user User

	err := c.ShouldBindJSON(&user); 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
		return
	}

	userList = append(userList, user)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": user,
	})
}

func getUsers(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"user": userList,
	})
}