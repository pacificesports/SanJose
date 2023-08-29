package service

import (
	"encoding/json"
	"io"
	"net/http"
	"sanjose/model"
	"sanjose/utils"
)

func GetSchoolForUser(userID string) model.School {
	var school model.School
	result := DB.Where("user_id = ?", userID).First(&school)
	if school.SchoolID != "" {
		school.School = FetchSchoolDetails(school.SchoolID)
	}
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

func FetchSchoolDetails(schoolID string) json.RawMessage {
	var responseJson json.RawMessage = []byte("{}")
	mappedService := MatchRoute("schools", "-")
	if mappedService.ID != 0 {
		proxyClient := &http.Client{}
		//proxyRequest, _ := http.NewRequest("GET", "http://localhost"+":"+strconv.Itoa(mappedService.Port)+"/schools/"+schoolID, nil) // Use this when not running in Docker
		proxyRequest, _ := http.NewRequest("GET", mappedService.URL+"/schools/"+schoolID, nil)
		proxyRequest.Header.Set("Request-ID", "-")
		proxyResponse, err := proxyClient.Do(proxyRequest)
		if err != nil {
			utils.SugarLogger.Errorln("Failed to get school information from " + mappedService.Name + ": " + err.Error())
			return responseJson
		}
		defer proxyResponse.Body.Close()
		proxyResponseBodyBytes, _ := io.ReadAll(proxyResponse.Body)
		json.Unmarshal(proxyResponseBodyBytes, &responseJson)
	}
	return responseJson
}
