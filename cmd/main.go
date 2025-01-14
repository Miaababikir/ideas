package main

import (
	"fmt"
	"net/http"
)

type App struct {
	Port string
}

func main() {
	app := App{
		Port: "8080",
	}

	http.HandleFunc("GET /health", app.HealthHandler)

	fmt.Printf("Server listening on port %s", app.Port)

	http.ListenAndServe(":"+app.Port, nil)

}
