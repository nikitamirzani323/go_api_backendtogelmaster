package main

import (
	"log"

	"github.com/nikitamirzani323/go_api_backendtogelmaster/db"
	"github.com/nikitamirzani323/go_api_backendtogelmaster/routes"
)

func main() {
	db.Init()
	app := routes.Init()
	log.Fatal(app.Listen(":7071"))
}
