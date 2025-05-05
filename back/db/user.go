package db

import "back/model"

func (uh *userService) SelectByUsername(username string) (model.User, error){
	var user model.User
	if err := uh.Db.Where("username = ?", username).First(&user); err != nil {
		return user, err.Error
	}
	return user, nil
}

func (uh *userService) SelectById(id uint) (model.User, error) {
	var user model.User
	if err := uh.Db.Where("id = ?", id).First(&user); err != nil {
		return user, err.Error
	}
	return user, nil
}

func (uh *userService) Create(user model.User) (uint, error) {
	result := uh.Db.Create(&user)
	return user.Id, result.Error
}