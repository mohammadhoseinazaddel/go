package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Methods("GET").Path("/CreateCookie").HandlerFunc(CreateCookie)
	router.Methods("GET").Path("/ReadCookie").HandlerFunc(ReadCookie)
	router.Methods("GET").Path("/DeleteCookie").HandlerFunc(DeleteCookie)
	log.Fatal(http.ListenAndServe(":8484", router))
}

func CreateCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "MyCookie",
		Value:  "myCookieValue",
		MaxAge: 3600,
		//Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "کوکی ست شد.")
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("MyCookie")
	if err != nil {
		fmt.Fprintln(w, "کوکی یافت نشد.")
		return
	}

	fmt.Fprintln(w, "Name : ", cookie.Name, "  Value: ", cookie.Value)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "MyCookie",
		MaxAge: -1,
		//Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "کوکی پاک شد.")
}
