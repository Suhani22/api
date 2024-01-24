package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID    string
	Name  string
	Email string
}

var contacts []Contact

func main() {
	router := mux.NewRouter()

	contacts = append(contacts, Contact{ID: "1", Name: "Suhani", Email: "suhani@gmail.com"})
	contacts = append(contacts, Contact{ID: "2", Name: "Subhagya", Email: "subhi@yahoo.com"})

	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	//router.HandleFunc("/contacts", AddContact).Methods("POST")
	//router.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	//router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")

	http.Handle("/", router)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contactID := params["id"]
	// params := map[string]string
	for _, contact := range contacts {
		if contact.ID == contactID {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Contact with ID %s not found.", contactID)
}
