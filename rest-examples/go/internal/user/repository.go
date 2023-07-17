package user

import (
	"github.com/igprad/how-to-code/internal/entity"
	dbcontext "github.com/igprad/how-to-code/pkg/dbctx"
	_ "github.com/lib/pq"
)

type Repository interface {
	FindAll() []entity.UserEntity
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
