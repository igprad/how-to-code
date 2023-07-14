package main

import (
	"net/http"

	"igprad/learn/rest/go/example/controller"
)

func main() {
	http.HandleFunc("/users", controller.GetAllUsers)
	http.HandleFunc("/user/add", controller.CreateUser)
	http.HandleFunc("/user/edit/", controller.EditUser)
	http.HandleFunc("/user/delete/", controller.DeleteUser)

	http.ListenAndServe(":6969", nil)
}
