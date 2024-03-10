package models

type Host struct {
	ID       int64  `json:"id"`
	Hostname string `json:"hostname"`
	Ip 	 string `json:"ip"`
	CreatedAt string `json:"created_at"`
}