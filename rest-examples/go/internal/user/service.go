package user

import (
	"fmt"
)

func GetAllUsers() []User {
	users := FindAll()
	userDomains := make([]User, 0)
	fmt.Println(users)

	for _, v := range users {
		userDomains = append(userDomains, User{v.Name, v.PhoneNumber, v.IdentityNumber})
	}
	return userDomains
}
