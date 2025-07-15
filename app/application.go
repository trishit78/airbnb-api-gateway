package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "AuthInGo/config/db"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
	//Store db.Storage
}

func NewConfig() Config {
	port :=config.GetString("PORT",":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
		//Store: *db.NewStorage(),
	}
}

func (app Application) Run() error {

	db,err := dbConfig.SetupDB()
	if err != nil{
		fmt.Println("Error setting up database",err)
		return err
	}

	ur:= repo.NewUserRepository(db)
	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler:router.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("port is listening on",server.Addr)

	return server.ListenAndServe()

}