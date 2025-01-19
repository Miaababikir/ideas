package router

import (
	"html/template"
	"net/http"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("ui/html/index.html"))

	template.Execute(w, nil)

}
