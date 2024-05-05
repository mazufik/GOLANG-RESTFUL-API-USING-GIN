package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mazufik/backend-api/models"
)

// type validation post input
type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// type error message
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// function get error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

// get all posts
func FindPosts(ctx *gin.Context) {
	// get data from database using model
	var posts []models.Post
	models.DB.Find(&posts)

	// return json
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Lists Data Posts",
		"data":    posts,
	})
}

// store a post
func StorePost(ctx *gin.Context) {
	// validate input
	var input ValidatePostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	// create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	// return response json
	ctx.JSON(201, gin.H{
		"status":  true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}

// get post by id
func FindPostById(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(200, gin.H{
		"succsess": true,
		"message":  "Detail Data Post By ID : " + ctx.Param("id"),
		"data":     post,
	})
}

// update post
func UpdatePost(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input ValidatePostInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": out})
		}
		return
	}

	// update post
	models.DB.Model(&post).Updates(input)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})

}

// delete post
func DeletePost(ctx *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&post).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// delete post
	models.DB.Delete(&post)

	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}
