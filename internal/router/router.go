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

	router.HandleFunc("/", app.Home)

	router.HandleFunc("GET /api/health", app.HealthHandler)

	router.HandleFunc("GET /api/ideas", app.GetIdeasHandler)
	router.HandleFunc("POST /api/ideas", app.CreateIdeaHandler)
	router.HandleFunc("GET /api/ideas/{id}", app.GetIdeaByIdHandler)
	router.HandleFunc("PUT /api/ideas/{id}", app.UpdateIdeaByIdHandler)
	router.HandleFunc("DELETE /api/ideas/{id}", app.DeleteIdeaByIdHandler)

	return router
}
