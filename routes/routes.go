package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	InitHandlers()

	router := mux.NewRouter()

	router.HandleFunc("/items", itemController.GetItems).Methods("GET")
	router.HandleFunc("/item/{id}", itemController.GetItem).Methods("GET")
	router.HandleFunc("/item", itemController.CreateItem).Methods("POST")

	http.ListenAndServe(":8000", router)
}
