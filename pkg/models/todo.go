package models

import (
	"time"

	"github.com/robsongomes/golang-todoapi/pkg/config"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint `gorm:"primary_key"`
	Description string
	Done        bool
	Owner       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.DB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) Create() *Todo {
	db.Create(t)
	return t
}

func GetAllTodos() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func GetTodoById(id uint) (*Todo, *gorm.DB) {
	var todo Todo
	result := db.First(&todo, id)
	return &todo, result
}

func DeleteTodo(id uint) *gorm.DB {
	result := db.Delete(&Todo{}, id)
	return result
}

func UpdateTodo(todo *Todo) *gorm.DB {
	result := db.Save(todo)
	return result
}
