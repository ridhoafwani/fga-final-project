package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"
)

type CommentHandler struct {
	DB *gorm.DB
}

func InitCommentHandler(db *gorm.DB) *CommentHandler {
	return &CommentHandler{DB: db}
}

// CreateComment handles creation of new comment
func (h CommentHandler) CreateComment(c *gin.Context) {

	// Binding JSON body into Comment struct
	var newComment models.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	if err := h.DB.Create(&newComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Response with created comment data
	c.JSON(http.StatusCreated, newComment)
}

// GetComments handles fetching all comments
func (h CommentHandler) GetComments(c *gin.Context) {
	// Fetch all comments
	var comments []models.Comment
	if err := h.DB.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	// Response with fetched comments
	c.JSON(http.StatusOK, comments)
}

// UpdateComment handles updating an existing comment
func (h CommentHandler) UpdateComment(c *gin.Context) {
	commentID := c.Param("commentId")

	// Fetch comment by ID
	var comment models.Comment
	if err := h.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check authorization (user can only update their own comment)

	// Binding JSON body into updated comment struct
	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update comment
	h.DB.Model(&comment).Updates(updatedComment)

	// Response with updated comment data
	c.JSON(http.StatusOK, comment)
}

// DeleteComment handles deletion of an existing comment
func (h CommentHandler) DeleteComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	commentID := c.Param("commentId")

	// Fetch comment by ID
	var comment models.Comment
	if err := db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check authorization (user can only delete their own comment)

	// Delete comment
	db.Delete(&comment)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
