package db

import "back/model"

func (rh *ratingService) SelectAllByUserId(userId uint) ([]model.Rating, error) {
	var ratings []model.Rating
	err := rh.Db.Where("user_id = ?", userId).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (rh *ratingService) SelectAllByAttractionId(attractionId uint) ([]model.Rating, error) {
	var ratings []model.Rating
	err := rh.Db.Where("attraction_id = ?", attractionId).Find(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (rh *ratingService) Create(rating model.Rating) (uint, error) {
	result := rh.Db.Create(&rating)
	return rating.Id, result.Error
}