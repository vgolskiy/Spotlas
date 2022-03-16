package utils

import (
	"database/sql"
	"endpoint/enteties"
	_ "github.com/lib/pq"
	"log"
)

func GetPGConnection(conf *enteties.PostgresConf) *sql.DB {
	conn, err := sql.Open("postgres", conf.GetConnectionString())
	if err != nil {
		log.Println(err.Error())
	}
	return conn
}