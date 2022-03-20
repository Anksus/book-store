package routes

import (
	"github.com/Anksus/book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookByIdBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBooks).Methods("DELETE")

}
