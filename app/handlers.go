package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"fullName" xml:"fullName"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"Zip_code" xml:"Zip_code"`
}


func greet(resp http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(resp, "hello world")
}


func getAllCustomers(resp http.ResponseWriter, req *http.Request)  {

	customers := []Customer{{Name: "Pratik Tiwari", City: "Surat", Zipcode: "395007"},{Name: "Neil Kawde", City: "Pune", Zipcode: "411014"}}
		
	if  req.Header.Get("Content-Type") == "application/json" {
		resp.Header().Add("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(customers)
	} else {
		resp.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(resp).Encode(customers) 
	}
}

