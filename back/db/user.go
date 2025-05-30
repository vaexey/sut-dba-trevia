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

func (uh *userService) SelectUsersWithMostComments(limit int) ([]model.User, error) {
	var results []model.User

	err := uh.Db.
		Table("comments").
		Select("users.id, users.username, users.display_name, COUNT(comments.id) as comment_count").
		Joins("JOIN users ON users.id = comments.user_id").
		Group("users.id").
		Order("comment_count DESC").
		Limit(limit).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (uh *userService) Create(user model.User) (uint, error) {
	result := uh.Db.Create(&user)
	return user.Id, result.Error
}