// // 1️⃣ Create logging middleware that logs all requests to login and logout.

// // 2️⃣ Implement an authentication middleware that restricts access to a protected route.

// // 3️⃣ Add rate limiting to protect against excessive API calls.

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var users = map[string]string{
	"akhil": "1234",
}

var sessions = map[string]string{}

type Client struct {
	Requests int
}

var clients = make(map[string]*Client)

const limit = 5

// Logging Middleware
func loginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		log.Printf(
			"[LOGIN/LOGOUT] %s %s | IP: %s\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
		)
	}
}

// Auth Middleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionID, err := c.Cookie("session_id")

		if err != nil {
			c.JSON(401, gin.H{"error": "Not logged in"})
			c.Abort()
			return
		}

		_, ok := sessions[sessionID]

		if !ok {
			c.JSON(401, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Rate Limiter
func rateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()

		client, exists := clients[ip]

		if !exists {
			clients[ip] = &Client{
				Requests: 1,
			}
			c.Next()
			return
		}

		// Check limit
		if client.Requests >= limit {
			c.JSON(429, gin.H{
				"error": "Too many requests (limit reached)",
			})
			c.Abort()
			return
		}

		client.Requests++

		c.Next()
	}
}

// Login
func login(c *gin.Context) {

	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	pass, ok := users[data.Username]

	if !ok || pass != data.Password {
		c.JSON(401, gin.H{"error": "Wrong credentials"})
		return
	}

	sessionID := uuid.NewString()
	sessions[sessionID] = data.Username

	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true)

	c.JSON(200, gin.H{"message": "Login successful"})
}

// Logout
func logout(c *gin.Context) {

	sessionID, err := c.Cookie("session_id")

	if err == nil {
		delete(sessions, sessionID)
	}

	c.SetCookie("session_id", "", -1, "/", "", false, true)

	c.JSON(200, gin.H{"message": "Logged out"})
}

// Protected Route
func profile(c *gin.Context) {

	sessionID, _ := c.Cookie("session_id")
	username := sessions[sessionID]

	c.JSON(200, gin.H{
		"user": username,
	})
}

func main() {

	r := gin.Default()

	// Rate Limiter
	r.Use(rateLimiter())

	// Logger only for login/logout
	logGroup := r.Group("/", loginLogger())
	{
		logGroup.POST("/login", login)
		logGroup.GET("/logout", logout)
	}

	// Protected Route
	r.GET("/profile", authMiddleware(), profile)

	r.Run(":8080")
}
