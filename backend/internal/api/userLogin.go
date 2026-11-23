package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uchuro/youtube_conv_mp3_mp4/internal/database"
	"github.com/uchuro/youtube_conv_mp3_mp4/internal/types"
)

/*
	/login API
	this API data reverify with DB
	if successfully need to redirect with
	/home API


*/

func UserLogin(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check data
		var user types.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Error during the Deocding the login data ")
			fmt.Println("Error: ==> ", err)
		}
		// verify logic
		if len(user.UserEmail) == 0 || len(user.UserPassword) < 3 {
			resError := types.ErrorResponse{
				Success:      false,
				Message:      "login or password are maybe not valid",
				EffectedData: "not any data",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(&resError)
			return
		}
		// pass to db
		if isItSave := database.LoginUserData(&user, db, r); isItSave == true {
			fmt.Println("Data is saved ")
		} else {
			fmt.Println("data is not saved ")
		}
		// return response
	}

}
