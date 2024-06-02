package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
    var err error
    // Подключение к базе данных
    db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    // Проверка подключения
    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %v\n", err)
    }

    fmt.Println("Successfully connected to database!")
}

func main() {
    initDB()
    defer db.Close()

	insert, err := db.Query("")
    // Ваш дальнейший код работы с базой данных
}