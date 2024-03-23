package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/ridhoafwani/fga-final-project/utils"
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

	if err := utils.ValidateCreateComment(newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user id
	authenticatedUserId, _ := c.Get("user_id")
	newComment.UserID = authenticatedUserId.(uint)

	// chek if photoid exists
	var photo models.Photo
	if err := h.DB.First(&photo, newComment.PhotoID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Photo not found"})
		return
	}

	// Create comment
	if err := h.DB.Create(&newComment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	commentCreateResponse := models.CommentCreateResponse{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
	}

	// Response with created comment data
	c.JSON(http.StatusCreated, commentCreateResponse)
}

func (h CommentHandler) GetComments(c *gin.Context) {
	// Fetch all comments
	var comments []models.Comment
	if err := h.DB.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	// Prepare response data
	var commentResponses []models.CommentResponse
	for _, comment := range comments {
		var user models.User
		if err := h.DB.First(&user, comment.UserID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
			return
		}

		var photo models.Photo
		if err := h.DB.First(&photo, comment.PhotoID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photo"})
			return
		}

		commentResponse := models.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: models.UserDetail{
				ID:       user.ID,
				Email:    user.Email,
				Username: user.Username,
			},
			Photo: models.PhotoDetail{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoURL,
				UserID:   photo.UserID,
			},
		}
		commentResponses = append(commentResponses, commentResponse)
	}

	// Response with formatted comments
	c.JSON(http.StatusOK, commentResponses)
}

// UpdateComment handles updating an existing comment
func (h CommentHandler) UpdateComment(c *gin.Context) {
	commentID := c.Param("id")

	// Fetch comment by ID
	var comment models.Comment
	if err := h.DB.First(&comment, commentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check authorization (user can only update their own comment)
	authenticatedUserId, _ := c.Get("user_id")
	if comment.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Binding JSON body into updated comment struct
	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateCreateComment(updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Message = updatedComment.Message
	comment.UpdatedAt = time.Now()

	if err := h.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	commentResponse := models.CommentUpdateResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, commentResponse)
}

// DeleteComment handles deletion of an existing comment
func (h CommentHandler) DeleteComment(c *gin.Context) {
	commentIDStr := c.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}
	authenticatedUserId, _ := c.Get("user_id")

	// Fetch comment by ID
	var comment models.Comment
	if err := h.DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	// Check authorization (user can only delete their own comment)
	if comment.UserID != authenticatedUserId.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	// Delete comment
	h.DB.Delete(&comment)

	// Response with success message
	c.JSON(http.StatusOK, gin.H{"message": "your comment has been successfully deleted"})
}
