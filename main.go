package main

import (
	"mygram_project/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Auth endpoints
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register)
		authGroup.POST("/login", handlers.Login)
	}

	// Photo endpoints
	photoGroup := router.Group("/photos")
	photoGroup.Use(handlers.AuthMiddleware())
	{
		photoGroup.POST("/", handlers.CreatePhoto)
		photoGroup.GET("/", handlers.GetPhotos)
		photoGroup.PUT("/:photoId", handlers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", handlers.DeletePhoto)
	}

	// Comment endpoints
	commentGroup := router.Group("/comments")
	commentGroup.Use(handlers.AuthMiddleware())
	{
		commentGroup.POST("/", handlers.CreateComment)
		commentGroup.GET("/", handlers.GetComments)
		commentGroup.PUT("/:commentId", handlers.UpdateComment)
		commentGroup.DELETE("/:commentId", handlers.DeleteComment)
	}

	// SocialMedia endpoints
	socialMediaGroup := router.Group("/socialmedias")
	socialMediaGroup.Use(handlers.AuthMiddleware())
	{
		socialMediaGroup.POST("/", handlers.CreateSocialMedia)
		socialMediaGroup.GET("/", handlers.GetSocialMedias)
		socialMediaGroup.PUT("/:socialMediaId", handlers.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:socialMediaId", handlers.DeleteSocialMedia)
	}

	router.Run(":8080")
}
