package main

import (
	"log"

	"github.com/christhefrog/go-blog/controllers"
	"github.com/christhefrog/go-blog/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load the .env file")
	}
}

func main() {
	loadDotEnv()

	database.Connect()
	database.Migrate()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/createPost", controllers.CreatePost)
		api.POST("/deletePost", controllers.DeletePost)
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPostByID)
	}

	r.Run()
}
