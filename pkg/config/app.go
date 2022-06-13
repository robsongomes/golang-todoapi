package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getEnv(key string, guard string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return guard
}

var db_user = getEnv("DB_USER", "root")
var db_password = getEnv("DB_PASSWORD", "root")
var db_host = getEnv("DB_HOST", "127.0.0.1")
var db_port = getEnv("DB_PORT", "3306")
var db_name = getEnv("DB_NAME", "tododb")

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_password,
		db_host,
		db_port,
		db_name)
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = con
}

func DB() *gorm.DB {
	return db
}
