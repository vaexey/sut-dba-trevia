package model

type RegionType struct {
	Id      uint `gorm:"primaryKey"`
	Name    string
	Regions []Region `gorm:"foreignKey:RegionTypeId"`
}