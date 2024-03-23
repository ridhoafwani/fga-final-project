package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/ridhoafwani/fga-final-project/utils"
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

	if err := utils.ValidateCreateSocialMedia(newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user id
	authenticatedUserId, _ := c.Get("user_id")
	newSocialMedia.UserID = authenticatedUserId.(uint)

	// Create social media account
	if err := h.DB.Create(&newSocialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media account"})
		return
	}

	socialMediaResponse := models.SocialMediaCreateResponse{
		ID:             newSocialMedia.ID,
		Name:           newSocialMedia.Name,
		SocialMediaURL: newSocialMedia.SocialMediaURL,
		UserID:         newSocialMedia.UserID,
		CreatedAt:      newSocialMedia.CreatedAt,
	}

	// Response with created social media account data
	c.JSON(http.StatusCreated, socialMediaResponse)
}

func (h *SocialMediaHandler) GetSocialMedias(c *gin.Context) {
	// Fetch all social media accounts
	var socialMedias []models.SocialMedia
	if err := h.DB.Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch social media accounts"})
		return
	}

	// Prepare response data
	var socialMediaResponses []models.SocialMediaResponse
	for _, socialMedia := range socialMedias {
		var user models.User
		if err := h.DB.First(&user, socialMedia.UserID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		socialMediaResponse := models.SocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaURL: socialMedia.SocialMediaURL,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: models.UserDetail{
				ID:       user.ID,
				Email:    user.Email,
				Username: user.Username,
			},
		}
		socialMediaResponses = append(socialMediaResponses, socialMediaResponse)
	}

	// Response with formatted social media accounts
	c.JSON(http.StatusOK, gin.H{"social_medias": socialMediaResponses})
}

// UpdateSocialMedia handles updating an existing social media account
func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	socialMediaID := c.Param("id")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := h.DB.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media account not found"})
		return
	}

	// Check authorization (user can only update their own social media account)
	authenticatedUserId, _ := c.Get("user_id")
	if socialMedia.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Binding JSON body into updated social media struct
	var updatedSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateCreateSocialMedia(updatedSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	socialMedia.Name = updatedSocialMedia.SocialMediaURL
	socialMedia.SocialMediaURL = updatedSocialMedia.SocialMediaURL
	socialMedia.UpdatedAt = time.Now()

	if err := h.DB.Save(&socialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update social media account"})
		return
	}

	socialMediaResponse := models.SocialMediaUpdateResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		UpdatedAt:      socialMedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, socialMediaResponse)
}

// DeleteSocialMedia handles deletion of an existing social media account
func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	socialMediaID := c.Param("id")
	authenticatedUserId, _ := c.Get("user_id")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := h.DB.First(&socialMedia, socialMediaID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media account not found"})
		return
	}

	// Check authorization (user can only delete their own social media account)
	if socialMedia.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete social media account
	h.DB.Delete(&socialMedia)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "your social media account has been successfully deleted"})
}
