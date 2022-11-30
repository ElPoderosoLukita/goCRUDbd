package main

import (
	"net/http"

	"github.com/ElPoderosoLukita/goCRUDbd/handlers"
	"github.com/ElPoderosoLukita/goCRUDbd/models"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/post/user", handlers.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/delete/user/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/update/user/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/get/user/{id}", handlers.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/get/users", handlers.GetUsersHandler).Methods(http.MethodGet)

	models.OpenDB()
	defer models.CloseDatabase()

	http.ListenAndServe(":8081", r)
}
