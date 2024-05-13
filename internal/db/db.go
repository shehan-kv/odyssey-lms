package db

import (
	"database/sql"
	"embed"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed schema/sqlite/migrations/*.sql
var sqliteMigrations embed.FS

//go:embed schema/postgresql/migrations/*.sql
var postgresqlMigrations embed.FS

//go:embed schema/mysql/migrations/*.sql
var mysqlMigrations embed.FS

var (
	colorReset   = "\033[0m"
	colorBlue    = "\033[34m"
	colorRedBold = "\033[31;1m"
	colorYellow  = "\033[1;33m"
)

func GetDatabaseConnection() *sql.DB {

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))
	connString := os.Getenv("CONNECTION_STRING")

	switch dbEngine {
	case "sqlite":
		log.Println(colorBlue + "[ INFO ] SQLite selected" + colorReset)
		if connString == "" {
			log.Println(colorYellow + "[ WARN ] Database connection string not found" + colorReset)
			log.Println(colorYellow + "[ WARN ] Using default SQLite database" + colorReset)

			return connectToDB("sqlite3", "file:lms.db")

		} else {
			log.Println(colorBlue + "[ INFO ] Database connection string found" + colorReset)
		}
		log.Println(colorBlue + "[ INFO ] Connecting to SQLite..." + colorReset)

		return connectToDB("sqlite3", connString)

	case "mysql":
		log.Println(colorBlue + "[ INFO ] MySQL selected" + colorReset)
		if connString == "" {
			log.Fatal(colorRedBold + "[ ERROR ] Database connection string not found" + colorReset)

		} else {
			log.Println(colorBlue + "[ INFO ] Database connection string found" + colorReset)
			log.Println(colorBlue + "[ INFO ] Connecting to MySQL..." + colorReset)
		}

		return connectToDB("mysql", connString)

	case "postgresql":
		log.Println(colorBlue + "[ INFO ] PostgreSQL selected" + colorReset)
		if connString == "" {
			log.Fatal(colorRedBold + "[ ERROR ] Database connection string not found" + colorReset)

		} else {
			log.Println(colorBlue + "[ INFO ] Database connection string found" + colorReset)
			log.Println(colorBlue + "[ INFO ] Connecting to PostgreSQL..." + colorReset)
		}

		return connectToDB("postgres", connString)

	default:
		log.Println(colorBlue + "[ INFO ] No valid database selected" + colorReset)
		log.Println(colorBlue + "[ INFO ] Using SQLite" + colorReset)
		log.Println(colorBlue + "[ INFO ] Connecting to SQLite..." + colorReset)

		return connectToDB("sqlite3", "file:lms.db")
	}

}

func RunMigrations() {

	dbConn := GetDatabaseConnection()

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))

	goose.SetLogger(goose.NopLogger())

	switch dbEngine {
	case "sqlite":
		goose.SetBaseFS(sqliteMigrations)
		if err := goose.SetDialect("sqlite3"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Setting migration dialect failed" + colorReset)
		}
		if err := goose.Up(dbConn, "schema/sqlite/migrations"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Database migration failed" + colorReset)
		}

	case "mysql":
		goose.SetBaseFS(mysqlMigrations)
		if err := goose.SetDialect("mysql"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Setting migration dialect failed" + colorReset)
		}
		if err := goose.Up(dbConn, "schema/mysql/migrations"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Database migration failed" + colorReset)
		}

	case "postgresql":
		goose.SetBaseFS(postgresqlMigrations)
		if err := goose.SetDialect("postgres"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Setting migration dialect failed" + colorReset)
		}
		if err := goose.Up(dbConn, "schema/postgresql/migrations"); err != nil {
			log.Fatal(colorRedBold + "[ ERROR ] Database migration failed" + colorReset)
		}

	}

	log.Println(colorBlue + "[ INFO ] Database migrations complete" + colorReset)
}

func connectToDB(engine string, connString string) *sql.DB {

	db, err := sql.Open(engine, connString)
	if err != nil {
		log.Fatal(colorRedBold + "[ ERROR ] Could not open database connection, check connection string" + colorReset)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(colorRedBold + "[ ERROR ] Could not connect to database" + colorReset)
	}

	log.Println(colorBlue + "[ INFO ] Connected to database" + colorReset)

	return db
}
