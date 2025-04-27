package model

type AttractionReport struct {
	Id           uint `gorm:"primaryKey"`
	Content      string
	AttractionId uint       // attraction
	Attraction   Attraction `gorm:"foreignKey:AttractionId"`
	UserId       uint       // user
	User         User       `gorm:"foreignKey:UserId"`
}