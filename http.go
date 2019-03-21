package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	home := template.Must(template.ParseFiles("views/homepage.tmpl"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string][]string)
		data["destinations"] = []string{
			"dev",
			"prod",
		}
		data["uploadTypes"] = []string{
			"yaml",
			"json",
		}
		
		home.Execute(w, data)
	})

	r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("uploadType") == "yaml" {
			//Do stuff
			return
		}
		if r.FormValue("uploadType") == "json" {
			//Do stuff
			return
		}
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
