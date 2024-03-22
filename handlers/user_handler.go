package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/ridhoafwani/fga-final-project/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	Db *gorm.DB
}

func InitUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{Db: db}
}

func (h UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user
	if err := utils.ValidateRegister(user, h.Db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password before storing in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save user to the database
	if err := h.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	userRegisterResponse := models.UserRegisterResponse{
		Id:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		Age:       user.Age,
	}

	c.JSON(http.StatusCreated, userRegisterResponse)
}

// Login handles user login
func (h UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	var user models.User
	if err := h.Db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	authenticatedUserId, _ := c.Get("user_id")

	// Authorization
	if userId != fmt.Sprint(authenticatedUserId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to update this user"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate user
	if err := utils.ValidateUpdateUser(user, h.Db); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	h.Db.First(&existingUser, userId)

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.UpdatedAt = time.Now()

	// Update user in the database
	if err := h.Db.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	var userUpdateResponse models.UserUpdateResponse
	userUpdateResponse.Id = int(existingUser.ID)
	userUpdateResponse.Username = existingUser.Username
	userUpdateResponse.Email = existingUser.Email
	userUpdateResponse.Age = existingUser.Age
	userUpdateResponse.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	c.JSON(http.StatusOK, userUpdateResponse)
}

func (h UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	authenticatedUserId, _ := c.Get("user_id")

	// Authorization
	if userId != fmt.Sprint(authenticatedUserId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to delete this user"})
		return
	}

	var user models.User
	h.Db.First(&user, userId)

	// Delete user from the database
	if err := h.Db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "your account has been successfully deleted"})
}
