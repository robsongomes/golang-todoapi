package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robsongomes/golang-todoapi/pkg/models"
	"github.com/robsongomes/golang-todoapi/pkg/utils"
	"gorm.io/gorm"
)

func GetTodos(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	todos := models.GetAllTodos()
	json.NewEncoder(rw).Encode(todos)
}

func GetTodo(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	idParam := vars["id"]
	id, _ := strconv.Atoi(idParam)
	todo, res := models.GetTodoById(uint(id))
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	utils.WriteJson(rw, todo)
}

func CreateTodo(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var todo models.Todo
	utils.ParseBody(req, &todo)
	todo.Create()
	utils.WriteJson(rw, todo)
}

func UpdateTodo(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	idParam := vars["id"]
	id, _ := strconv.Atoi(idParam)
	todoDB, res := models.GetTodoById(uint(id))
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	var todo models.Todo
	utils.ParseBody(req, &todo)
	todoDB.Description = todo.Description
	todoDB.Owner = todo.Owner
	todoDB.Done = todo.Done
	models.UpdateTodo(todoDB)
	utils.WriteJson(rw, todoDB)
}

func DeleteTodo(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	idParam := vars["id"]
	id, _ := strconv.Atoi(idParam)
	res := models.DeleteTodo(uint(id))
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
