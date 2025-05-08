package db

import "back/model"

func (ah *attractionTypeService) SelectByName(name string) (model.AttractionType, error) {
	var attractionType model.AttractionType
	if err := ah.Db.Where("name = ?", name).First(&attractionType); err != nil {
		return attractionType, err.Error
	}
	return attractionType, nil
}