package db

import "back/model"

func (rh *roleService) SelectByName(name string) (model.Role, error) {
	var role model.Role
	if err := rh.Db.Where("name = ?", name).First(&role); err != nil {
		return role, err.Error
	}
	return role, nil
}

func (rh *roleService) SelectById(id uint) (model.Role, error) {
	var role model.Role
	if err := rh.Db.Where("id = ?", id).First(&role); err != nil {
		return role, err.Error
	}
	return role, nil
}