package user

import "github.com/igprad/how-to-code/internal/request"

type Service interface {
	GetAllUsers() []User
	CreateUser(ur *request.CreateUserRequest) bool
	EditUser(userId string, er request.EditUserRequest) bool
	DeleteUser(userId string) bool
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) GetAllUsers() []User {
	users := s.r.FindAll()
	userDomains := make([]User, 0)

	for _, v := range users {
		userDomains = append(userDomains, User{v.UserId, v.Name, v.PhoneNumber, v.IdentityNumber})
	}
	return userDomains
}

func (s service) CreateUser(ur *request.CreateUserRequest) bool {
	return s.r.Insert(ur.Name, ur.PhoneNumber, ur.IdentityNumber)

}

func (s service) EditUser(userId string, er request.EditUserRequest) bool {
	return s.r.Edit(userId, er.PhoneNumber, er.IdentityNumber)
}

func (s service) DeleteUser(userId string) bool {
	return s.r.Delete(userId)
}
