package db

import (
	"embed"
	"log"
	"os"
	"strings"

	"github.com/pressly/goose/v3"

	"odyssey.lms/internal/colors"
)

//go:embed schema/sqlite/migrations/*.sql
var sqliteMigrations embed.FS

//go:embed schema/postgresql/migrations/*.sql
var postgresqlMigrations embed.FS

//go:embed schema/mysql/migrations/*.sql
var mysqlMigrations embed.FS

func RunMigrations() {

	dbConn := GetDatabaseConnection()

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))

	goose.SetLogger(goose.NopLogger())
	goose.SetTableName("lms_db_version")

	switch dbEngine {
	case "sqlite":
		goose.SetBaseFS(sqliteMigrations)
		if err := goose.SetDialect(string(goose.DialectSQLite3)); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "schema/sqlite/migrations"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	case "mysql":
		goose.SetBaseFS(mysqlMigrations)
		if err := goose.SetDialect("mysql"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "schema/mysql/migrations"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	case "postgresql":
		goose.SetBaseFS(postgresqlMigrations)
		if err := goose.SetDialect("postgres"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "schema/postgresql/migrations"); err != nil {
			log.Println(err)
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}

	default:
		goose.SetBaseFS(sqliteMigrations)
		if err := goose.SetDialect("sqlite3"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Setting migration dialect failed" + colors.Reset)
		}
		if err := goose.Up(dbConn, "schema/sqlite/migrations"); err != nil {
			log.Fatal(colors.RedBold + "[ ERROR ] Database migration failed" + colors.Reset)
		}
	}

	log.Println(colors.Blue + "[ INFO ] Database migrations complete" + colors.Reset)
}
