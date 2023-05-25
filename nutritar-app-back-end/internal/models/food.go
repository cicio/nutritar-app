package models

import "time"



type Food struct {
	ID int 	`json:"id"`
	DataBankID string 	`json:"data_bank_id"`
	FoodName string 	`json:"food_name"`
	FoodDescription string 	`json:"food_description"`
	FoodImage string 	`json:"food_image"`
	CreatedAT time.Time 	`json:"-"`
	UpdatedAT time.Time 	`json:"-"` 	
}
