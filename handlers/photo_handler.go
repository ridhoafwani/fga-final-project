package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"
)

// CreatePhoto handles creation of new photo
func CreatePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Binding JSON body into Photo struct
	var newPhoto models.Photo
	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create photo
	if err := db.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	// Response with created photo data
	c.JSON(http.StatusCreated, newPhoto)
}

// GetPhotos handles fetching all photos
func GetPhotos(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Fetch all photos
	var photos []models.Photo
	if err := db.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	// Response with fetched photos
	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto handles updating an existing photo
func UpdatePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	photoID := c.Param("photoId")

	// Fetch photo by ID
	var photo models.Photo
	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check authorization (user can only update their own photo)

	// Binding JSON body into updated photo struct
	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update photo
	db.Model(&photo).Updates(updatedPhoto)

	// Response with updated photo data
	c.JSON(http.StatusOK, photo)
}

// DeletePhoto handles deletion of an existing photo
func DeletePhoto(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	photoID := c.Param("photoId")

	// Fetch photo by ID
	var photo models.Photo
	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check authorization (user can only delete their own photo)

	// Delete photo
	db.Delete(&photo)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
