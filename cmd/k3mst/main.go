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

	possibleErrors := make(chan error, 2) //( 2 is beacuse there are 2 servers listnening ot two ports)

	go func() {
		log.Print("Application is in state of preparation to serve...")

		server := &http.Server{
			Addr:    ":" + blPort,
			Handler: router,
		}

		err := server.ListenAndServe()
		//server.Shutdown()
		if err != nil {
			possibleErrors <- err
		}
	}()

	log.Print("Diagnostics are in state of preparation to serve...")
	diagRoutes := diagnostics.NewDiagnostics()

	diagserver := &http.Server{
		Addr:    ":" + diagPort,
		Handler: diagRoutes,
	}

	err := diagserver.ListenAndServe()
	if err != nil {
		possibleErrors <- err
	}

	select {
	case err := <-possibleErrors:
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("The hello handler was called")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
