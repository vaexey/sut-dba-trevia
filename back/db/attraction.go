package db

import (
	"back/model"
	"strings"
)

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

func (ah *attractionService) SelectAllByLocationIdAndCategory(id uint, category string) ([]model.Attraction, error) {
	var attractionsWithFunFact []model.Attraction
	err := ah.Db.Joins("JOIN attraction_types ON attraction_types.id = attractions.attraction_type_id").
    	Where("attractions.region_id = ? AND attraction_types.name = ?", id, category).
    	Preload("AttractionType").
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

func (ah *attractionService) SelectAllByRegionIds(regionIds []uint) ([]model.Attraction, error) {
	var attractions []model.Attraction
	err := ah.Db.Where("region_id IN (?)", regionIds).Find(&attractions).Error
	return attractions, err
}

func (ah *attractionService) SelectAllByRegionIdsAndCategory(regionIds []uint, category string) ([]model.Attraction, error) {
	var attractions []model.Attraction
	err := ah.Db.
	Joins("JOIN attraction_types ON attraction_types.id = attractions.attraction_type_id").
	Where("attractions.region_id IN ? AND LOWER(attraction_types.name) = ?", regionIds, strings.ToLower(category)).
	Find(&attractions).Error
	return attractions, err
}

func (ah *attractionService) SelectAttractionsWithMostComments(numberOfAttractions int) ([]model.Attraction, error) {
	var attractions []model.Attraction
		err := ah.Db.
		Model(&model.Attraction{}).
		Select("attractions.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.attraction_id = attractions.id").
		Group("attractions.id").
		Having("COUNT(comments.id) > 0").
		Order("comment_count DESC").
		Limit(numberOfAttractions).
		Find(&attractions).Error
	if err != nil {
		return nil, err
	}
	return attractions, nil
}

func (ah *attractionService) Create(attraction model.Attraction) (uint, error) {
	result := ah.Db.Create(&attraction)
	return attraction.Id, result.Error
}