package model

type Book struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"size:255;not null" json:"title"`
	Author     string `gorm:"size:255;not null" json:"author"`
	CategoryID uint   `gorm:"not null" json:"category_id"`

	Category *Category `gorm:"foreignKey:CategoryID" json:"category"`
}
