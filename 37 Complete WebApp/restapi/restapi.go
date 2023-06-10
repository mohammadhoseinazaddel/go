package restapi

import (
	"TopLearn/WebApp/datalayer"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const STATIC_DIR = "/Static/"
const IMAGE_DIR = "/Images/"

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
	r.Methods("GET").Path("/").HandlerFunc(handler.Index)
	r.Methods("GET").Path("/Post/{PostId}").HandlerFunc(handler.SinglePost)
	r.Methods("GET").Path("/Group/{GroupId}").HandlerFunc(handler.Groups)
	r.Methods("GET").Path("/GetMenu").HandlerFunc(handler.Menu)
	r.Methods("GET").Path("/Login").HandlerFunc(handler.GetLogin)
	r.Methods("POST").Path("/Login").HandlerFunc(handler.PostLogin)

	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("./mainTemplate/"))))
	r.PathPrefix(IMAGE_DIR).Handler(http.StripPrefix(IMAGE_DIR, http.FileServer(http.Dir("./Content/Images"))))
}
