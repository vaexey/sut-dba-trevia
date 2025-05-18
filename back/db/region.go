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

func (rh *regionService) SelectAllRegionIds(regionId uint) ([]uint, error) {
	var allIds []uint
	var queue []uint = []uint{regionId}

	for len(queue) > 0 {
		currentId := queue[0]
		queue = queue[1:]

		allIds = append(allIds, currentId)
		
		var subregions []model.Region
		if err := rh.Db.Where("parent_region_id = ?", currentId).Find(&subregions).Error; err != nil {
			return nil, err
		}
		
		for _, sub := range subregions {
			queue = append(queue, sub.Id)
		}
	}
	return allIds, nil
}

func (rh *regionService) SelectAllSubregionsIds(regionId uint) ([]uint, error){
	var subregions []model.Region
	if err := rh.Db.Where("parent_region_id = ?", regionId).Find(&subregions).Error; err != nil {
		return nil, err
	}
	var subRegionsIds []uint
	for _, region := range subregions {
		subRegionsIds = append(subRegionsIds, uint(region.Id))
	}
	return subRegionsIds, nil
} 