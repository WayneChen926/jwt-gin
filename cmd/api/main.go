package main

import (
	"jwt-gin/config"
)

var app config.AppConfig

func main() {

	//models.ConnectDataBase()
	repo := NewRepo(&app)
	NewHandlers(repo)
	r := routes()
	r.Run(":8080")
}
