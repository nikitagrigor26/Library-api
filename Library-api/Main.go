package main

import (
	"fmt"
	"net/http"
)

func main() {
	InitBook()
	mux := http.NewServeMux()
	loggegMux := LogingMiddelware(mux)
	go BackStats()
	mux.HandleFunc("GET /", HelloUserHandler)
	mux.HandleFunc("GET /books", BooksHandler)
	mux.HandleFunc("GET /book/{id}", BookHandler)
	mux.HandleFunc("POST /add/book", addBookHandler)
	mux.HandleFunc("DELTE /delete/book/{id}", deleteBookHandler)
	mux.HandleFunc("PUT /update/book/{id}", updateBookHandler)
	fmt.Println("Сервер запущен по адресу : http://localhost:8000")
	http.ListenAndServe(":8000", loggegMux)
}
