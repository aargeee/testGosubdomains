package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{subdomain:[a-z]+}", subdomainHandler)
	router.HandleFunc("/", defaultHandler)

	http.Handle("/", router)
	fmt.Printf("Port : %d", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%d", os.Getenv("PORT")), nil)
}

func subdomainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subdomain := vars["subdomain"]
	fmt.Fprintf(w, "Subdomain: %s, Path: %s", subdomain, r.URL.Path)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "No subdomain, Path: ", r.URL.Path)
}
