package model

type Rating struct {
	Id           uint `gorm:"primaryKey"`
	Rating       int
	AttractionId uint       // attraction
	Attraction   Attraction `gorm:"foreignKey:AttractionId"`
	UserId       uint       // user
	User         User       `gorm:"foreignKey:UserId"`
}