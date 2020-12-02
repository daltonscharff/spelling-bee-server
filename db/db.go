package db

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type credentials struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func (d *credentials) connectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.user, d.password, d.dbname)
}

func ReadConfig(configFile string) (credentials, error) {
	creds := credentials{}

	if _, err := os.Stat(configFile); err != nil {
		return creds, err
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return creds, err
	}

	creds.host = viper.GetString("database.host")
	creds.port = viper.GetInt("database.port")
	creds.user = viper.GetString("database.user")
	creds.password = viper.GetString("database.password")
	creds.dbname = viper.GetString("database.dbname")

	return creds, nil
}
