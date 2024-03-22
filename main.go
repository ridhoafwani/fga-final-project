package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/handlers"
)

func main() {
	router := gin.Default()
	db := database.DatabaseConnection()

	authHandler := handlers.InitAuthHandler(db)

	// Auth endpoints
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// // Photo endpoints
	// photoGroup := router.Group("/photos")
	// photoGroup.Use(middleware.AuthMiddleware())
	// {
	// 	photoGroup.POST("/", handlers.CreatePhoto)
	// 	photoGroup.GET("/", handlers.GetPhotos)
	// 	photoGroup.PUT("/:photoId", handlers.UpdatePhoto)
	// 	photoGroup.DELETE("/:photoId", handlers.DeletePhoto)
	// }

	// // Comment endpoints
	// commentGroup := router.Group("/comments")
	// commentGroup.Use(middleware.AuthMiddleware())
	// {
	// 	commentGroup.POST("/", handlers.CreateComment)
	// 	commentGroup.GET("/", handlers.GetComments)
	// 	commentGroup.PUT("/:commentId", handlers.UpdateComment)
	// 	commentGroup.DELETE("/:commentId", handlers.DeleteComment)
	// }

	// // SocialMedia endpoints
	// socialMediaGroup := router.Group("/socialmedias")
	// socialMediaGroup.Use(middleware.AuthMiddleware())
	// {
	// 	socialMediaGroup.POST("/", handlers.CreateSocialMedia)
	// 	socialMediaGroup.GET("/", handlers.GetSocialMedias)
	// 	socialMediaGroup.PUT("/:socialMediaId", handlers.UpdateSocialMedia)
	// 	socialMediaGroup.DELETE("/:socialMediaId", handlers.DeleteSocialMedia)
	// }

	router.GET("", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	router.Run(":8080")
}
