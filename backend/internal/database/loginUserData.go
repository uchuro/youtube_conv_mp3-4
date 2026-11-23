package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/uchuro/youtube_conv_mp3_mp4/internal/types"
)

func LoginUserData(usr *types.User, db *sql.DB, r *http.Request) bool {
	query := "INSERT INTO `user` (u_email, u_password) VALUES (?, ?)"

	result, err := db.ExecContext(
		r.Context(),
		query,
		usr.UserEmail,
		usr.UserPassword,
	)
	if err != nil {
		fmt.Println("Error during inserting data ")
		fmt.Println("Error: ==> ", err)
		return false
	}

	fmt.Println(result)
	return true
}
