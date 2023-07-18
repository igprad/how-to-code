package user

import (
	"fmt"

	"github.com/igprad/how-to-code/internal/entity"
	dbcontext "github.com/igprad/how-to-code/pkg/dbctx"
	_ "github.com/lib/pq"
)

type Repository interface {
	FindAll() []entity.UserEntity
	Insert(name string, phoneNumber string, identityNumber string) bool
}

type repository struct {
	db dbcontext.DB
}

func NewRepository(dbctx dbcontext.DB) Repository {
	return repository{dbctx}
}

func (r repository) FindAll() []entity.UserEntity {
	results, _ := r.db.Db.Query("SELECT * FROM users")

	userResults := make([]entity.UserEntity, 0)

	for results.Next() {
		user := entity.UserEntity{}
		results.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.IdentityNumber, &user.CreatedOn, &user.UpdatedOn)
		userResults = append(userResults, user)
	}

	return userResults
}

func (r repository) Insert(name string, phoneNumber string, identityNumber string) bool {
	_, err := r.db.Db.Query("INSERT into users(name, phone_number, identity_number) values('" + name + "','" + phoneNumber + "','" + identityNumber + "')")

	if err == nil {
		return true
	} else {
		fmt.Println("error insert user with err: ", err)
		return false
	}
}
