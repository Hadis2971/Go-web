package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)


func ConnectToDatabase () *sql.DB {
	cfg := mysql.Config{
        User:   "go_web",
        Passwd: "password",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "go_web",
    }

	db, err := sql.Open("mysql", cfg.FormatDSN());

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping();

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected");

	return db;
}