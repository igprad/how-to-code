package request

type CreateUserRequest struct {
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	IdentityNumber string `json:"identity_number"`
}

type EditUserRequest struct {
	PhoneNumber    string `json:"phone_number"`
	IdentityNumber string `json:"identity_number"`
}
