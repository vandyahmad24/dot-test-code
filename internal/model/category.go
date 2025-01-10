package model

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`

	Books []*Book `gorm:"foreignKey:CategoryID" json:"books"`
}
