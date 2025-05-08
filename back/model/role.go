package model

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
	// persmissions

	Users []User `gorm:"foreignKey:RoleId"`
}