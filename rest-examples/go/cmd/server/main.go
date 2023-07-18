package main

import (
	"fmt"
	"net/http"

	"github.com/igprad/how-to-code/internal/config"
	"github.com/igprad/how-to-code/internal/user"
	dbcontext "github.com/igprad/how-to-code/pkg/dbctx"
)

func main() {
	/*
	* Init section, if you're coming from Spring Java Realm
	* This is like bean creation ;)
	 */
	db := dbcontext.InitConnection(config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName)
	userApi := user.CreateUserApi(user.NewService(user.NewRepository(db)))

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userApi.GetUsers)
	mux.HandleFunc("/user/add", userApi.CreateUser)
	mux.HandleFunc("/user/edit/", userApi.EditUser)
	mux.HandleFunc("/user/delete/", userApi.DeleteUser)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 6969),
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error run server: ", err)
	}
}
