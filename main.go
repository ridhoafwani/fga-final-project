package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/database"
	"github.com/ridhoafwani/fga-final-project/handlers"
	middleware "github.com/ridhoafwani/fga-final-project/middlewares"
)

func WelcomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Welcome")
}

func main() {
	router := gin.Default()
	db := database.DatabaseConnection()

	userHandler := handlers.InitUserHandler(db)
	photoHandler := handlers.InitPhotoHandler(db)

	// Auth endpoints
	authGroup := router.Group("/users")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}

	// User endpoints
	userGroup := router.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}

	// Photo endpoints
	photoGroup := router.Group("/photos")
	photoGroup.Use(middleware.AuthMiddleware())
	{
		photoGroup.POST("/", photoHandler.CreatePhoto)
		photoGroup.GET("/", photoHandler.GetPhotos)
		photoGroup.PUT("/:photoId", photoHandler.UpdatePhoto)
		photoGroup.DELETE("/:photoId", photoHandler.DeletePhoto)
	}

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

	router.GET("", WelcomeHandler)

	router.Run(":8001")
}
