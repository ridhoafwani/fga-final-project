package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"

	"net/http"
)

type SocialMediaHandler struct {
	DB *gorm.DB
}

func InitSocialMediaHandler(db *gorm.DB) *SocialMediaHandler {
	return &SocialMediaHandler{
		DB: db,
	}
}

// CreateSocialMedia handles creation of new social media account
func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context) {
	// Binding JSON body into SocialMedia struct
	var newSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create social media account
	if err := h.DB.Create(&newSocialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media account"})
		return
	}

	// Response with created social media account data
	c.JSON(http.StatusCreated, newSocialMedia)
}

// GetSocialMedias handles fetching all social media accounts
func (h *SocialMediaHandler) GetSocialMedias(c *gin.Context) {
	// Fetch all social media accounts
	var socialMedias []models.SocialMedia
	if err := h.DB.Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch social media accounts"})
		return
	}

	// Response with fetched social media accounts
	c.JSON(http.StatusOK, gin.H{"social_medias": socialMedias})
}

// UpdateSocialMedia handles updating an existing social media account
func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	socialMediaID := c.Param("socialMediaId")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := h.DB.Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media account not found"})
		return
	}

	// Check authorization (user can only update their own social media account)

	// Binding JSON body into updated social media struct
	var updatedSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update social media account
	h.DB.Model(&socialMedia).Updates(updatedSocialMedia)

	// Response with updated social media account data
	c.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMedia handles deletion of an existing social media account
func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	socialMediaID := c.Param("socialMediaId")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := h.DB.Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media account not found"})
		return
	}

	// Check authorization (user can only delete their own social media account)

	// Delete social media account
	h.DB.Delete(&socialMedia)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "Social media account deleted successfully"})
}
