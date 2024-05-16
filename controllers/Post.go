package controllers

import (
	"net/http"
	"strconv"

	"github.com/christhefrog/go-blog/database"
	"github.com/christhefrog/go-blog/models"
	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 10

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	if body.Body == "" || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	post := models.Post{Body: body.Body, Title: body.Title}
	result := database.Handle.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	if body.Body == "" || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var post models.Post
	result := database.Handle.Find(&post, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "Wrong post ID"})
		return
	}

	database.Handle.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := database.Handle.Delete(&models.Post{}, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong post ID or post is already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func GetPosts(c *gin.Context) {
	from, err := strconv.Atoi(c.Query("from"))

	if err != nil {
		from = 0
	}

	if from < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "From should be greater than 0"})
		return
	}

	var posts []models.Post

	if from == 0 {
		database.Handle.Limit(PAGE_SIZE).Order("id desc").Find(&posts)
	} else {
		database.Handle.Limit(PAGE_SIZE).Where("id <= ?", from).Order("id desc").Find(&posts)
	}

	c.JSON(http.StatusOK, gin.H{"from": from, "posts": posts})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := database.Handle.Find(&post, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "Wrong post ID"})
		return
	}

	c.JSON(http.StatusOK, post)
}
