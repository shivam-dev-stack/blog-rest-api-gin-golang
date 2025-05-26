package models

type Post struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`

	UserID uint // foreign key
	User   User `gorm:"foreignKey:UserID"` // association
}
