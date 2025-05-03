package db

import "back/model"

func (rh *regionService) SelectById(id uint) (model.Region, error) {
	var region model.Region
	if err := rh.Db.Where("id = ?", id).First(&region); err != nil {
		return region, err.Error
	}
	return region, nil
}