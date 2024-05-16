package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DemoLoadHTML(r *gin.Engine) {
	r.LoadHTMLGlob("./demo/*")
}

func DemoIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
