package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	error "github.com/igprad/how-to-code/internal/errors"
	"github.com/igprad/how-to-code/internal/request"
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
	if r.Method == "GET" {
		userDomains := u.s.GetAllUsers()
		userResults := make([]response.UserResponse, 0)
		for _, v := range userDomains {
			userResults = append(userResults, response.UserResponse{v.Name, v.PhoneNumber})
		}
		sendOkResponse(w, userResults)
	} else {
		throwNotAllowed(w)
	}
}

func (u userApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var createReq request.CreateUserRequest
		json.NewDecoder(r.Body).Decode(&createReq)
		defer r.Body.Close()

		if success := u.s.CreateUser(&createReq); success {
			sendOkResponse(w, "success")
		} else {
			sendOkResponse(w, "failure")
		}
	} else {
		throwNotAllowed(w)
	}
}

func (u userApi) EditUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		fmt.Println("TODO - Edit an User")
	} else {
		throwNotAllowed(w)
	}
}

func (u userApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		fmt.Println("TODO - Delete an User")
	} else {
		throwNotAllowed(w)
	}
}

func sendOkResponse[T any](w http.ResponseWriter, body T) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	jsonResponse, _ := json.Marshal(body)
	w.Write([]byte(jsonResponse))
}

func throwNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	errRes := error.ToErrorResponse("405", "Method not allowed.")
	fmt.Println(errRes)
	errJsonRes, _ := json.Marshal(errRes)
	w.Write([]byte(errJsonRes))
}
