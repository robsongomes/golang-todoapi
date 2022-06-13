package routes

import (
	"github.com/gorilla/mux"
	"github.com/robsongomes/golang-todoapi/pkg/controllers"
)

var Routes = func(router *mux.Router) {
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}
