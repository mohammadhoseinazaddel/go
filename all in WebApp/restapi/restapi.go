package restapi

import (
	"fmt"
	"log"
	"main/datalayer"
	"main/security/jwt"
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
	ManageSiteRoutes(r, db)
	ManageAdminRoutes(r, db)

	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("./mainTemplate/"))))
	r.PathPrefix(IMAGE_DIR).Handler(http.StripPrefix(IMAGE_DIR, http.FileServer(http.Dir("./Content/Images"))))
}

func ManageSiteRoutes(r *mux.Router, db datalayer.SQLHandler) {
	handler := newPersonRestApiHandler(db)
	r.Methods("GET").Path("/").HandlerFunc(handler.Index)
	r.Methods("GET").Path("/Post/{PostId}").HandlerFunc(handler.SinglePost)
	r.Methods("GET").Path("/Group/{GroupId}").HandlerFunc(handler.Groups)
	r.Methods("GET").Path("/GetMenu").HandlerFunc(handler.Menu)
	r.Methods("GET").Path("/Login").HandlerFunc(handler.GetLogin)
	r.Methods("POST").Path("/Login").HandlerFunc(handler.PostLogin)

	r.Methods("GET").Path("/Upload").HandlerFunc(handler.Upload)
	r.Methods("POST").Path("/Upload").HandlerFunc(handler.PostUpload)
}
func ManageAdminRoutes(r *mux.Router, db datalayer.SQLHandler) {
	handler := newAdminRestApiHandler(db)
	adminRoute := r.PathPrefix("/Admin").Subrouter()
	adminRoute.Use(Authentication)
	adminRoute.Methods("GET").Path("/").HandlerFunc(handler.Dashboard)                      // /Admin/
	adminRoute.Methods("GET").Path("/Posts").HandlerFunc(handler.PostList)                  // /Admin/Posts
	adminRoute.Methods("GET").Path("/AddPost").HandlerFunc(handler.AddPost)                 // /Admin/Posts
	adminRoute.Methods("POST").Path("/AddPost").HandlerFunc(handler.PostAddPost)            // /Admin/Posts
	adminRoute.Methods("GET").Path("/EditPost/{PostId}").HandlerFunc(handler.EditPost)      // /Admin/EditPost
	adminRoute.Methods("POST").Path("/EditPost/{PostId}").HandlerFunc(handler.PostEditPost) // /Admin/EditPost

	adminRoute.Methods("GET").Path("/Groups").HandlerFunc(handler.GroupList)                   // /Admin/Groups
	adminRoute.Methods("GET").Path("/AddGroup").HandlerFunc(handler.AddGroup)                  // /Admin/AddGroup
	adminRoute.Methods("POST").Path("/AddGroup").HandlerFunc(handler.PostAddGroup)             // /Admin/AddGroup
	adminRoute.Methods("GET").Path("/EditGroup/{GroupId}").HandlerFunc(handler.EditGroup)      // /Admin/EditGroup
	adminRoute.Methods("POST").Path("/EditGroup/{GroupId}").HandlerFunc(handler.PostEditGroup) // /Admin/EditGroup
	adminRoute.Methods("GET").Path("/DeleteGroup/{GroupId}").HandlerFunc(handler.DeleteGroup)
}

func Authentication(next http.Handler) http.Handler {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if info, logedin := jwt.IsLogedin(r); logedin {
			// We found the token in our map
			log.Printf("Authenticated user %v\n", info)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "شما به این صفحه دسترسی ندارید", http.StatusForbidden)
		}
	})

	return h
}
