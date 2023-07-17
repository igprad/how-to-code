package user

type Service interface {
	GetAllUsers() []User
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
		userDomains = append(userDomains, User{v.Name, v.PhoneNumber, v.IdentityNumber})
	}
	return userDomains
}
