package main

import "log"

import "net/http"

func main() {
	log.Print("Ko sta")	

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
		
}

