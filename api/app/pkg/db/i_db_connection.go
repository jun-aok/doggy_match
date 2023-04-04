package db

import (
	"database/sql"
)

type IDbConnection interface {
	GetDbConnection() *sql.DB
	CloseDbConnection()
}
