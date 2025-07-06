package main

import (
	"AuthInGo/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	err:=godotenv.Load(".env")
	if err != nil{
		log.Fatalln("Error loading in env file",err)
	}

	port:= os.Getenv("PORT")
	cfg:= app.NewConfig(port)
	app:= app.NewApplication(cfg)

	app.Run()
}