package controllers

import (
	"net/http"

	"github.com/TheAnswer16/rapper_phrases_generator/services"
	"github.com/TheAnswer16/rapper_phrases_generator/models"
	"github.com/TheAnswer16/rapper_phrases_generator/utils"
)


func GetPhrase(w http.ResponseWriter, r *http.Request) {

	var clientInfo models.ClientInfo

	clientInfo.Hostname = r.Host
	clientInfo.IP = r.RemoteAddr

	existentHostID, err := services.VerifyExistentHost(clientInfo.Hostname, clientInfo.IP)

	if err != nil {
		utils.ErrorHandler(w, "Error verifying existent host", http.StatusInternalServerError, err)
		return
	}

	if existentHostID < 1 {
		lastInsertID, err := services.SaveClientInfo(clientInfo)
		existentHostID = lastInsertID
		if err != nil {
			utils.ErrorHandler(w, "Error saving client info", http.StatusInternalServerError, err)
			return
		}
		
	}
	
	phrase, err := services.GetPhrase(existentHostID)

	if err != nil {
		utils.ErrorHandler(w, "Error getting phrase", http.StatusInternalServerError, err)
		return
	}

	utils.Responser(w, phrase)

	
}




