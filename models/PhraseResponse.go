package models 

type PhraseResponse struct {

	idPhrase int `json:"id_phrase"`
	idAuthor int `json:"id_author"`
	phrase string `json:"phrase"`
	music string `json:"music"`
	created_at string `json:"created_at"`

}