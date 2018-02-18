package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

func main() {

	fmt.Println("Rest API - Mux Routers and Bongo")

	router := mux.NewRouter().StrictSlash(true)

	usuarioRouter := router.PathPrefix("/usuario").Subrouter()
	usuarioRouter.HandleFunc("", usuarioResourceList).Methods("GET")
	usuarioRouter.HandleFunc("", usuarioResourcePost).Methods("POST")
	usuarioRouter.HandleFunc("/{id}", usuarioResourceFindById).Methods("GET")
	usuarioRouter.HandleFunc("/{id}", usuarioResourcePut).Methods("PUT")
	usuarioRouter.HandleFunc("/{id}", usuarioResourceDelete).Methods("DELETE")

	router.Use(requestMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func requestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)

	})
}
