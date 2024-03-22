package models

import "time"

type UserUpdateResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	UpdatedAt string `json:"updated_at"`
}

type UserRegisterResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

type PhotoCreateResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PhotoGetsResponse struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	User      UserDetail `json:"User"`
}

type PhotoUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetail struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
