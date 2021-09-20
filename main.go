package main

import (
	"log"

	"github.com/nikitamirzani323/go-angular/database"
	"github.com/nikitamirzani323/go-angular/routes"
)

func main() {
	database.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7071"))
}
