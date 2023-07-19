package entity

import "time"

type UserEntity struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	IdentityNumber string    `json:"identity_number"`
	CreatedOn      time.Time `json:"created_on"`
	UpdatedOn      time.Time `json:"updated_on"`
	UserId         string    `json:"uuid"`
}
