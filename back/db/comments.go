package db

import (
	"back/model"

	"gorm.io/gorm"
)

type commentService struct {
	Db *gorm.DB
}

func (cs *commentService) SelectAllByAttractionId(attractionId uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := cs.Db.
		Preload("User").
		Where("attraction_id = ?", attractionId).
		Order("id DESC").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (cs *commentService) Create(comment model.Comment) (uint, error) {
	result := cs.Db.Create(&comment)
	return comment.Id, result.Error
}

func (cs *commentService) SelectById(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := cs.Db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
