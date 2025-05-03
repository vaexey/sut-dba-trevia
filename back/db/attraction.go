package db

import "back/model"

func (ah *attractionService) SelectById(id uint) (model.Attraction, error) {
	var attraction model.Attraction
	if err := ah.Db.Where("id = ?", id).First(&attraction); err != nil {
		return attraction, err.Error
	}
	return attraction, nil
}

func (ah *attractionService) SelectAllByLocationId(id uint) ([]model.Attraction, error) {
	var attractionsWithFunFact []model.Attraction
	err := ah.Db.Where("region_id = ?", id).
		Find(&attractionsWithFunFact).Error

	if err != nil {
		return nil, err
	}
	return attractionsWithFunFact, nil
}

func (ah *attractionService) SelectAllWithFunFact() ([]model.Attraction, error) {
	var attractionsWithFunFact []model.Attraction
	err := ah.Db.Where("fun_fact IS NOT NULL AND fun_fact != ''").
		Find(&attractionsWithFunFact).Error

	if err != nil {
		return nil, err
	}
	return attractionsWithFunFact, nil
}

func (ah *attractionService) Create(attraction model.Attraction) (uint, error) {
	result := ah.Db.Create(&attraction)
	return attraction.Id, result.Error
}