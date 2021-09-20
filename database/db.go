package database

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/nikitamirzani323/go-angular/helpers"
)

var db *sql.DB
var err error

func Init() {
	err := godotenv.Load()
	helpers.ErrorCheck(err)

	dbUser := os.Getenv("DB1_USER")
	dbPass := os.Getenv("DB1_PASS")
	dbHost := os.Getenv("DB1_HOST")
	dbPort := os.Getenv("DB1_PORT")
	dbName := os.Getenv("DB1_NAME")

	conString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err = sql.Open("mysql", conString)
	helpers.ErrorCheck(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	err = db.Ping()
	helpers.ErrorCheck(err)
}
func CreateCon() *sql.DB {
	return db
}
