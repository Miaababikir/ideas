package main

import "net/http"

func (app *App) HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("OK"))

}
