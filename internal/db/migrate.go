package db

import (
	"embed"
	"log"
	"os"
	"strings"

	"github.com/pressly/goose/v3"

	"odyssey.lms/internal/colors"
)

//go:embed migrations/sqlite/*.sql
var sqliteMigrations embed.FS

//go:embed migrations/postgresql/*.sql
var postgresqlMigrations embed.FS

//go:embed migrations/mysql/*.sql
var mysqlMigrations embed.FS

func RunMigrations() {

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))

	goose.SetLogger(goose.NopLogger())
	goose.SetTableName("lms_db_version")

	switch dbEngine {
	case "sqlite":
		goose.SetBaseFS(sqliteMigrations)
		if err := goose.SetDialect(string(goose.DialectSQLite3)); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "migrations/sqlite"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	case "mysql":
		goose.SetBaseFS(mysqlMigrations)
		if err := goose.SetDialect("mysql"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "migrations/mysql"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	case "postgresql":
		goose.SetBaseFS(postgresqlMigrations)
		if err := goose.SetDialect("postgres"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "migrations/postgresql"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	default:
		goose.SetBaseFS(sqliteMigrations)
		if err := goose.SetDialect("sqlite3"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "migrations/sqlite"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}
	}

	log.Println(colors.Blue + "[ INFO ] Database migrations complete" + colors.Reset)
}
