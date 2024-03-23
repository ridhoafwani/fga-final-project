package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/ridhoafwani/fga-final-project/utils"
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

	// Get user id
	authenticatedUserId, _ := c.Get("user_id")
	userId := authenticatedUserId.(uint)
	photo.UserID = userId

	// Validate photo data
	if err := utils.ValidateCreatePhoto(photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save photo to database
	if err := h.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	photoCreateResponse := models.PhotoCreateResponse{
		Id:        int(photo.ID),
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoURL,
		UserId:    int(photo.UserID),
		CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// Response with created photo data
	c.JSON(http.StatusCreated, photoCreateResponse)
}

func (h *PhotoHandler) GetPhotos(c *gin.Context) {
    // Fetch all photos
    var photos []models.Photo
    if err := h.DB.Preload("User").Find(&photos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
        return
    }

    // Prepare response data
    var photoResponses []models.PhotoGetsResponse
    for _, photo := range photos {
		if photo.User.ID == 0 {
            continue
        }
        photoResponse := models.PhotoGetsResponse{
            ID:        photo.ID,
            Title:     photo.Title,
            Caption:   photo.Caption,
            PhotoURL:  photo.PhotoURL,
            UserID:    photo.UserID,
            CreatedAt: photo.CreatedAt,
            UpdatedAt: photo.UpdatedAt,
            User: models.UserDetail{
				ID:       photo.User.ID,
                Email:    photo.User.Email,
                Username: photo.User.Username,
            },
        }
        photoResponses = append(photoResponses, photoResponse)
    }

    // Response with formatted photos
    c.JSON(http.StatusOK, photoResponses)
}

// UpdatePhoto handles updating an existing photo
func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {
	photoID := c.Param("id")

	var photo models.Photo
	if err := h.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Authorization
	authenticatedUserId, _ := c.Get("user_id")
	if photo.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Binding JSON body into updated photo struct
	var updatedPhoto models.Photo
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateCreatePhoto(updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photo.Title = updatedPhoto.Title
	photo.Caption = updatedPhoto.Caption
	photo.PhotoURL = updatedPhoto.PhotoURL
	photo.UpdatedAt = time.Now()

	if err := h.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	photoResponse := models.PhotoUpdateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	// Response with updated photo data
	c.JSON(http.StatusOK, photoResponse)
}

// DeletePhoto handles deletion of an existing photo
func (h *PhotoHandler) DeletePhoto(c *gin.Context) {
	photoID := c.Param("id")
	authenticatedUserId, _ := c.Get("user_id")

	// Fetch photo by ID
	var photo models.Photo
	if err := h.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Check authorization (user can only delete their own photo)
	if photo.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete photo
	h.DB.Delete(&photo)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "your photo has been successfully deleted"})
}
