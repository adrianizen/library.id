package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"adrianizen/library.id/internal/config"
	"adrianizen/library.id/internal/service"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")

	fmt.Println(dir)
	config.RootDirectory = dir

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", service.IndexUIHandler)
	mux.HandleFunc("/list", service.ListUIHandler)
	mux.HandleFunc("/upsert-form", service.UpsertUIHandler)
	mux.HandleFunc("/upsert-handle", service.UpsertHandler)
	http.ListenAndServe(":"+port, mux)
}
