package models

import "time"

type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string `gorm:"unique_index"`
	Password     string
	Age          int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo       `gorm:"foreignkey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments     []Comment     `gorm:"foreignkey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SocialMedias []SocialMedia `gorm:"foreignkey:UserID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Photo struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Caption   string
	PhotoURL  string `json:"photo_url"`
	UserID    uint   `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment `gorm:"foreignkey:PhotoID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Comment struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `json:"user_id"`
	PhotoID   uint `json:"photo_id"`
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SocialMedia struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint   `json:"user_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
