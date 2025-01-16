package router

import (
	"database/sql"
	"net/http"
)

type App struct {
	Port string
	Db   *sql.DB
}

func (app *App) RegisterRoutes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("GET /health", app.HealthHandler)

	router.HandleFunc("GET /ideas", app.GetIdeasHandler)
	router.HandleFunc("POST /ideas", app.CreateIdeaHandler)
	router.HandleFunc("GET /ideas/{id}", app.GetIdeaByIdHandler)

	return router
}
