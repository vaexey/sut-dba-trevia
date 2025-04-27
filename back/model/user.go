package model

type User struct {
	Id                uint   `gorm:"primaryKey"`
	Login             string `gorm:"unique"`
	Password          string
	Name              string
	Attractions       []Attraction       `gorm:"foreignKey:UserId"` // attraction
	AttractionReports []AttractionReport `gorm:"foreignKey:UserId"` // attraction reports
	Ratings           []Rating           `gorm:"foreignKey:UserId"` // ratings
	CommentReports    []CommentReport    `gorm:"foreignKey:UserId"` // comment reports
	RoleId            uint               // role
	Role              Role               `gorm:"foreignKey:RoleId"`
}