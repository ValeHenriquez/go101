package models

type Task struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null;unique_index" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      string `json:"userId"`
}
