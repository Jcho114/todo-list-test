package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // Check path here
		http.NotFound(w, r)
		return
	}

	log.Println("Inside HomePageHandler")
	tmplFile := "./modules/templates/home.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	err = tmpl.Execute(w, "")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
}
