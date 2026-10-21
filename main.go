package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()
	router.LoadHTMLFiles("html/index.html")
	// Define a simple GET endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/submit", func(ctx *gin.Context) {
		password := ctx.PostForm("password")
		fmt.Printf("Password: %s\n", password)
		ctx.Redirect(http.StatusFound, "/")
	})

	router.GET("/hacking", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
