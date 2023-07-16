package main

import (
	"igprad/learn/rest/go/example/internal/user"
	"net/http"
)

func main() {
	http.HandleFunc("/users", user.GetUsers)
	http.HandleFunc("/user/add", user.CreateUser)
	http.HandleFunc("/user/edit/", user.EditUser)
	http.HandleFunc("/user/delete/", user.DeleteUser)

	http.ListenAndServe(":6969", nil)
}
