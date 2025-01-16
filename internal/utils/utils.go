package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx status code:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errorResponse{Error: message})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	response := Response{
		Code: code,
		Data: payload,
	}

	data, err := json.Marshal(response)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed marshalling payload")
		return
	}

	w.WriteHeader(code)

	_, err = w.Write(data)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed writing response")
		return
	}

}
