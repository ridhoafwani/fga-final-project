package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username     string
    Email        string `gorm:"unique_index"`
    Password     string
    Age          int
    Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE"`
    Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE"`
    SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE"`
}

type Photo struct {
    gorm.Model
    Title     string
    Caption   string
    PhotoURL  string `json:"photo_url"`
    UserID    uint   `json:"user_id"`
	User	User	`gorm:"foreignKey:UserID"`
    Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE"`
}

type Comment struct {
    gorm.Model
    UserID    uint `json:"user_id"`
    PhotoID   uint `json:"photo_id"`
	User User `gorm:"foreignKey:UserID"`
	Photo Photo `gorm:"foreignKey:PhotoID"`
    Message   string
}

type SocialMedia struct {
    gorm.Model
    Name           string
    SocialMediaURL string `json:"social_media_url"`
    UserID         uint   `json:"user_id"`
	User	User `gorm:"foreignKey:UserID"`
}