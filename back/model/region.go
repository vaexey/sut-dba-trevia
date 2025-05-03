package model

type Region struct {
	Id             uint `gorm:"primaryKey"`
	Name           string
	RegionTypeId   uint
	RegionType     RegionType `gorm:"foreignKey:RegionTypeId"`
	Description    string
	Attractions    []Attraction `gorm:"foreignKey:RegionId"`
	ParentRegionId uint
	ParentRegion   *Region  `gorm:"foreignKey:ParentRegionId"`
	Subregions     []Region `gorm:"foreignKey:ParentRegionId"`
}