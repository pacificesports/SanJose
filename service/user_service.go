package service

import (
	"sanjose/model"
	"sanjose/utils"
)

func GetAllUsers() []model.User {
	var users []model.User
	result := DB.Find(&users)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}

	return users
}
