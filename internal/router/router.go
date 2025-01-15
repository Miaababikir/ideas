package router

import "net/http"

type App struct {
	Port string
}

func (app *App) RegisterRoutes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/health", app.HealthHandler)

	return router
}
