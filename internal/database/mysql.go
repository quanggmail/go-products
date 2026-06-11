package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {

	dsn := "root:root@tcp(mysql-local:3306)/productdb"

	return sql.Open("mysql", dsn)
}