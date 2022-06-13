package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robsongomes/golang-todoapi/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)

	fmt.Println("Server running on port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
