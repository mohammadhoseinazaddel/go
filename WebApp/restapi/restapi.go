package restapi

import (
	"fmt"
	"main/datalayer"
	"net/http"

	"github.com/gorilla/mux"
)

//Http Get search /api/person/name/{name}
//http POST add or Edit /api/person/add | localhost:8181/api/person/edit

func RunApi(endpoint string, db datalayer.SQLHandler) error {
	r := mux.NewRouter()
	RunApiOnRouter(r, db)
	fmt.Println("Server Started ...")
	return http.ListenAndServe(endpoint, r)
}

func RunApiOnRouter(r *mux.Router, db datalayer.SQLHandler) {
	handler := newPersonRestApiHandler(db)
	apiRouter := r.PathPrefix("/api/person").Subrouter()
	apiRouter.Methods("GET").Path("/name/{search}").HandlerFunc(handler.SearchByName)
	apiRouter.Methods("GET").Path("/AllPeople").HandlerFunc(handler.GetPeople)
	apiRouter.Methods("POST").PathPrefix("/{operator}").HandlerFunc(handler.Operation)
}
