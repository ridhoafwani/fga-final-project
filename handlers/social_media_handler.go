package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"

	"net/http"
)

// CreateSocialMedia handles creation of new social media account
func CreateSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Binding JSON body into SocialMedia struct
	var newSocialMedia models.SocialMedia
	if err := c.ShouldBindJSON(&newSocialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create social media account
	if err := db.Create(&newSocialMedia).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create social media account"})
		return
	}

	// Response with created social media account data
	c.JSON(http.StatusCreated, newSocialMedia)
}

// GetSocialMedias handles fetching all social media accounts
func GetSocialMedias(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Fetch all social media accounts
	var socialMedias []models.SocialMedia
	if err := db.Find(&socialMedias).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch social media accounts"})
		return
	}

	// Response with fetched social media accounts
	c.JSON(http.StatusOK, gin.H{"social_medias": socialMedias})
}

// UpdateSocialMedia handles updating an existing social media account
func UpdateSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	socialMediaID := c.Param("socialMediaId")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := db.Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
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
	db.Model(&socialMedia).Updates(updatedSocialMedia)

	// Response with updated social media account data
	c.JSON(http.StatusOK, socialMedia)
}

// DeleteSocialMedia handles deletion of an existing social media account
func DeleteSocialMedia(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	socialMediaID := c.Param("socialMediaId")

	// Fetch social media account by ID
	var socialMedia models.SocialMedia
	if err := db.Where("id = ?", socialMediaID).First(&socialMedia).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media account not found"})
		return
	}

	// Check authorization (user can only delete their own social media account)

	// Delete social media account
	db.Delete(&socialMedia)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "Social media account deleted successfully"})
}
