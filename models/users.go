package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty"`

	Posts []Post
}
