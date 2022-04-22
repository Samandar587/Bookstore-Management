package routes

import (
	"github.com/gorilla/mux"
	"github.com/samandar587/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router){
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}/", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{bookId}/", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}/", controllers.DeleteBook).Methods("DELETE")

}
