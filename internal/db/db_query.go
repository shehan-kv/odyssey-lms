package db

import (
	"database/sql"
	"os"
	"strings"

	"odyssey.lms/internal/db/queries/mysql"
	"odyssey.lms/internal/db/queries/postgresql"
	"odyssey.lms/internal/db/queries/sqlite"
)

var dbConn = GetDatabaseConnection()
var QUERY = GetDBQuery(dbConn)

func GetDBQuery(dbConn *sql.DB) DBQuery {

	dbEngine := strings.ToLower(os.Getenv("DB_VARIANT"))

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
