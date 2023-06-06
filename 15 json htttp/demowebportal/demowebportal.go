package demowebportal

import (
	"fmt"
	"net/http"
)

//RunWebPortal can Run webPortal
func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welome to my Server. %s", r.RemoteAddr)
}
