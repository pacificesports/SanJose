package service

import (
	"sanjose/model"
	"sanjose/utils"
)

func GetSchoolForUser(userID string) model.School {
	var school model.School
	result := DB.Where("user_id = ?", userID).First(&school)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return school
}

func SetSchoolForUser(userID string, school model.School) error {
	school.UserID = userID
	if DB.Where("user_id = ?", userID).Select("*").Updates(&school).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New school created for user with id: " + userID)
		if result := DB.Create(&school); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("School for user with id: " + userID + " has been updated!")
	}
	return nil
}
