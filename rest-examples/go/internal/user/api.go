package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
			userResults = append(userResults, response.UserResponse{v.UserId, v.Name, v.PhoneNumber})
		}
		sendOkResponse(w, http.StatusOK, userResults)
	} else {
		sendErrResponse(w, http.StatusMethodNotAllowed, "40500001", "Method not allowed.")
	}
}

func (u userApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var createReq request.CreateUserRequest
		json.NewDecoder(r.Body).Decode(&createReq)
		defer r.Body.Close()

		if success := u.s.CreateUser(&createReq); success {
			sendOkResponse(w, http.StatusCreated, 0)
		} else {
			sendErrResponse(w, http.StatusInternalServerError, "5000001", "Failed to create user.")
		}
	} else {
		sendErrResponse(w, http.StatusMethodNotAllowed, "40500001", "Method not allowed.")
	}
}

func (u userApi) EditUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		paths := strings.Split(r.URL.Path, "/")
		userId := paths[len(paths)-1]

		var editUserRequest request.EditUserRequest
		json.NewDecoder(r.Body).Decode(&editUserRequest)
		defer r.Body.Close()

		if success := u.s.EditUser(userId, editUserRequest); success {
			sendOkResponse(w, http.StatusNoContent, 0)
		} else {
			sendErrResponse(w, http.StatusInternalServerError, "5000001", "Failed to update user.")
		}
	} else {
		sendErrResponse(w, http.StatusMethodNotAllowed, "40500001", "Method not allowed.")
	}
}

func (u userApi) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		paths := strings.Split(r.URL.Path, "/")
		userId := paths[len(paths)-1]

		if success := u.s.DeleteUser(userId); success {
			sendOkResponse(w, http.StatusNoContent, 0)
		} else {
			sendErrResponse(w, http.StatusInternalServerError, "5000001", "Failed to delete user.")
		}
	} else {
		sendErrResponse(w, http.StatusMethodNotAllowed, "40500001", "Method not allowed.")
	}
}

func sendOkResponse[T any](w http.ResponseWriter, httpStatus int, body T) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/text")
	if jsonResponse, err := json.Marshal(body); err == nil {
		if len(jsonResponse) > 1 {
			w.Write([]byte(jsonResponse))
		} else {
			w.Write([]byte(nil))
		}
	} else {
		sendErrResponse(w, http.StatusInternalServerError, "5000000", "Failed to parse the response")
	}
}

func sendErrResponse(w http.ResponseWriter, httpStatus int, errCode string, errMessage string) {
	w.WriteHeader(httpStatus)
	errRes := error.ToErrorResponse(errCode, errMessage)
	fmt.Println(errRes)
	errJsonRes, _ := json.Marshal(errRes)
	w.Write([]byte(errJsonRes))
}
