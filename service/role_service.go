package service

import (
	"sanjose/model"
	"sanjose/utils"
	"time"
)

func GetRolesForUser(userID string) []string {
	var roles []string
	result := DB.Table("user_role").Where("user_id = ?", userID).Select("role").Find(&roles)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return roles
}

func SetRolesForUser(userID string, roles []string) error {
	DB.Where("user_id = ?", userID).Delete(&model.Role{})
	for _, role := range roles {
		if result := DB.Create(&model.Role{
			UserID:    userID,
			Role:      role,
			CreatedAt: time.Time{},
		}); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
