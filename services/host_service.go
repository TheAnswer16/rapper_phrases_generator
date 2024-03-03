package services 

import (
	"fmt"
	"github.com/TheAnswer16/rapper_phrases_generator/database"
	"github.com/TheAnswer16/rapper_phrases_generator/models"
	"time"
)

func SaveClientInfo(clientInfo models.ClientInfo) {

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// Insert the client info into the database
	_, err = db.Exec("INSERT INTO hosts (hostname, ip, created_at) VALUES ($1, $2, $3)", clientInfo.Hostname, clientInfo.IP, time.Now())
	fmt.Println("Client info saved")
	if err != nil {
		fmt.Println(err)
		return
	}
}