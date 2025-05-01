package db

import "back/model"

func (rh *roleService) SelectRoleByName(name string) (model.Role, error) {
	var role model.Role
	if err := rh.Db.Where("name = ?", name).First(&role); err != nil {
		return role, err.Error
	}
	return role, nil
}