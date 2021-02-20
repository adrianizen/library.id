package service

import (
	"fmt"
	"net/http"
	"text/template"
)

var tplUpsert = template.Must(template.ParseFiles("html/upsert.html"))

func UpsertHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// name := r.FormValue("name")
	// address := r.FormValue("address")

	tplUpsert.Execute(w, nil)
}
