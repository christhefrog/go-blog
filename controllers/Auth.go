package controllers

import (
	"net/http"

	"github.com/christhefrog/go-blog/database"
	"github.com/christhefrog/go-blog/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthenticateUser(c *gin.Context) {
	var body struct {
		Login    string
		Password string
	}

	session := sessions.Default(c)

	logged := session.Get("logged")

	if logged == true {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User already authenticated"})
		return
	}

	c.Bind(&body)
	if body.Login == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	result := database.Handle.Where("Name = ?", body.Login).Find(&user)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "User doesn't exist"})
		return
	}

	match := CheckPasswordHash(body.Password, user.Hash)

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password doesn't match"})
		return
	}

	session.Set("logged", true)
	session.Save()

	c.JSON(http.StatusOK, user)
}

func UnauthenticateUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("logged")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}
