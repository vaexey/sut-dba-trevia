package db

import "back/model"

func (rh *regionService) SelectById(id uint) (model.Region, error) {
	var region model.Region
	if err := rh.Db.Where("id = ?", id).First(&region); err != nil {
		return region, err.Error
	}
	return region, nil
}

func (rh *regionService) SelectByNameFragment(nameFragment string) ([]model.Region, error) {
	var regions []model.Region
	pattern := nameFragment + "%"
	err := rh.Db.Preload("RegionType").Where("name ILIKE ?", pattern).Find(&regions).Error
	if err != nil {
		return nil, err
	}
	return regions, nil
}