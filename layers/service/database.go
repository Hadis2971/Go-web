package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Hadis2971/go_web/util"
	"github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() *sql.DB {

	envConfig := util.GetEnvConfig("DB_USER", "DB_PASSWORD", "DB_NET", "DB_Addr", "DB_NAME")

	cfg := mysql.Config{
		User:   envConfig["DB_USER"],
		Passwd: envConfig["DB_PASSWORD"],
		Net:    envConfig["DB_NET"],
		Addr:   envConfig["DB_Addr"],
		DBName: envConfig["DB_NAME"],
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	} // These are good places to do a fatal. No point starting if the database isn't ready.

	fmt.Println("Database Connected")

	return db
}
