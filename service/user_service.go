package service

import (
	"sanjose/model"
	"sanjose/utils"
	"strconv"
)

func GetAllUsers() []model.User {
	var users []model.User
	result := DB.Find(&users)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	for i := range users {
		users[i].Roles = GetRolesForUser(users[i].ID)
		users[i].Privacy = GetPrivacyForUser(users[i].ID)
		users[i].School = GetSchoolForUser(users[i].ID)
		users[i].Verification = GetVerificationForUser(users[i].ID)
		users[i].Connections = GetConnectionsForUser(users[i].ID)
	}
	return users
}

func GetUserByID(userID string) model.User {
	var user model.User
	result := DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	user.Roles = GetRolesForUser(user.ID)
	user.Privacy = GetPrivacyForUser(user.ID)
	user.School = GetSchoolForUser(user.ID)
	user.Verification = GetVerificationForUser(user.ID)
	user.Connections = GetConnectionsForUser(user.ID)
	return user
}

func CreateUser(user model.User) error {
	if DB.Where("id = ?", user.ID).Updates(&user).RowsAffected == 0 {
		utils.SugarLogger.Infoln("New user created with id: " + user.ID)
		if result := DB.Create(&user); result.Error != nil {
			return result.Error
		}
		DiscordLogNewUser(user)
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has been updated!")
	}
	utils.SugarLogger.Infoln("Setting (" + strconv.Itoa(len(user.Roles)) + ") roles for user with id: " + user.ID)
	if err := SetRolesForUser(user.ID, user.Roles); err != nil {
		return err
	}
	if user.Privacy.UserID != "" {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has non-empty privacy object, setting privacy in db...")
		if err := SetPrivacyForUser(user.ID, user.Privacy); err != nil {
			return err
		}
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has empty privacy object, nothing to do here!")
	}
	if user.School.UserID != "" {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has non-empty school object, setting school in db...")
		if err := SetSchoolForUser(user.ID, user.School); err != nil {
			return err
		}
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has empty school object, nothing to do here!")
	}
	if user.Verification.UserID != "" {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has non-empty verification object, setting verification in db...")
		if err := SetVerificationForUser(user.ID, user.Verification); err != nil {
			return err
		}
	} else {
		utils.SugarLogger.Infoln("User with id: " + user.ID + " has empty verification object, nothing to do here!")
	}
	if err := SetConnectionsForUser(user.ID, user.Connections); err != nil {
		return err
	}
	return nil
}
