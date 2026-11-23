package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Env      string       `json:"env"`
	Port     string       `json:"port"`
	Database DbConncetion `json:"database"`
}
type DbConncetion struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func MustLoadConfig() *Config {
	var cfg = Config{}

	// get data from file
	file, err := os.Open("./config/configLocal.json")
	if err != nil {
		fmt.Println("error during the config file")
		fmt.Println("Error: ==> ", err)
	}

	er := json.NewDecoder(file).Decode(&cfg)

	if er != nil {
		fmt.Println("Error during decoding the config json file")
		fmt.Println("Error: ==> ", er)
	}

	return &cfg

}
