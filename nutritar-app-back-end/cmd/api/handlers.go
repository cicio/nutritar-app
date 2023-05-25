package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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
		// out, err := json.Marshal(payload)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// // Return the Payload as a JSON object to the requestor of the API service
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(out)
	// Refactored payload serialization with writeJSON
	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) AllFoods(w http.ResponseWriter, r *http.Request) {
	// var foods []models.Food;

	// kaleRaw := models.Food {
	// 	ID: 1,
	// 	DataBankID: "72119190",
	// 	FoodName: "Kale_raw",
	// 	FoodDescription: "Kale, raw",
	// 	FoodImage: "../../sql/freshkale.jpg",
	// 	CreatedAT: time.Now(),
	// 	UpdatedAT: time.Now(),
	// };
	// kaleFreshCooked := models.Food {
	// 	ID: 2,
	// 	DataBankID: "72119211",
	// 	FoodName: "Kale fresh cooked",
	// 	FoodDescription: "Kale, fresh, cooked, no added fat",
	// 	FoodImage: "../../sql/KaleCooked.jpg",
	// 	CreatedAT: time.Now(),
	// 	UpdatedAT: time.Now(),
	// };
	// foods = append(foods, kaleRaw)
	// foods = append(foods, kaleFreshCooked)

	// Instead of Hardcoding the foods, lets us get the values from the Postgres DB
	foods, err := app.DB.AllFoods()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// Serialize the Foods as a JSON object
		// out, err := json.Marshal(foods)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// // Return the Foods as a JSON object to the requestor of the API service
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(out)
	
	//Refactor serialization of JSON object with writeJSON
	_ = app.writeJSON(w, http.StatusOK, foods) 
		
}


func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read a JSON payload
	var requestPayload struct {
		Email		string 		`json:"email"`
		Password 	string		`json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return 
	}


	// validate user against database
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return 
	}
	
	// Check user password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"),http.StatusBadRequest)
		return
	}



	// create a jwt user

	u := jwtUser {
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}

	// Generate Tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w,err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	// w.Write([]byte(tokens.Token))
	//Refactor using writeJSON utility function
	app.writeJSON(w, http.StatusAccepted, tokens )
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	// range over all cookies
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, 
										func(token *jwt.Token)(interface{}, error){
											return []byte(app.JWTSecret), nil
										})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}
			
			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}
			// save user data into u variable
			u := jwtUser {
				ID: user.ID,
				FirstName: user.FirstName,
				LastName: user.LastName,
			}

			// Generate tokenpairs for User
			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating token"), http.StatusUnauthorized)
				return
			}			
			
			//Set a new refresh cookie to send with our response 
			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			//send back response as JSON
			app.writeJSON(w, http.StatusOK, tokenPairs)

		}
	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}