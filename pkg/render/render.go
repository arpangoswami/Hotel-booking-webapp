package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//RenderTemplate renders html/templates inside templates folder
func RenderTemplate(rw http.ResponseWriter, templateTitle string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateTitle)
	err := parsedTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing template ", err.Error())
		return
	}

}
