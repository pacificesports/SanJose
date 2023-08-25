package service

import (
	"sanjose/model"
	"sanjose/utils"
)

func GetPrivacyForUser(userID string) model.Privacy {
	var privacy model.Privacy
	result := DB.Where("user_id = ?", userID).First(&privacy)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return privacy
}

func SetPrivacyForUser(userID string, privacy model.Privacy) error {
	privacy.UserID = userID
	if DB.Where("user_id = ?", userID).Select("*").Updates(&privacy).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New privacy created for user with id: " + userID)
		if result := DB.Create(&privacy); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Privacy for user with id: " + userID + " has been updated!")
	}
	return nil
}
