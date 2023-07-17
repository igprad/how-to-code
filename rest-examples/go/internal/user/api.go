package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igprad/how-to-code/internal/response"
)

type UserApi interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	EditUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userApi struct {
	s Service
}

func CreateUserApi(s Service) UserApi {
	return userApi{s}
}

func (u userApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	userDomains := u.s.GetAllUsers()
	userResults := make([]response.UserResponse, 0)
	for _, v := range userDomains {
		userResults = append(userResults, response.UserResponse{v.Name, v.PhoneNumber})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	jsonResponse, _ := json.Marshal(userResults)
	w.Write([]byte(jsonResponse))
}

func (u userApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Create an User")
}

func (u userApi) EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Edit an User")
}

func (u userApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO - Delete an User")
}
