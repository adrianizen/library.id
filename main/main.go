package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"runtime"

	"adrianizen/library.id/internal/config"
	"adrianizen/library.id/internal/service"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	config.RootDirectory = dir

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upsert-data", service.UpsertHandler)
	http.ListenAndServe(":"+port, mux)
}
