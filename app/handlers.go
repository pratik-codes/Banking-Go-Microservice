package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/pratik-codes/Banking-Go-Microservice/service"
)

type Customer struct {
	Name string `json:"fullName" xml:"fullName"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"Zip_code" xml:"Zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}


func (ch *CustomerHandlers) getAllCustomers(resp http.ResponseWriter, req *http.Request)  {

	customers, _ := ch.service.GetAllCustomer()
		
	if  req.Header.Get("Content-Type") == "application/json" {
		resp.Header().Add("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(customers)
	} else {
		resp.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(resp).Encode(customers) 
	}
}


