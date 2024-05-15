package controllers

import (
	"net/http"

	"github.com/christhefrog/go-blog/database"
	"github.com/christhefrog/go-blog/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *gin.Context) {
	var body struct {
		Login    string
		Password string
	}

	c.Bind(&body)
	if body.Login == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	hash, err := HashPassword(body.Password)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user := models.User{Name: body.Login, Hash: hash}
	database.Handle.Create(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	result := database.Handle.Delete(&models.User{}, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user ID or user is already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}
