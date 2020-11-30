package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DbCreds struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func (d *DbCreds) isValid() bool {
	if len(d.host) == 0 || len(d.port) == 0 || len(d.user) == 0 || len(d.password) == 0 || len(d.dbName) == 0 {
		return false
	}

	return true
}

type GameData struct {
	Date         string   `json:"gameDate"`
	Words        []string `json:"words"`
	Letters      []string `json:"letters"`
	CenterLetter string   `json:"centerLetter"`
}

func (g *GameData) JSON() string {
	data, err := json.Marshal(g)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	var test bool
	var confLoc string

	flag.BoolVar(&test, "t", false, "Scrape data without updating database")
	flag.StringVar(&confLoc, "c", "../../config.yaml", "Location of configuration file holding database credentials")
	flag.Parse()

	if test {
		fmt.Println(Scrape().JSON())
		return
	}

	if _, err := os.Stat(confLoc); err != nil {
		panic(err)
	}

	viper.SetConfigFile(confLoc)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dbCreds := DbCreds{
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		dbName:   viper.GetString("database.dbName"),
	}

	if !dbCreds.isValid() {
		panic("Database credentials are not valid")
	}
	// gameData := Scrape()

}
