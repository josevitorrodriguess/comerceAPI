package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "12345"
	dbname   = "comerce"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL database:", dbname)

	return db, nil
}