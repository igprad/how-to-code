package dbcontext

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Db *sql.DB
}

func InitConnection(host string, port int, user string, password string, dbName string) (db DB) {
	postgresConn := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	con, _ := sql.Open("postgres", postgresConn)
	return DB{con}
}

func (db *DB) CloseConnection() {
	db.Db.Close()
}
