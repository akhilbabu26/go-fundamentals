package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type LoginInput struct{
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}

func main(){
	r := gin.Default()

	r.POST("/login", loginHandler)
	r.Run(":8080")

}

func loginHandler(c *gin.Context){
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login data is valid",
		"username": input.Username,
	})
}