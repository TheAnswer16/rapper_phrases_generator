package controllers

import (
	"net/http"

	"github.com/TheAnswer16/rapper_phrases_generator/services"
	"github.com/TheAnswer16/rapper_phrases_generator/models"
)

func GetPhrase(w http.ResponseWriter, r *http.Request) {

	var clientInfo models.ClientInfo

	clientInfo.Hostname = r.Host
	clientInfo.IP = r.RemoteAddr

 	services.SaveClientInfo(clientInfo)

}
