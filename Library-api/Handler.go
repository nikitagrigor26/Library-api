package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func HelloUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Добро пожаловать в библиотеку")
}
func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(BookCache)
}
func BookHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		return
	}
	if _, ok := BookCache[id]; !ok {
		fmt.Fprintf(w, "Ошибка, книги с таким id не существует")
		return

	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(BookCache[id])
}
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	if _, ok := BookCache[newBook.ID]; ok {
		fmt.Fprintf(w, "Книга с таким Id существует, введите новый id")
		return
	}
	BookCache[newBook.ID] = &newBook
	err := SaveBook(BookCache)
	if err != nil {
		fmt.Fprintf(w, "Ошибка записи")
		return
	}
	fmt.Fprintf(w, "Книга добавлена")
}
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		return
	}
	if _, ok := BookCache[id]; !ok {
		fmt.Fprintf(w, "Книги с таким id не существует, удаление невозможно")
		return
	}
	delete(BookCache, id)
	err = SaveBook(BookCache)
	if err != nil {
		fmt.Fprintf(w, "Ошибка сохранения")
		return
	}
	fmt.Fprintf(w, "Книга удалена")
}
