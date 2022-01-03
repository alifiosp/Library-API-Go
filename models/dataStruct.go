package models

type Books struct {
	Code            string `json:"Code" gorm:"primary_key"`
	Title           string `json:"Title"`
	Author          string `json:"Author"`
	Genre           string `json:"Genre"`
	PublicationYear string `json:"Publication Year"`
}
