package service

import (
	"sanjose/model"
	"sanjose/utils"
)

func GetConnectionsForUser(userID string) []model.Connection {
	var connections []model.Connection
	result := DB.Where("user_id = ?", userID).Find(&connections)
	if result.Error != nil {
		utils.SugarLogger.Errorln(result.Error.Error())
	}
	return connections
}

func SetConnectionsForUser(userID string, connections []model.Connection) error {
	DB.Where("user_id = ?", userID).Delete(&model.Connection{})
	for _, connection := range connections {
		connection.UserID = userID
		if result := DB.Create(&connection); result.Error != nil {
			return result.Error
		}
	}
	return nil
}
