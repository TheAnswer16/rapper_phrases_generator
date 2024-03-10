package models 

type Phrase struct {

	IDPhrase int `json:"id_phrase"`
	IDAuthor int `json:"id_author"`
	Phrase string `json:"phrase"`
	Author string `json:"author"`
	Music string `json:"music"`
	CreatedAt string `json:"created_at"`

}