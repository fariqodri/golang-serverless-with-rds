package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	REGION = os.Getenv("REGION")
)

func GetDatabase() *sqlx.DB {
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
		DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)

	//var db *sqlx.DB
	db, err := sqlx.Open("mysql", dnsStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}