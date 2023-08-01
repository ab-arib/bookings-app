package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ab-arib/bookings-app/pkg/config"
	"github.com/ab-arib/bookings-app/pkg/models"
)

/** render html template with template caching  to make it more dynamic and flexible*/

var app *config.AppConfig

// set the config for templates package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// use to sets default template data value if necessary
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// render template html
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template cache
	t, ok := tc[tmpl]
	if !ok {
		// kill the process
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer) // buffer content is bytes used to provide more clear error
	td = AddDefaultData(td)  // set template data default
	//!notes: indirectly execute the map value using buffer it will provide an error from that map value itself
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writting template to browser", err)
	}
}

// create template cache to store html to render
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		// get name file for map key
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// get all of the layout files
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// implement layout to html templates
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// add to template cache
		myCache[name] = ts
	}

	return myCache, nil
}
