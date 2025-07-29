package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "embed"

	"github.com/google/uuid"
)

//go:embed page.html
var page string

func main() {
	id := uuid.New()
	log.Println("uuid:" + id.String())

	http.HandleFunc("/", serveHello)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}
}

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
