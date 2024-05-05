package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mazufik/backend-api/controllers"
	"github.com/mazufik/backend-api/models"
)

func main() {
	// inisialisasi gin
	router := gin.Default()

	// panggil koneksi database
	models.ConnectDatabase()

	// membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {
		// return response JSON
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// membuat route get all posts
	router.GET("/api/posts", controllers.FindPosts)

	// membuat route store post
	router.POST("/api/posts", controllers.StorePost)

	// membuat route detail post
	router.GET("/api/posts/:id", controllers.FindPostById)

	// membuat route update post
	router.PUT("/api/posts/:id", controllers.UpdatePost)

	// membuat route delete post
	router.DELETE("/api/posts/:id", controllers.DeletePost)

	// mulai server dengan port 3000
	router.Run(":3000")
}
