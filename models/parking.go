package models

type Parking struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Type	string `json:"type"`
	City 	string `json:"city"`
}