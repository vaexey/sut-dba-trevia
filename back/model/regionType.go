package model

type RegionType struct {
	Id      uint `gorm:"primaryKey"`
	Type    string
	Regions []Region `gorm:"foreignKey:RegionTypeId"`
}