package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Conn *sql.DB
}

func NewMySQL(user, pass, host, port, dbname string) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, dbname,
	)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}
