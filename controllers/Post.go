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
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 0
	}

	if page < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pages start at 0"})
		return
	}

	var posts []models.Post
	database.Handle.Limit(PAGE_SIZE).Offset(page * PAGE_SIZE).Find(&posts)

	c.JSON(http.StatusOK, gin.H{"page": page, "posts": posts})
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
