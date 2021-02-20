package service

import (
	"adrianizen/library.id/internal/config"
	"net/http"
	"text/template"
)

func IndexUIHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(config.RootDirectory + "/html/index.html"))
	tpl.Execute(w, nil)
}
