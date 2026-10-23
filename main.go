package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type InputGuess struct {
	Hash string `json:"hash" binding:"required"`
}

func main() {
	var password string
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()
	router.Static("/assets", "./assets")
	// router.LoadHTMLFiles("html/index.html", "html/hacking.tmpl")
	router.LoadHTMLGlob("templates/*")

	// Define a simple GET endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", nil)
	})

	router.POST("/submit", func(ctx *gin.Context) {
		password = ctx.PostForm("password")
		fmt.Printf("Password: %s\n", password)
		passwordHash, err := GeneratePassword(password)
		if err != nil {
			fmtErr := url.QueryEscape(err.Error())
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/error?msg=%s", fmtErr))
		}
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/hacking?hash=%s", passwordHash))
	})

	router.GET("/hacking", func(ctx *gin.Context) {
		hash := ctx.Query("hash")
		ctx.HTML(http.StatusOK, "hacking.tmpl", gin.H{
			"hash": hash,
		})
	})

	router.GET("/error", func(ctx *gin.Context) {
		msg := ctx.Query("msg")
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": msg,
		})
	})

	router.GET("/success", func(ctx *gin.Context) {
		pswd := ctx.Query("password")
		ctx.HTML(http.StatusOK, "success.tmpl", gin.H{
			"password": pswd,
		})
	})

	router.POST("/guess", func(ctx *gin.Context) {
		var json InputGuess
		if err := ctx.ShouldBindJSON(&json); err != nil {
			fmtErr := url.QueryEscape(err.Error())
			ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/error?msg=%s", fmtErr))
			return
		}
		pswd := HackPassword(json.Hash)
		ctx.JSON(http.StatusOK, gin.H{"password": pswd})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
