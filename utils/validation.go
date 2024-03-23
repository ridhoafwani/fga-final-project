package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/ridhoafwani/fga-final-project/models"
	"gorm.io/gorm"
)

func ValidateRegister(user models.User, db *gorm.DB) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}

	if err := validator.New().Var(user.Email, "required,email"); err != nil {
		return errors.New("invalid email format")
	}

	// Check email uniqueness
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}

	if err := validator.New().Var(user.Username, "required"); err != nil {
		return errors.New("username required")
	}

	// Check username uniqueness
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	// Custom validation for password length
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Custom validation for age
	if user.Age < 9 {
		return errors.New("age must be at least 9")
	}

	return nil
}

func ValidateUpdateUser(user models.User, db *gorm.DB) error {
	if err := validator.New().Struct(user); err != nil {
		return err
	}

	if err := validator.New().Var(user.Email, "required,email"); err != nil {
		return errors.New("invalid email format")
	}

	// Check email uniqueness
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already exists")
	}

	if err := validator.New().Var(user.Username, "required"); err != nil {
		return errors.New("username required")
	}

	// Check username uniqueness
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	return nil
}

func ValidateCreatePhoto(photo models.Photo) error {
	if err := validator.New().Struct(photo); err != nil {
		return err
	}

	if err := validator.New().Var(photo.Title, "required"); err != nil {
		return errors.New("title required")
	}

	if err := validator.New().Var(photo.PhotoURL, "required"); err != nil {
		return errors.New("photo url required")
	}

	return nil
}

func ValidateCreateComment(comment models.Comment) error {
	if err := validator.New().Struct(comment); err != nil {
		return err
	}

	if err := validator.New().Var(comment.Message, "required"); err != nil {
		return errors.New("message required")
	}

	return nil
}

func ValidateCreateSocialMedia(socialMedia models.SocialMedia) error {
	if err := validator.New().Struct(socialMedia); err != nil {
		return err
	}

	if err := validator.New().Var(socialMedia.Name, "required"); err != nil {
		return errors.New("name required")
	}

	if err := validator.New().Var(socialMedia.SocialMediaURL, "required"); err != nil {
		return errors.New("social media url required")
	}

	return nil
}
