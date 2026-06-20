package main

import (
	"database/sql"
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
	var args []any
	Query := `SELECT id,title,author,year, is_available FROM Books WHERE 1=1`
	if yearFilter != "" {
		var err error
		year, err = strconv.Atoi(yearFilter)
		if err != nil {
			fmt.Println("Ошибка ввода года написания попробуйте снова")
			return
		}
		args = append(args, year)
		Query += fmt.Sprintf(" AND year = %d", len(args))
	}
	if titleFilter != "" {
		args = append(args, titleFilter)
		Query += fmt.Sprintf(" AND title = %s", len(args))
	}
	if authorFilter != "" {
		args = append(args, authorFilter)
		Query += fmt.Sprintf(" AND author = %s", len(args))
	}
	rows, err := DB.Query(Query, args...)
	if err != nil {
		fmt.Fprintf(w, "Ошибка запроса")
		return
	}
	defer rows.Close()
	var result []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.IsAvailable)
		if err != nil {
			fmt.Println("Ошибка чтения строки:", err)
			continue
		}
		result = append(result, book)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(result)
}
func BookHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		fmt.Println("Ошибка перевода id")
		return
	}
	var book Book
	bookQuery := `SELECT id,title,author,is_available FROM books WHERE id = @p1`
	err = DB.QueryRow(bookQuery, id).Scan(&book.ID, &book.Title, &book.Author, &book.IsAvailable)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Книга не найдена", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Ошибка чтения БД", http.StatusInternalServerError)
		}
	}
}
func addBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	insertQuery := `INSERT INTO books(title, author, year, is_available) VALUES (@p1, @p2, @p3, @p4)`
	_, err := DB.Exec(insertQuery, newBook.Title, newBook.Author, newBook.Year, newBook.IsAvailable)

	if err != nil {
		fmt.Println("Ошибка при добавлении в БД:", err)
		http.Error(w, "Ошибка сохранения в базу данных", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "книга добавлена")
}
func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		return
	}
	deleteQuery := `DELETE  FROM books WHERE id = @p1`
	result, err := DB.Exec(deleteQuery, id)
	if err != nil {
		fmt.Fprintf(w, "Ошибка удаления")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, "Ошибка удаления")
		return
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Книги с таким id не существует")
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
	updateQuery := `UPDATE books SET title = @p1, author = @p2, year=@p3, is_available = @p4 WHERE id = @p5`
	result, err := DB.Exec(updateQuery, newBook.Title, newBook.Author, newBook.Year, newBook.IsAvailable, id)
	if err != nil {
		fmt.Println("Ошибка сохранения")
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Fprintf(w, "Ошибка с БД")
	}
	if rowsAffected == 0 {
		fmt.Fprintf(w, "Ошибка, данной книги не существует", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Данные обновлены")
}
