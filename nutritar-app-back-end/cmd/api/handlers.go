package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// Payload that holds the state of the API service parameter values
	var payload = struct {
		//define the Payload struct fields
		Status string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		//Assign values to the Payload fields
		Status: "active",
		Message: "Food ANDI service up and running",
		Version: "1.0.0",

	}
	// Serialize the Paylod as a JSON object
	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	// Return the Payload as a JSON object to the requestor of the API service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func (app *application) AllFoods(w http.ResponseWriter, r *http.Request) {
	var foods []models.Food;

	kaleRaw := models.Food {
		FoodID: 72119190,
		Name: "Kale raw",
		Description: "Kale, raw",
		GombsCategory: "Green",
		FoodGroup: "Vegetables",
		ANDI: 1000,
	};
	kaleFreshCooked := models.Food {
		FoodID: 72119211,
		Name: "Kale fresh cooked",
		Description: "Kale, fresh, cooked, no added fat",
		GombsCategory: "Green",
		FoodGroup: "Vegetables",
		ANDI: 1000,
	};
	foods = append(foods, kaleRaw)
	foods = append(foods, kaleFreshCooked)

	// Serialize the Foods as a JSON object
	out, err := json.Marshal(foods)
	if err != nil {
		fmt.Println(err)
	}

	// Return the Foods as a JSON object to the requestor of the API service
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
		
}
