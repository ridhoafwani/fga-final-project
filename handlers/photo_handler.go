package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"
)

type PhotoHandler struct {
	DB *gorm.DB
}

func InitPhotoHandler(db *gorm.DB) *PhotoHandler {
	return &PhotoHandler{
		DB: db,
	}
}

// CreatePhoto handles creation of new photo
func (h *PhotoHandler) CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check authorization (user can only create photo for their own user)

	// Save photo to database
	if err := h.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	// Response with created photo data
	c.JSON(http.StatusCreated, photo)
}

// GetPhotos handles fetching all photos
func (h *PhotoHandler) GetPhotos(c *gin.Context) {

	// Fetch all photos
	var photos []models.Photo
	if err := h.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	// Response with fetched photos
	c.JSON(http.StatusOK, photos)
}

// UpdatePhoto handles updating an existing photo
func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	// Fetch photo by ID
	var photo models.Photo
	if err := h.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
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
	h.DB.Model(&photo).Updates(updatedPhoto)

	// Response with updated photo data
	c.JSON(http.StatusOK, photo)
}

// DeletePhoto handles deletion of an existing photo
func (h *PhotoHandler) DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	// Fetch photo by ID
	var photo models.Photo
	if err := h.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check authorization (user can only delete their own photo)

	// Delete photo
	h.DB.Delete(&photo)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
