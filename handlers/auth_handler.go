package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/fga-final-project/models"
	"github.com/ridhoafwani/fga-final-project/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Db *gorm.DB
}

func InitAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{Db: db}
}

func (h AuthHandler) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate user
    if err := utils.ValidateUser(user, h.Db); err != nil {
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

    c.JSON(http.StatusCreated, user)
}



// Login handles user login
func (h AuthHandler) Login(c *gin.Context) {
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
