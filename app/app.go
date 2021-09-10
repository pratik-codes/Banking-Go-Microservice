package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratik-codes/Banking-Go-Microservice/domain"
	"github.com/pratik-codes/Banking-Go-Microservice/service"
)

func Start() {
	router := mux.NewRouter()

	//wiring 
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// routes
	router.HandleFunc( "/customers", ch.getAllCustomers ).Methods(http.MethodGet)


	// / starting the server
	fmt.Printf("server running on http://localhost:8000...")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}


