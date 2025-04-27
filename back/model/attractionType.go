package model

type AttractionType struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Attractions []Attraction `gorm:"foreignKey:AttractionTypeId"`
}