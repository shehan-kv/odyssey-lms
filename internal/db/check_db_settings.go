package db

import (
	"log"
	"os"
	"strings"

	"odyssey.lms/internal/colors"
)

func CheckDBSettings() {
	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))
	connString := os.Getenv("CONNECTION_STRING")

	switch dbEngine {
	case "sqlite":
		log.Println(colors.Blue + "[ INFO ] SQLite selected" + colors.Reset)
		if connString == "" {
			log.Println(colors.Yellow + "[ WARN ] Database connection string not found" + colors.Reset)
			log.Println(colors.Yellow + "[ WARN ] Using default SQLite database" + colors.Reset)

		} else {
			log.Println(colors.Blue + "[ INFO ] Database connection string found" + colors.Reset)
		}

	case "mysql", "postgresql":
		if dbEngine == "mysql" {
			log.Println(colors.Blue + "[ INFO ] MySQL selected" + colors.Reset)
		} else {
			log.Println(colors.Blue + "[ INFO ] PostgreSQL selected" + colors.Reset)
		}
		if connString == "" {
			log.Fatal(colors.RedBold + "[ ERROR ] Database connection string not found" + colors.Reset)
		}
		log.Println(colors.Blue + "[ INFO ] Database connection string found" + colors.Reset)

	default:
		log.Println(colors.Blue + "[ INFO ] No valid database selected" + colors.Reset)
		log.Println(colors.Blue + "[ INFO ] Using SQLite" + colors.Reset)
	}
}
