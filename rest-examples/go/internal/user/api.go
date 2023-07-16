package user

import (
	"encoding/json"
	"fmt"
	"igprad/learn/rest/go/example/internal/response"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userDomains := GetAllUsers()
	userResults := make([]response.UserResponse, 0)
	for _, v := range userDomains {
		userResults = append(userResults, response.UserResponse{v.Name, v.PhoneNumber})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	jsonResponse, _ := json.Marshal(userResults)
	w.Write([]byte(jsonResponse))
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
