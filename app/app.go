package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc( "/greet", greet )
	mux.HandleFunc( "/customers", getAllCustomers )

	// / starting the server
	fmt.Printf("server running on http://localhost:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}