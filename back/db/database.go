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
	Attraction attractionService
	Rating ratingService
	AttractionType attractionTypeService
	Comment *commentService
}

func NewDatabase(Db *gorm.DB) Database{
	return Database{
		Db: Db,
		Role: roleService{Db: Db},
		User: userService{Db: Db},
		Region: regionService{Db: Db},
		Attraction: attractionService{Db: Db},
		Rating: ratingService{Db: Db},
		AttractionType: attractionTypeService{Db: Db},
		Comment: &commentService{Db: Db},
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

type attractionService struct {
	Db *gorm.DB
}

type attractionTypeService struct {
	Db *gorm.DB
}

type ratingService struct {
	Db *gorm.DB
}