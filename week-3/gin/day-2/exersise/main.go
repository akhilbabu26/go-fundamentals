// 1️⃣ Build a simple web server with Gin that serves a homepage.

// 2️⃣ Add a route group /api for handling REST API requests.

// 3️⃣ Implement /users API that supports GET and POST requests.

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

var users = []User{}

func main(){
	r := gin.Default()

	r.GET("/", homePage)
	r.LoadHTMLGlob("template/*")

	api := r.Group("/api")
	{
		api.GET("/users", getUsers)
		api.POST("/users",createUser)
	}

	r.Run(":8080") // http://localhost:8080
}

func homePage(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin Web Server",
	})
}

func getUsers(c *gin.Context){
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context){
	var newUser User

	err := c.ShouldBindJSON(&newUser)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}
