package models

type SurfSpots struct {
	ID      	uint    `json:"id" gorm:"primary_key"`
	Name    	string 	`json:"name"`
	Image		string 	`json:"image"`
	City 		string 	`json:"city"`
	Latitude  	float64 `json:"latitude"`
	Longitude 	float64 `json:"longitude"`
	Risk 		string 	`json:"risk"`
	Description string 	`json:"description"`
}