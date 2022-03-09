package router

import (
	"github.com/gorilla/mux"
	"github.com/yogesh-p-thakare3110/Seperate/handler"
)

func HandlerRouting() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")

	return r
}
