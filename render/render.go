package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/torenware/sawler-course/config"
)

// Functions to exend templates
var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}

	log.Print("starting create templates")
	pages, err := filepath.Glob("./templates/*.pages.tmpl")
	if err != nil {
		// log.Fatalf("loading page templates failed: %s", err)
		return tCache, err
	}

	for ndx, page := range pages {
		base_path := "./templates/base.layouts.tmpl"
		name := filepath.Base(page)
		log.Printf("%d: %s", ndx, name)
		tmpl, err := template.New(name).Funcs(functions).ParseFiles(base_path, page)
		if err != nil {
			return tCache, err
		}
		tCache[name] = tmpl
	}

	return tCache, nil

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tCache map[string]*template.Template
	name := filepath.Base(tmpl)

	if app.UseCache {
		tCache = app.TemplateCache
	} else {
		tCache, _ = CreateTemplateCache()
	}

	ts, ok := tCache[name]
	if !ok {
		log.Fatalf("Template %s not found", tmpl)
	}

	buf := new(bytes.Buffer)
	data := map[string]interface{}{
		"Title": "Hello World!",
	}
	err := ts.Execute(buf, data)
	if err != nil {
		log.Fatalf("error rendering template: %s", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalf("error writing template: %s", err)
	}

}
