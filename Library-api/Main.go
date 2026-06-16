package main

import "net/http"

func main() {
	InitBook()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", HelloUserHandler)
	mux.HandleFunc("GET /books", BooksHandler)
	mux.HandleFunc("GET /book/{id}", BookHandler)
	mux.HandleFunc("PUT /add/book", addBookHandler)
	mux.HandleFunc("DELTE /delete/book/{id}", deleteBookHandler)
	
	http.ListenAndServe(":8080", mux)
}
