package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("vim-go")
}

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Corey", "Morgantown", "12345"},
		{"Caleb", "Portland", "12345"},
	}
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
