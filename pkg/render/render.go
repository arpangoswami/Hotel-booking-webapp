package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/arpangoswami/Hotel-booking-webapp/pkg/config"
	"github.com/arpangoswami/Hotel-booking-webapp/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(appLocal *config.AppConfig) {
	app = appLocal
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

//RenderTemplate renders html/templates inside templates folder
func RenderTemplate(rw http.ResponseWriter, templateTitle string, templateData *models.TemplateData) {

	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()

	}

	temp, ok := templateCache[templateTitle]
	if !ok {
		log.Fatal("Page not found inside template folder")
	}
	buffer := new(bytes.Buffer)
	AddDefaultData(templateData)
	_ = temp.Execute(buffer, templateData)
	_, err := buffer.WriteTo(rw)
	if err != nil {
		fmt.Println("Error writing template to browser ", err.Error())
		return
	}
}

//Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	fmt.Println(pages)
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
