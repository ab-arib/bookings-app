package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// render html template manual --notes: remember to always set variable start with UpperCase for accessible in other file
func RenderTemplateManual(w http.ResponseWriter, tmpl string) {
	// parse routes page html file and base template file -- // template is use as the base or index html file
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

/** [start] create template caching variant: simple - logic to make the RenderTemplate function more dynamic */

// template cache map
var tc = make(map[string]*template.Template)

func RenderTemplateVariant(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// find template by it's key in the template cache
	var _, inMap = tc[t]
	if inMap {
		// we have the tamplate in template cache
		fmt.Println("using cahced template")
	} else {
		// we don't have template in template cache
		log.Println("creating template and adding to cache")
		err = createTemplateCacheManual(t)
		if err != nil {
			log.Println(err)
		}
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// create new render template cache
func createTemplateCacheManual(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add to template cache
	tc[t] = tmpl

	return nil
}

/** [end] create template caching variant: simple - logic to make the RenderTemplate function more dynamic */
