package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var BookCache map[int]*Book

func LoadBook() (map[int]*Book, error) {
	var books map[int]*Book
	temp, err := os.ReadFile("Data/book.json")
	if err != nil {
		return books, err
	}
	err = json.Unmarshal(temp, &books)
	if err != nil {
		return books, err
	}
	return books, nil
}
func InitBook() {
	var err error
	BookCache, err = LoadBook()
	if err != nil {
		fmt.Println("Файл не найден, создаем пустую базу")
		BookCache = make(map[int]*Book)
	}
}
func SaveBook(books map[int]*Book) error {
	temp, err := json.MarshalIndent(books, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("Data/book.json", temp, 0644)
	if err != nil {
		return err
	}
	return nil
}

func BackStats() {
	for {
		fmt.Printf("Фоновая запись : колличество книг на данные момент : %d\n", len(BookCache))
		time.Sleep(1 * time.Minute)
	}
}
