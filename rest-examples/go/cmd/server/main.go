package main

import (
	"net/http"

	"github.com/igprad/how-to-code/internal/user"
)

func main() {
	http.HandleFunc("/users", user.GetUsers)
	http.HandleFunc("/user/add", user.CreateUser)
	http.HandleFunc("/user/edit/", user.EditUser)
	http.HandleFunc("/user/delete/", user.DeleteUser)

	http.ListenAndServe(":6969", nil)
}
