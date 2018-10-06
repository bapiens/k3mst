package main

import (
	"fmt"
	"log"

	"github.com/bapiens/k3mst/internal/diagnostics"
	"github.com/gorilla/mux"

	"net/http"
	"os"
)

func main() {
	log.Print("Starting the application...")

	blPort := os.Getenv("PORT")
	if len(blPort) == 0 {
		log.Fatal("The app port should be set")
	}

	diagPort := os.Getenv("DIAG_PORT")
	if len(diagPort) == 0 {
		log.Fatal("The diagnostics port should be set")
	}
	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	diagRoutes := diagnostics.NewDiagnostics()

	err := http.ListenAndServe(":8585", diagRoutes)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
