package db

import (
	"back/model"

	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
	
}

func NewDatabase(Db *gorm.DB) Database{
	return Database{
		Db: Db,
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