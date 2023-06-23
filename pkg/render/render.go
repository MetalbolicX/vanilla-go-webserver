package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MetalbolicX/vanilla-go-webserver/internal/config"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/types"
	"github.com/MetalbolicX/vanilla-go-webserver/pkg/utils"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates set the config for the temaplate package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate the requested template from the template
// cache, renders it using the provided data and sends
// the output to the user's browser. It handles potential
// errors during template execution and writing the output.
func RenderTemplate(w http.ResponseWriter, tmplFileName string, tmplData *types.TemplateData) {
	// Get the template cache from the app config
	var templateCache map[string]*template.Template
	if app.GetIsUsingCache() {
		templateCache = app.GetTemplateCache()
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	// Get the requested template from cache
	tmpl, templateExists := templateCache[tmplFileName]
	if !templateExists {
		log.Fatal("Could not get template from cache")
	}
	bufferTemplate := new(bytes.Buffer)
	err := tmpl.Execute(bufferTemplate, tmplData)
	if err != nil {
		log.Println(err)
	}
	// Render the template
	_, err = bufferTemplate.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

// CreatetemplateCache is responsible for creating and
// populating a cache of parsed templates. It retrieves
// all files with the extension *.page.tmpl from the
// ./templates directory and parses them using the
// Go html/template package.
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)
	rootPath := utils.GetRootDir()
	// Get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(filepath.Join(rootPath, "templates", "*-page.html"))
	if err != nil {
		return templateCache, err
	}
	// Range through all files ended with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		// Parse the page template file
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}
		// Find layout templates (*.layout.tmpl) in the ./templates directory
		matches, err := filepath.Glob(filepath.Join(rootPath, "templates", "*-layout.html"))
		if err != nil {
			return templateCache, err
		}
		// If layout templates are found, parse them and associate them with the page template
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(filepath.Join(rootPath, "templates", "*-layout.html"))
			if err != nil {
				return templateCache, err
			}
		}
		// Store the template set in the cache map with the template name as the key
		templateCache[name] = templateSet
	}
	return templateCache, nil
}
