package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/igprad/how-to-code/internal/entity"
	dbcontext "github.com/igprad/how-to-code/pkg/dbctx"
	_ "github.com/lib/pq"
)

type Repository interface {
	FindAll() []entity.UserEntity
	Insert(name string, phoneNumber string, identityNumber string) bool
	Edit(userId string, phoneNumber string, identityNumber string) bool
	Delete(userId string) bool
}

type repository struct {
	dbctx dbcontext.DB
}

func NewRepository(dbctx dbcontext.DB) Repository {
	return repository{dbctx}
}

func (r repository) FindAll() []entity.UserEntity {
	results, _ := r.dbctx.Db.Query("SELECT id, uuid, name, phone_number, identity_number, created_on, updated_on FROM users")

	userResults := make([]entity.UserEntity, 0)

	for results.Next() {
		user := entity.UserEntity{}
		results.Scan(&user.Id, &user.UserId, &user.Name, &user.PhoneNumber, &user.IdentityNumber, &user.CreatedOn, &user.UpdatedOn)
		userResults = append(userResults, user)
	}

	return userResults
}

func (r repository) Insert(name string, phoneNumber string, identityNumber string) bool {
	userId := uuid.New()
	_, err := r.dbctx.Db.Query("INSERT into users(uuid, name, phone_number, identity_number) values($1,$2,$3,$4)",
		userId,
		name,
		phoneNumber,
		identityNumber)

	if err == nil {
		return true
	} else {
		fmt.Println("error insert user with err: ", err)
		return false
	}
}

func (r repository) Edit(userId string, phoneNumber string, identityNumber string) bool {
	_, err := r.dbctx.Db.Query("UPDATE users set phone_number = $2, identity_number = $3 WHERE uuid = $1",
		userId,
		phoneNumber,
		identityNumber)

	if err == nil {
		return true
	} else {
		fmt.Println("error edit user with err: ", err)
		return false
	}
}

func (r repository) Delete(userId string) bool {
	_, err := r.dbctx.Db.Query("DELETE from users WHERE uuid = $1", userId)
	if err == nil {
		return true
	} else {
		fmt.Println("error delete user with er: ", err)
		return false
	}
}
