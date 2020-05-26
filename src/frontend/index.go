package frontend

import (
	"net/http"
	"html/template"
	"fmt"
	"log"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/frontend/index.html")
	if err != nil {
		log.Printf("Error while extracting index.html, %s\n", err.Error())
		fmt.Fprintf(w, "Internal error\n")
		return
	}
	if r.Method != http.MethodPost {
		log.Println("Not posting")
		tmpl.Execute(w, nil)
	} else {
		handle := r.FormValue("handle")
		log.Println("handle is:",handle)
		http.Redirect(w,r,"/userstats/" + handle,http.StatusFound)
	}
}
