package services

import (
	"fmt"
	"github.com/TheAnswer16/rapper_phrases_generator/database"
	"github.com/TheAnswer16/rapper_phrases_generator/models"
)



func GetPhrase (existentHostID int64) (models.Phrase, error) {

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return models.Phrase{}, err
	}

	defer db.Close()

	var phrase models.Phrase

	
	err = db.QueryRow("SELECT phrases.id_phrase, phrases.id_author, phrases.phrase, authors.author_name, phrases.music, phrases.created_at FROM phrases INNER JOIN authors ON phrases.id_author = authors.id_author WHERE id_phrase NOT IN (SELECT id_phrase FROM host_phrase_relationship WHERE id_host = $1) ORDER BY RANDOM() LIMIT 1", existentHostID).Scan(&phrase.IDPhrase, &phrase.IDAuthor, &phrase.Phrase, &phrase.Author, &phrase.Music, &phrase.CreatedAt)
	if err == sql.ErrNoRows {
		return models.Phrase{}, fmt.Errorf("no rows in result set")
	}
	if err != nil {
		fmt.Println(err)
		return models.Phrase{}, err
	}

	_, err = db.Exec("INSERT INTO host_phrase_relationship (id_host, id_phrase) VALUES ($1, $2)", existentHostID, phrase.IDPhrase)

	if err != nil {
		fmt.Println(err)
		return models.Phrase{}, err
	}

	return phrase, nil


}
