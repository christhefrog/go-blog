package controllers

import (
	"log"
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

func DeletePost(c *gin.Context) {
	var body struct {
		Post uint
	}

	c.Bind(&body)
	log.Print(body.Post)

	result := database.Handle.Delete(&models.Post{}, body.Post)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong post ID or post is already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": body.Post, "deleted": true})
}

func GetPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 1
	}

	if page < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pages start at 1"})
		return
	}

	var posts []models.Post
	database.Handle.Where("id >= ?", (page-1)*PAGE_SIZE).Limit(PAGE_SIZE).Find(&posts)

	c.JSON(http.StatusOK, gin.H{"page": page, "posts": posts})
}

func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := database.Handle.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"eror": "Wrong post ID or post is deleted"})
		return
	}

	c.JSON(http.StatusOK, post)
}
