package router

import (
	"encoding/json"
	"fmt"
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

type CreateIdeaRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (app *App) CreateIdeaHandler(w http.ResponseWriter, r *http.Request) {

	createIdeaRequest := &CreateIdeaRequest{}

	json.NewDecoder(r.Body).Decode(createIdeaRequest)

	result, err := app.Db.Exec(
		"INSERT INTO ideas (title, content) VALUES(?, ?)",
		createIdeaRequest.Title,
		createIdeaRequest.Content,
	)

	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create idea")
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create idea")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, id)

}

func (app *App) GetIdeaByIdHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	fmt.Println(id)

	idea := Idea{}

	err := app.Db.QueryRow("SELECT * FROM ideas WHERE id = ?", id).Scan(&idea.Id, &idea.Title, &idea.Content, &idea.CreatedAt, &idea.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch idea")
		return
	}

	utils.RespondWithJson(w, http.StatusOK, idea)

}

func (app *App) GetIdeasHandler(w http.ResponseWriter, r *http.Request) {

	ideas := []Idea{}

	rows, err := app.Db.Query("SELECT * FROM ideas ORDER BY created_at DESC")

	if err != nil {
		fmt.Println(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch ideas")
		return
	}

	for rows.Next() {
		idea := Idea{}
		err = rows.Scan(&idea.Id, &idea.Title, &idea.Content, &idea.CreatedAt, &idea.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch ideas")
			return
		}
		ideas = append(ideas, idea)
	}

	utils.RespondWithJson(w, http.StatusOK, ideas)
}
