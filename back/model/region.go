package model

type Region struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	RegionTypeId uint
	RegionType   RegionType `gorm:"foreignKey:RegionTypeId"`
	Description  string
	Attractions  []Attraction `gorm:"foreignKey:RegionId"`
}