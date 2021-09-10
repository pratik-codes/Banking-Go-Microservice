package main

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc( "/greet", greet )
	http.HandleFunc( "/customers", getAllCustomers )

	// / starting the server
	fmt.Printf("server running on http://localhost:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}