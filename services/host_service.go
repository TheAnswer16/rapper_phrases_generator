package services 

import (
	"fmt"
	"github.com/TheAnswer16/rapper_phrases_generator/database"
	"github.com/TheAnswer16/rapper_phrases_generator/models"
	"time"
	"database/sql"
)

func SaveClientInfo (clientInfo models.ClientInfo) (int64, error) {

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer db.Close()

	// Insert the client info into the database
	var lastInsertID int64
	err = db.QueryRow("INSERT INTO hosts (hostname, ip, created_at) VALUES ($1, $2, $3) RETURNING id_host", clientInfo.Hostname, clientInfo.IP, time.Now()).Scan(&lastInsertID)

	if err != nil {
		fmt.Println(err)
		return 0, err;
	}

	// Use the last inserted ID
	fmt.Println("Last inserted ID:", lastInsertID)

	return lastInsertID, nil
}

func VerifyExistentHost (hostname string, ip string) (int64, error) {

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer db.Close()

	// Check if the host already exists
	var existentHostID int64
	err = db.QueryRow("SELECT id_host FROM hosts WHERE hostname = $1 AND ip = $2", hostname, ip).Scan(&existentHostID)

	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return existentHostID, nil
}