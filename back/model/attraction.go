package model

type Attraction struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Description string
	FunFact     string
	Photo       string

	// region
	RegionId uint
	Region   Region `gorm:"foreignKey:RegionId"`

	// attraction type
	AttractionTypeId uint
	AttractionType   AttractionType `gorm:"foreignKey:AttractionTypeId"`

	// ratings
	Ratings []Rating `gorm:"foreignKey:AttractionId"`

	// attraction reports
	AttractionReports []AttractionReport `gorm:"foreignKey:AttractionId"`

	// user
	UserId uint
	User   User `gorm:"foreignKey:UserId"`

	//comments
	Comments []Comment `gorm:"foreignKey:AttractionId"`
}