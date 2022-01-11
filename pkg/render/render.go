package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/monjurmorshed2022/go_booking/pkg/config"
	"github.com/monjurmorshed2022/go_booking/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

//SetTemplate sets the config for the template package
func SetTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(tempData *models.TemplateDate) *models.TemplateDate {
	return tempData
}

// RenderTemplate renders templates using html/template
func RenderTemplate(rw http.ResponseWriter, temp string, tempData *models.TemplateDate) {
	// get the template catch from config package
	templateCatch := app.TemplateCache

	currentTemplate, ok := templateCatch[temp]

	if !ok {
		log.Fatal("Could not get template from template catch")
	}

	buffer := new(bytes.Buffer)

	tempData = AddDefaultData(tempData)

	_ = currentTemplate.Execute(buffer, tempData)

	_, err := buffer.WriteTo(rw)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

//CreateTemplateCatch creates template catch as a map
func CreateTemplateCatch() (map[string]*template.Template, error) {

	templateCatch := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return templateCatch, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCatch, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return templateCatch, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return templateCatch, err
			}
		}
		templateCatch[name] = templateSet
	}
	return templateCatch, nil
}
