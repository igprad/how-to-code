package user

import (
	"database/sql"
	"fmt"
	"igprad/learn/rest/go/example/internal/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbName   = "training"
)

var db *sql.DB

func initConnection() {
	postgresConn := fmt.Sprintf("host=%s port =%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	con, _ := sql.Open("postgres", postgresConn)
	db = con
}

func close() {
	db.Close()
}

func FindAll() []entity.UserEntity {
	initConnection()
	results, _ := db.Query("SELECT * FROM users")
	defer close()

	userResults := make([]entity.UserEntity, 0)

	for results.Next() {
		user := entity.UserEntity{}
		results.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.IdentityNumber, &user.CreatedOn, &user.UpdatedOn)
		userResults = append(userResults, user)
	}

	return userResults
}
