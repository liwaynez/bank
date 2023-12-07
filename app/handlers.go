package app

import (
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashay", City: "bhan", Zipcode: "888348"},
		{Name: "Wayor", City: "ATL", Zipcode: "4edd"},
	}
}

if r.Header.Get("Content-Type") == "application/xml" {
w.Header().Add("Content-Type", "applicaton/xml")
xml.NewEncoder(w).Encode(customers)
} else {
w.Header().Add("Content-Type", "application/json")
json.NewEncoder(w).Encode(customers)
}
