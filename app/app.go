package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// routes
	router.HandleFunc( "/greet", greet ).Methods(http.MethodGet)
	router.HandleFunc( "/customers", getAllCustomers ).Methods(http.MethodGet)
	router.HandleFunc( "/customers", createCustomer ).Methods(http.MethodPost)
	router.HandleFunc( "/customers/{customer_id:[0-9]+}", getCustomers ).Methods(http.MethodGet)

	// / starting the server
	fmt.Printf("server running on http://localhost:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}


func getCustomers(resp http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	fmt.Fprint(resp, vars["customer_id"])
}

func createCustomer(resp http.ResponseWriter, req *http.Request){
	fmt.Fprint(resp, "post request recieved")
}