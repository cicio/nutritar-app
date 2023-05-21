package models



type Food struct {
	FoodID int `json:"foodID"`
	Name string `json:"name"`
	Description string `json:"description"`
	GombsCategory string `json:"gombsCategory"`
	FoodGroup string `json:"foodGroup"`
	ANDI int `json:"ANDI"`
}
