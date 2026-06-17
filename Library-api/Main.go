package main

import (
	"fmt"
	"net/http"
)

func main() {
	InitBook()
	mux := http.NewServeMux()
	loggegMux := LogingMiddelware(mux)
	mux.HandleFunc("GET /", HelloUserHandler)
	mux.HandleFunc("GET /books", BooksHandler)
	mux.HandleFunc("GET /book/{id}", BookHandler)
	mux.HandleFunc("PUT /add/book", addBookHandler)
	mux.HandleFunc("DELTE /delete/book/{id}", deleteBookHandler)
	fmt.Println("Сервер запущен по адресу : http://localhost:8000")
	http.ListenAndServe(":8000", loggegMux)
}
