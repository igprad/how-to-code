package response

type UserResponse struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
