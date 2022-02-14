package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yogesh-p-thakare3110/Seperate/router"
)

func main() {
	r := router.HandlerRouting()

	fmt.Println("API with seperate ones application")
	log.Fatal(http.ListenAndServe(":8000", r))
}
