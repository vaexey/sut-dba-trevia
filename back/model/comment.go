package model

type Comment struct {
	Id             uint `gorm:"primaryKey"`
	Content        string
	AttractionId   uint
	Attraction     Attraction      `gorm:"foreignKey:AttractionId"`
	CommentReports []CommentReport `gorm:"foreignKey:CommentId"`
	UserId         uint
	User           User `gorm:"foreignKey:UserId"`
	Comment		   string
}