package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default()

	r.GET("/set-cookie", setInCookie)
	r.GET("/get-cookie", getFromCookei)
	r.GET("/delete-cookie", deleteCookie)

	r.Run(":8080")
}

// set cookie
func setInCookie(c *gin.Context){
	c.SetCookie("user", "Akhil", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user set in cookei"})
}

// get cookie
func getFromCookei(c *gin.Context){
	user, err := c.Cookie("user")

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name":user})
}

// delete cookei
func deleteCookie(c *gin.Context){
	c.SetCookie("user", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user delete successfully"})
}
