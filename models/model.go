package models

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Email     string `gorm:"unique_index"`
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []Photo   `gorm:"foreignkey:UserID"` // One-to-many relationship
	Comments  []Comment `gorm:"foreignkey:UserID"` // One-to-many relationship
}

type Photo struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Caption   string
	PhotoURL  string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignkey:UserID"` // Belongs-to relationship
}

type Comment struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint
	PhotoID   uint
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User  `gorm:"foreignkey:UserID"`  // Belongs-to relationship
	Photo     Photo `gorm:"foreignkey:PhotoID"` // Belongs-to relationship
}

type SocialMedia struct {
	ID             uint `gorm:"primary_key"`
	Name           string
	SocialMediaURL string
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User `gorm:"foreignkey:UserID"` // Belongs-to relationship
}
