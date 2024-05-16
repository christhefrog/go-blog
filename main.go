package main

import (
	"log"
	"os"

	"github.com/christhefrog/go-blog/controllers"
	"github.com/christhefrog/go-blog/database"
	"github.com/christhefrog/go-blog/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load the .env file")
	}
}

func useSessions(r *gin.Engine) {
	store := cookie.NewStore([]byte(os.Getenv("SECRET")))
	r.Use(sessions.Sessions("session", store))
}

func main() {
	loadDotEnv()

	database.Connect()
	database.Migrate()

	r := gin.Default()
	useSessions(r)

	api := r.Group("/api")
	{
		api.POST("/auth", controllers.AuthenticateUser)

		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPostByID)

		api.Use(middleware.Protected())

		api.DELETE("/auth", controllers.UnauthenticateUser)

		api.POST("/posts", controllers.CreatePost)
		api.PUT("/posts/:id", controllers.UpdatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)

		api.POST("/users", controllers.CreateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}

	controllers.DemoLoadHTML(r)
	r.GET("/demo", controllers.DemoIndex)

	r.Run()
}
