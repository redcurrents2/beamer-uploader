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

    var uploadTypes = map[string]string{
        "1": "yaml",
        "2": "json",
    }
    var destinations = map[string]string{
        "dev": "https://www.withtopic.com/manage/",
        "preprod": "https://www.wiztopic.work/manage/",
        "other": "",
    }

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]map[string]string{}
        data["destinations"] = destinations
        data["uploadTypes"] = uploadTypes
        home.Execute(w, data)
    })

    r.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        switch uploadType := r.FormValue("uploadType"); uploadType {
        case "yaml":
            handleYaml(uploadType, w)
        case "json":
            handleJson(uploadType, w)
        default:
            fmt.Fprintf(w, "Something that should never happened, happened: %s\n", uploadType)
        }
    }).Methods("POST")

    http.ListenAndServe(":8080", r)
}

func handleYaml(ftype string, w http.ResponseWriter) {
    fmt.Fprintf(w, "coucou: %s\n", ftype)
}

func handleJson(ftype string, w http.ResponseWriter) {
    fmt.Fprintf(w, "coucou: %s\n", ftype)
}
