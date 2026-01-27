package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// database
var users = make(map[string]string)

type AuthInput struct{
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3,max=6"`
}

func main(){
	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)

	r.Run(":8080")
}

func register(c *gin.Context){
	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := users[input.Username]; exists{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		c.JSON(500, gin.H{"error": "faild to hash password"})
		return
	}

	users[input.Username] = string(hashedPassword)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func login(c *gin.Context) {

	var input AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Get stored hash
	hashedPassword, exists := users[input.Username]

	if !exists {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare password with hash
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
	})
}