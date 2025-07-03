package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	user := "root"
	pass := "root"
	host := "127.0.0.1"
	port := "3306"
	name := "taskdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	return sql.Open("mysql", dsn)
}
