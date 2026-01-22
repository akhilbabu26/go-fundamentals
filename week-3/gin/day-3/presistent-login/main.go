package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// fake database
var users = map[string]string{
	"akhil": "1234",
}

// Session store
var sessions = map[string]string{}

func main(){
	r := gin.Default()

	r.POST("/login", login)
	r.GET("/profile", authMiddleware(), profile)
	r.GET("/logout", logout)

	r.Run(":8080")
}

// Login api
func login(c *gin.Context){
	var data struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BindJSON(&data)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// checking if database has password and username
	pass, ok := users[data.Username]

	if !ok || pass != data.Password{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid password"})
		return
	}

	// create session
	sessionID := "session_" + data.Username
	sessions[sessionID] = data.Username

	// Store session in cookie
	c.SetCookie("session_id", sessionID, 86400, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// auth middleware
func authMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		sessionID, err := c.Cookie("session_id")
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"not logged in"})
			c.Abort()
			return
		}

		username, ok := sessions[sessionID]

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			c.Abort()
			return
		}

		c.Set("user", username)
		c.Next()
	}
}

//protect route
func profile(c *gin.Context){
	user,_:= c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome" + user.(string),
	})
}

//logout
func logout(c *gin.Context){
	sessionID, err := c.Cookie("session_id")
	if err == nil{
		delete(sessions, sessionID)
	}

	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}