package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"odyssey.lms/internal/colors"
)

func GetDatabaseConnection() *sql.DB {

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))
	connString := os.Getenv("CONNECTION_STRING")

	switch dbEngine {
	case "sqlite":
		if connString == "" {
			return connectToDB("sqlite3", "file:lms.db")
		}
		return connectToDB("sqlite3", connString)

	case "mysql":
		return connectToDB("mysql", connString)

	case "postgresql":
		return connectToDB("postgres", connString)

	default:
		return connectToDB("sqlite3", "file:lms.db")
	}

}

func connectToDB(engine string, connString string) *sql.DB {
	CheckDBSettings()

	log.Println(colors.Blue + "[ INFO ] Connecting to database..." + colors.Reset)

	db, err := sql.Open(engine, connString)
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Could not open database connection, check connection string" + colors.Reset)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(colors.RedBold + "[ ERROR ] Could not connect to database" + colors.Reset)
	}

	log.Println(colors.Blue + "[ INFO ] Connected to database" + colors.Reset)

	return db
}
