package service

import (
	"sanjose/model"
	"sanjose/utils"
)

func GetVerificationForUser(userID string) model.Verification {
	var verification model.Verification
	result := DB.Where("user_id = ?", userID).First(&verification)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return verification
}

func SetVerificationForUser(userID string, verification model.Verification) error {
	verification.UserID = userID
	if verification.Status == "ACCEPTED" {
		verification.IsVerified = true
	} else {
		verification.IsVerified = false
	}
	if DB.Where("user_id = ?", userID).Select("*").Updates(&verification).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New verification created for user with id: " + userID)
		if result := DB.Create(&verification); result.Error != nil {
			return result.Error
		}
	} else {
		utils.SugarLogger.Infoln("Verification for user with id: " + userID + " has been updated!")
	}
	return nil
}
