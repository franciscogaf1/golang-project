package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConnection interface {
	SetupPostgresDBConnection(dbUser string, dbPassword string, dbHost string, dbPort string, dbName string) *sql.DB
}

type DBConn struct{}

func NewDBConnection() DBConnection {
	return &DBConn{}
}

func (*DBConn) SetupPostgresDBConnection(dbUser string, dbPassword string, dbHost string, dbPort string, dbName string) *sql.DB {
	connStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	
	return db
}
