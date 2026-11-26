package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/uchuro/youtube_conv_mp3_mp4/internal/types"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserData(usr *types.User, db *sql.DB, r *http.Request) (bool, error) {

	var passHash string

	userExitAndPasswQuery := "SELECT u_password FROM user WHERE u_email = ?"

	err := db.QueryRowContext(
		r.Context(),
		userExitAndPasswQuery,
		usr.UserEmail,
	).Scan(&passHash)

	// User already exists
	if err == nil {

		// Check the entered password
		err := bcrypt.CompareHashAndPassword(
			[]byte(passHash),
			[]byte(usr.UserPassword),
		)

		// Password is incorrect
		if err != nil {
			return false, errors.New("invalid credentials")
		}

		// Password is correct
		return true, nil
	}

	// If the error is not because the user doesn't exist
	if !errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("database error: %w", err)
	}

	// User does not exist, create a new user
	createNewUserQuery := "INSERT INTO `user` (u_email, u_password) VALUES (?, ?)"

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(usr.UserPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return false, fmt.Errorf("password hashing failed: %w", err)
	}

	_, err = db.ExecContext(
		r.Context(),
		createNewUserQuery,
		usr.UserEmail,
		hash,
	)

	if err != nil {
		return false, fmt.Errorf("user creation failed: %w", err)
	}

	return true, nil
}
