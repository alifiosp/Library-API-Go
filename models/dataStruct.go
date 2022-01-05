package models

type Books struct {
	Code            string `json:"Code" gorm:"primary_key"`
	Title           string `json:"Title"`
	Author          string `json:"Author"`
	Genre           string `json:"Genre"`
	PublicationYear string `json:"Publication Year"`
	Placement       string `json:"Placement"`
}

type Users struct {
	Id      string `json:"Id" gorm:"primary_key"`
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Address string `json:"Address"`
	Status  string `json:"Status"`
}
