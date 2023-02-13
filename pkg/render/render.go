package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/suryasatriah/learn-go/pkg/config"
	"github.com/suryasatriah/learn-go/pkg/model"
)

//var functions = template.FuncMap{}

var app *config.AppConfig

// this function set the config for template page
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData{
	return td
}

// RenderTemplate render the templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {

	var tc map[string]*template.Template
	// create templates chache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateChache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateChache() (map[string]*template.Template, error) {
	// myChache := make(map[string]*template.Template), atau sama aja pake :
	myChache := map[string]*template.Template{}

	// gets all the files starting with *.pages.tmpl
	pages, err := filepath.Glob("./templates/*.pages.tmpl")
	if err != nil {
		return myChache, err
	}

	// range to all files

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myChache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myChache, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myChache, err
			}
			myChache[name] = ts
		}
	}

	return myChache, nil
}
