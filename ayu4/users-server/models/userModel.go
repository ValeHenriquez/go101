package models

type User struct {
	ID          string `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Password  string `gorm:"not null;unique_index" json:"password"`
}
