package dataAccess

import "database/sql"

type DataAccess struct {
	dbConnection *sql.DB
}

func NewDataAccess (dbConnection * sql.DB) *DataAccess {
	return &DataAccess{dbConnection: dbConnection}
}