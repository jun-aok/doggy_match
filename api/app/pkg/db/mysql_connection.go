package db

import (
	"app/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct{}

var db *sql.DB = &sql.DB{}
var singletonInstance *MySQLConnection = newMySQLDbConnection()

func newMySQLDbConnection() *MySQLConnection {
	config := config.NewConfig()
	d, err := sql.Open("mysql", config.MySqlConnectionString)
	if err != nil {
		panic(err.Error())
	}
	db = d
	return &MySQLConnection{}
}

func GetMySQLConnection() *MySQLConnection {
	return singletonInstance
}

func (c MySQLConnection) GetDbConnection() *sql.DB {
	return db
}

func (c MySQLConnection) CloseDbConnection() {
	// https://yaruki-strong-zero.hatenablog.jp/entry/go_sql_open_close
	// この辺を参考にした実装
	db.Close()
}
