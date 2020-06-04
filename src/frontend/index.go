package frontend

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//HandleIndex ... the routing function that implements index page
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/frontend/index.html")
	if err != nil {
		log.Printf("Error while extracting index.html, %s\n", err.Error())
		fmt.Fprintf(w, "Internal error\n")
		return
	}

	tmpl.Execute(w, nil)
}