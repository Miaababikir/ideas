package main

import (
	"fmt"
	"net/http"

	"github.com/miaababikir/ideas/internal/database"
	"github.com/miaababikir/ideas/internal/env"
	"github.com/miaababikir/ideas/internal/router"
)

func main() {

	dbConn := database.Connect(env.GetString("DB_ADDR", "root:root@tcp(127.0.0.1:3306)/ideas?parseTime=true"))

	defer dbConn.Close()

	port := env.GetString("PORT", "8080")

	app := router.App{
		Port: port,
		Db:   dbConn,
	}

	fmt.Printf("Server listening on port %s", app.Port)

	routes := app.RegisterRoutes()

	http.ListenAndServe(":"+app.Port, routes)

}
