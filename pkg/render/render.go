package render

import (
	"bytes"
	"github/klyvechen/booking-go/pkg/config"
	"github/klyvechen/booking-go/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders a template in html
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string](*template.Template)
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache
	myCache := map[string]*template.Template{}

	// get all of the files names in *.page.tmpl

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all files ending page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//
//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	// check to see if we already have the template in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		// need to create the template
//		log.Println("creating template and adding to cache")
//		err = createTemplateCached(t)
//		if err != nil {
//			log.Println("Error", err)
//		}
//	} else {
//		// we have the template in the cache
//		log.Println("using cached template")
//	}
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//}
//
//func createTemplateCached(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	tc[t] = tmpl
//	return nil
//}
