package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"main/datalayer"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type PersonRestApiHandler struct {
	dbhandler datalayer.SQLHandler
}

func newPersonRestApiHandler(db datalayer.SQLHandler) *PersonRestApiHandler {
	return &PersonRestApiHandler{
		dbhandler: db,
	}
}

func (handler PersonRestApiHandler) SearchByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	search, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Search not Found.")
	}

	person, err := handler.dbhandler.GetPersonByName(search)
	if err != nil {
		fmt.Println("person not found!")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "person by name %s not Found.", search)
	}

	json.NewEncoder(w).Encode(person)
	return
}

func (handler PersonRestApiHandler) GetPeople(w http.ResponseWriter, r *http.Request) {

	People, err := handler.dbhandler.GetPeople()
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(People)
	return
}

func (handler PersonRestApiHandler) Operation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator, ok := vars["operator"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Search not Found.")
	}

	var person datalayer.Person
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not Decode request body %v", err)
	}

	switch strings.ToLower(operator) {
	case "add":
		err := handler.dbhandler.AddPeople(person)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not Add Person in to DataBase %v", err)
		}
	case "edit":
		name := r.RequestURI[len("/api/person/edit/"):]
		fmt.Printf("Edit Requested for person by Name %s", name)
		err := handler.dbhandler.UpdatePerson(person)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not Update Person in DataBase %v", err)
		}
	}
}
