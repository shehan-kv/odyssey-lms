package db

import (
	"os"
	"strings"

	"odyssey.lms/internal/db/queries/mysql"
	"odyssey.lms/internal/db/queries/postgresql"
	"odyssey.lms/internal/db/queries/sqlite"
)

func GetDBQuery() DBQuery {

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))
	dbConn := GetDatabaseConnection()

	switch dbEngine {
	case "sqlite":
		return sqlite.New(dbConn)

	case "mysql":
		return mysql.New(dbConn)

	case "postgresql":
		return postgresql.New(dbConn)

	default:
		return sqlite.New(dbConn)
	}
}
