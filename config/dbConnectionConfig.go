package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type DbConfigUserServicePostgresql struct {
	JsonName         string `json:"JsonName"`
	DatabaseHost     string `json:"DatabaseHost"`
	DatabasePort     string `json:"DatabasePort"`
	DatabaseUser     string `json:"DatabaseUser"`
	DatabasePassword string `json:"DatabasePassword"`
	DatabaseName     string `json:"DatabaseName"`
}

var DbConfigPathLocal string = "../env/configUserServicePostgresql.json"
var ConfigUserDB DbConfigUserServicePostgresql

func loadConfigFile(path string) []byte {
	configDB, err := os.ReadFile(path)
	check(err)
	return configDB
}

func loadConfigDB() (string, error) {
	configData := loadConfigFile(DbConfigPathLocal)
	err := json.Unmarshal(configData, &ConfigUserDB)
	check(err)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", ConfigUserDB.DatabaseHost, ConfigUserDB.DatabasePort, ConfigUserDB.DatabaseUser, ConfigUserDB.DatabasePassword, ConfigUserDB.DatabaseName)
	return connStr, err
}

func ConnectToUserDB() (*sql.DB, error) {
	connStr, err := loadConfigDB()
	check(err)

	db, err := sql.Open("postgres", connStr)
	check(err)

	err = db.Ping()
	check(err)
	return db, err
}
