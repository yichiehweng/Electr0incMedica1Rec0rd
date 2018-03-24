package main

import (
	"ElectronicMedicalRecord/medication"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := medication.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
