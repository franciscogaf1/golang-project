package repository

import (
	"database/sql"
	"level7/go/src/app/infrastructure"
)

var dbHost string = "postgres"
var dbPort string = "5432"
var dbName string = "postgres"
var dbUser string = "postgres"
var dbPassword string = "postgres"

var dbConn = infrastructure.NewDBConnection().SetupPostgresDBConnection(dbUser, dbPassword, dbHost, dbPort, dbName)

func GetDBConn() *sql.DB {
	return dbConn
}