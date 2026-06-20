package main

import (
	"fmt"
	"time"
)

func BackStats() {
	for {
		time.Sleep(10 * time.Second)

		var count int
		err := DB.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)

		if err != nil {
			fmt.Println("[Статистика] Ошибка подсчета книг:", err)
			continue
		}

		fmt.Printf("[Статистика] Прямо сейчас в базе хранится книг: %d шт.\n", count)
	}
}
