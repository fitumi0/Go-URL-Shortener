package main

import (
	"gourlshortener/internal/app"
	"log"
	"net/http"
)

func main() {
    app := app.NewApp()
	log.Println("Starting server at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", app.Router))
}
