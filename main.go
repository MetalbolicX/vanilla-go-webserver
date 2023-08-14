package main

import (
	"log"
	"os"

	"github.com/MetalbolicX/vanilla-go-webserver/internal/config"
	"github.com/MetalbolicX/vanilla-go-webserver/internal/routes"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/render"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/server"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/utils"
)

func main() {

	if err := utils.LoaderEnvFile(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("SERVER_PORT")
	DB_URL := os.Getenv("DB_URL")
	DB_MANAGEMENT_SYSTEM := os.Getenv("DB_MANAGEMENT_SYSTEM")
	STATIC_FOLDER := os.Getenv("STATIC_FOLDER")

	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create a template cache")
	}
	app := config.NewAppConfig(tmplCache, false)
	render.NewTemplates(app)

	server := server.NewServer(PORT)
	routes.BindRoutes(server)
	server.SetupStaticFileServer("./"+STATIC_FOLDER, "resources")

	if err := server.SetDBConfig(DB_MANAGEMENT_SYSTEM, DB_URL); err != nil {
		log.Fatal("The database cannot be configurated")
	}

	if err := server.Listen(); err != nil {
		log.Fatal("Server cannot be started")
	}

}
