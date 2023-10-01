package model

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID string `json:"userId"`
	User        *User  `json:"User"`
}