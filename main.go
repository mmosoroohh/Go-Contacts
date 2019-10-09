package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mmosoroohh/go-contact/app"
	"github.com/mmosoroohh/go-contact/controllers"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // attach JWT auth middleware

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/user/{id}/contacts", controllers.GetContactsFor).Methods("GET")

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT") // Get port from .env file, when we do not specified any port we should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) // Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}