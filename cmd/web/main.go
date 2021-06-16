package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/torenware/sawler-course/config"
	"github.com/torenware/sawler-course/handlers"
	"github.com/torenware/sawler-course/render"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)
	handlers.NewHandlers(&app)

	fmt.Println("Starting at 3000")
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	http.ListenAndServe(":3000", nil)
}
