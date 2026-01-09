package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goApp/models"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Request could not be parsed"})
		return
	}

	err := user.Save()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User saved successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Request could not be parsed"})
		return
	}

	err := user.ValidateCredentials()

	if err != nil {
		c.JSON(401, gin.H{"error": "Could not authenticate user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}
