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
	commentHandler := handlers.InitCommentHandler(db)
	socialMediaHandler := handlers.InitSocialMediaHandler(db)

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
		photoGroup.PUT("/:id", photoHandler.UpdatePhoto)
		photoGroup.DELETE("/:id", photoHandler.DeletePhoto)
	}

	// Comment endpoints
	commentGroup := router.Group("/comments")
	commentGroup.Use(middleware.AuthMiddleware())
	{
		commentGroup.POST("/", commentHandler.CreateComment)
		commentGroup.GET("/", commentHandler.GetComments)
		commentGroup.PUT("/:id", commentHandler.UpdateComment)
		commentGroup.DELETE("/:id", commentHandler.DeleteComment)
	}

	// SocialMedia endpoints
	socialMediaGroup := router.Group("/socialmedias")
	socialMediaGroup.Use(middleware.AuthMiddleware())
	{
		socialMediaGroup.POST("/", socialMediaHandler.CreateSocialMedia)
		socialMediaGroup.GET("/", socialMediaHandler.GetSocialMedias)
		socialMediaGroup.PUT("/:id", socialMediaHandler.UpdateSocialMedia)
		socialMediaGroup.DELETE("/:id", socialMediaHandler.DeleteSocialMedia)
	}

	router.GET("", WelcomeHandler)

	router.Run(":8001")
}
