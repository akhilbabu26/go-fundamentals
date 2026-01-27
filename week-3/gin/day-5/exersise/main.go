// 1️⃣ Validate user input before login (empty fields, invalid format).

// 2️⃣ Hash the password before storing it in memory.

// 3️⃣ Implement a protected dashboard page that requires authentication.

package main

import (
	"errors"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)
var sessions = make(map[string]string)

var mu sync.Mutex

func validateInput(username string, password string)error{
	if username == "" || password == ""{
		return errors.New("field cannot be empty")
	}

	if len(username) < 3{
		return errors.New("userneme must be at least 3 characters")
	}

	if len(password) < 6{
		return errors.New("password must be at least 6 characters")
	}

	return nil
}

func authMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		sessionID, err := c.Cookie("session_id")

		if err != nil{
			c.JSON(401, gin.H{"error": "Login required"})
			c.Abort()
			return
		}

		mu.Lock()
		_,ok := sessions[sessionID]
		mu.Unlock()

		if !ok{
			c.JSON(401, gin.H{"error":"Invalid input"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func register(c *gin.Context){
	var data struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&data); err != nil{
		c.JSON(400, gin.H{"error": "Invalid json"})
		return
	}

	if err := validateInput(data.Username, data.Password); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(data.Password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		c.JSON(500, gin.H{"error": "server error"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _,exists := users[data.Username]; exists{
		c.JSON(400, gin.H{"error":"user already exists"})
		return
	}

	users[data.Username] = string(hash)

	c.JSON(201, gin.H{"message": "registration succcessful"})
}

func login(c *gin.Context){
	var data struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validate input
	if err := validateInput(data.Username, data.Password); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	hash, ok := users[data.Username]
	mu.Unlock()

	if !ok {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify password
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(data.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create session
	sessionID := uuid.NewString()

	mu.Lock()
	sessions[sessionID] = data.Username
	mu.Unlock()

	// Set cookie (secure in production)
	c.SetCookie(
		"session_id",
		sessionID,
		3600,
		"/",
		"",
		false, // true in HTTPS
		true,  // HttpOnly
	)

	c.JSON(200, gin.H{"message": "Login successful"})
}

// Dashboard (Protected)
func dashboard(c *gin.Context) {

	sessionID, _ := c.Cookie("session_id")

	mu.Lock()
	username := sessions[sessionID]
	mu.Unlock()

	c.JSON(200, gin.H{
		"message": "Welcome to dashboard",
		"user":    username,
	})
}

// Logout
func logout(c *gin.Context) {

	sessionID, err := c.Cookie("session_id")

	if err == nil {
		mu.Lock()
		delete(sessions, sessionID)
		mu.Unlock()
	}

	c.SetCookie("session_id", "", -1, "/", "", false, true)

	c.JSON(200, gin.H{"message": "Logged out"})
}

func main() {

	r := gin.Default()

	// Public routes
	r.POST("/register", register)
	r.POST("/login", login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("/dashboard", dashboard)
		protected.POST("/logout", logout)
	}

	r.Run(":8080")
}