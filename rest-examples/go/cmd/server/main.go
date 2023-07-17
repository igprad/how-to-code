package main

import (
	"net/http"

	"github.com/igprad/how-to-code/internal/config"
	"github.com/igprad/how-to-code/internal/user"
	dbcontext "github.com/igprad/how-to-code/pkg/dbctx"
)

func main() {
	// init
	// construct db pointer
	db := dbcontext.InitConnection(config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName)
	// construct controller -> service -> repository
	userApi := user.CreateUserApi(user.NewService(user.NewRepository(db)))

	http.HandleFunc("/users", userApi.GetUsers)
	http.HandleFunc("/user/add", userApi.CreateUser)
	http.HandleFunc("/user/edit/", userApi.EditUser)
	http.HandleFunc("/user/delete/", userApi.DeleteUser)

	http.ListenAndServe(":6969", nil)
}
