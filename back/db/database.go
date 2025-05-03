package db

import (
	"back/model"

	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
	Role roleService
	User userService
	Region regionService
}

func NewDatabase(Db *gorm.DB) Database{
	return Database{
		Db: Db,
		Role: roleService{Db: Db},
		User: userService{Db: Db},
		Region: regionService{Db: Db},
	}
}

func (d *Database) Migrate() error {
	return d.Db.AutoMigrate(
		&model.Attraction{},
		&model.AttractionReport{},
		&model.AttractionType{},
		&model.Comment{},
		&model.CommentReport{},
		&model.Rating{},
		&model.Region{},
		&model.RegionType{},
		&model.Role{},
		&model.User{},
	)
}

type roleService struct {
	Db *gorm.DB
}

type userService struct {
	Db *gorm.DB
}

type regionService struct {
	Db *gorm.DB
}