package router

import (
	"github.com/gorilla/mux"
	"github.com/yogesh-p-thakare3110/Seperate/handler"
)

func HandlerRouting() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", handler.getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.deleteMovie).Methods("DELETE")

	return r
}
