/*
	file that start the server and load the server details

*/

package service

import (
	"fmt"
	"net/http"

	"github.com/uchuro/youtube_conv_mp3_mp4/config"
	"github.com/uchuro/youtube_conv_mp3_mp4/internal/api"
	"github.com/uchuro/youtube_conv_mp3_mp4/internal/database"
	"github.com/uchuro/youtube_conv_mp3_mp4/internal/service/policy"
)

func Start() {
	var cfg = config.MustLoadConfig() // create the config variable

	db, err := database.DatabseConnection(*cfg) // db connection loading
	if err != nil {
		fmt.Println("Error during the DB configuration ")
		fmt.Println("Error: ==> ", err)
	}

	defer db.Close()

	handler := http.NewServeMux() //
	// TODO: pass the API here with function

	handler.HandleFunc("/login", api.UserLogin(db))
	handler.HandleFunc("/home",api.UserLink(db))

	/*
		if i want to pass the api function -
			1) handler
			2) db connection
	*/

	server := http.Server{ // server varaible
		Addr:    cfg.Port,
		Handler: policy.EnableCORS(handler),
	}

	// handler is my frontend so frontend is 3000 port so need cors policy apppiled here ->

	server.ListenAndServe() // start linsteing
}
