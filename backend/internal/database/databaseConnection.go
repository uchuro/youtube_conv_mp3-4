package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uchuro/youtube_conv_mp3_mp4/config"
)

func DatabseConnection(cfg config.Config) (*sql.DB, error) {
	dbConnection := cfg.Database

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConnection.Username,
		dbConnection.Password,
		dbConnection.Host,
		dbConnection.Port,
		dbConnection.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error in DB CONFIG loading ")
		fmt.Println("Error: ==> ", err)
		return nil, err
	}
	if erro := db.Ping(); erro != nil {
		fmt.Println("DB connection is disconnected ")
		fmt.Println("Error ==> ", erro)
		return nil, err
	}

	return db, nil

}
