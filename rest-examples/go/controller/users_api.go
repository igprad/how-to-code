package controller

import (
	"fmt"
	"net/http"
)

type User struct {
	Name        string
	PhoneNumber string
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Get All Users From DB")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Create an User")
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Edit an User")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Delete an User")
}
