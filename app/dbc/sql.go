package dbc

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewSqlConnection() *sqlx.DB {
	user := viper.GetString("DB_USERNAME")
	pass := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	name := viper.GetString("DB_NAME")
	sslm := viper.GetString("DB_SSLMODE")

	log.Println(host)

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta", user, pass, host, port, name, sslm)
	log.Println(dsn)

	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	return conn
}
