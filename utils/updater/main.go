package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/daltonscharff/spelling-bee-server/utils/updater/updater"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DbCreds struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (d *DbCreds) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.dbname)
}

func readConfig(configFile string) (DbCreds, error) {
	dbCreds := DbCreds{}

	if _, err := os.Stat(configFile); err != nil {
		return dbCreds, err
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return dbCreds, err
	}

	dbCreds.host = viper.GetString("database.host")
	dbCreds.port = viper.GetInt("database.port")
	dbCreds.user = viper.GetString("database.user")
	dbCreds.password = viper.GetString("database.password")
	dbCreds.dbname = viper.GetString("database.dbname")

	return dbCreds, nil
}

func main() {
	var configFile string

	flag.StringVar(&configFile, "c", "../../config.yaml", "Location of configuration file holding database credentials")
	flag.Parse()

	dbCreds, err := readConfig(configFile)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", dbCreds.ConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Connection String: " + dbCreds.ConnectionString())
		panic(err)
	}

	if err = updater.Update(db); err != nil {
		panic(err)
	}
}
