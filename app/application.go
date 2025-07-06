package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig(addr string) Config {
	return Config{
		Addr: addr,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app Application) Run() error {
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler:nil,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("port is listening on",server.Addr)

	return server.ListenAndServe()

}