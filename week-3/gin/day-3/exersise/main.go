package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/google/uuid"
)

// data base hardcoded user database
var users = map[string]string{
	"akhil": "1234",
}

// session store
var sessions = map[string]string{}

func main(){
	r := gin.Default()

	r.POST("/login", login) // http://localhost:8080/login
	r.GET("/profile", authMiddleware(), profile)
	r.GET("/logout", logout)

	r.Run(":8080")
}

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

	pass, ok := users[data.Username]
	if !ok || pass != data.Password{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password or username"})
		return
	}

	sessionID := uuid.NewString()
	sessions[sessionID] = data.Username

	c.SetCookie("session_id", sessionID, 86400, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

//middleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		sessionID, err := c.Cookie("session_id")
		if err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"Not logged in"})
			c.Abort()
			return
		}

		username, ok := sessions[sessionID]
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid session"})
			c.Abort()
			return
		}

		c.Set("user", username)
		c.Next()
	}
}

//protected route
func profile(c *gin.Context){
	user,_:= c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message":"welcome"+ user.(string)})
}

//logout
func logout(c *gin.Context){
	sessionID, err := c.Cookie("session_id")
	if err == nil{
		delete(sessions, sessionID)
	}
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}