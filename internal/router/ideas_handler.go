package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/miaababikir/ideas/internal/utils"
)

type Idea struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IdeaRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *App) CreateIdeaHandler(w http.ResponseWriter, r *http.Request) {

	ideaRequest := &IdeaRequest{}

	json.NewDecoder(r.Body).Decode(ideaRequest)

	result, err := app.Db.Exec(
		"INSERT INTO ideas (title, content) VALUES(?, ?)",
		ideaRequest.Title,
		ideaRequest.Content,
	)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create idea")
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create idea")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, id)

}

func (app *App) UpdateIdeaByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	ideaRequest := &IdeaRequest{}

	err := json.NewDecoder(r.Body).Decode(ideaRequest)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to parse request")
		return
	}

	_, err = app.Db.Exec(
		"UPDATE ideas SET title = ?, content = ? WHERE id = ?",
		ideaRequest.Title,
		ideaRequest.Content,
		id,
	)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, id)

}

func (app *App) DeleteIdeaByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := app.Db.Exec(
		"DELETE FROM ideas WHERE id = ?",
		id,
	)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, id)

}

func (app *App) GetIdeaByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	idea, err := app.findById(id)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusNotFound, "Not found")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, idea)

}

func (app *App) GetIdeasHandler(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")

	ideas := []Idea{}

	query := "SELECT * FROM ideas WHERE true = true"
	args := []interface{}{}

	if search != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY created_at DESC"

	rows, err := app.Db.Query(query, args...)

	if err != nil {
		log.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch ideas")
		return
	}

	for rows.Next() {
		idea := Idea{}
		err = rows.Scan(&idea.Id, &idea.Title, &idea.Content, &idea.CreatedAt, &idea.UpdatedAt)
		if err != nil {
			log.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch ideas")
			return
		}
		ideas = append(ideas, idea)
	}

	utils.RespondWithJson(w, http.StatusOK, ideas)
}

func (app *App) findById(id string) (Idea, error) {

	idea := Idea{}

	err := app.Db.QueryRow("SELECT * FROM ideas WHERE id = ?", id).Scan(&idea.Id, &idea.Title, &idea.Content, &idea.CreatedAt, &idea.UpdatedAt)

	if err != nil {
		return idea, err
	}

	return idea, nil

}
