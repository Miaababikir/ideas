package router

import (
	"net/http"

	"github.com/miaababikir/ideas/internal/utils"
)

func (app *App) HealthHandler(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithJson(w, http.StatusOK, struct{ status string }{status: "ok"})

}
