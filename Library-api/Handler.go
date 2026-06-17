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
	authorFilter := r.URL.Query().Get("author")
	titleFilter := r.URL.Query().Get("title")
	yearFilter := r.URL.Query().Get("year")
	var year int
	if yearFilter != "" {
		var err error
		year, err = strconv.Atoi(yearFilter)
		if err != nil {
			fmt.Println("Ошибка ввода года написания попробуйте снова")
			return
		}
	}

	filterBook := make(map[int]*Book)
	for id, book := range BookCache {
		matchAuthor := authorFilter == "" || book.Author == authorFilter
		matchTitle := titleFilter == "" || book.Title == titleFilter
		matchYear := yearFilter == "" || book.Year == year
		if matchAuthor && matchTitle && matchYear {
			filterBook[id] = book
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	newEncoder := json.NewEncoder(w)
	newEncoder.SetIndent("", " ")
	newEncoder.Encode(filterBook)
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
func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "Ошибка: неверный формат ID", http.StatusBadRequest)
		return
	}
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	if _, ok := BookCache[id]; !ok {
		http.Error(w, "Книги с таким ID не существует", http.StatusNotFound)
		return
	}
	BookCache[id] = &newBook
	err = SaveBook(BookCache)
	if err != nil {
		fmt.Println("Ошибка сохранения")
		return
	}
	fmt.Fprintf(w, "Данные обновлены")
}
