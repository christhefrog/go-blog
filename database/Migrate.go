package database

import (
	"github.com/christhefrog/go-blog/models"
)

func Migrate() {
	Handle.AutoMigrate(&models.User{})
	Handle.AutoMigrate(&models.Post{})
}
