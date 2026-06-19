package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb" //Подключение драйвера
)

// Глобальная переменная для базы данных.
// Через нее весь наш проект будет отправлять запросы.
var DB *sql.DB

func InitDB() {
	// ... (твоя строка подключения)
	connString := "server=localhost;port=1433;database=LibraryDB;Trusted_Connection=True;"

	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Ошибка настройки подключения: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Сервер MS SQL не отвечает: ", err)
	}

	fmt.Println("[БД] Ура! Мы успешно подключились к MS SQL Server!")

	// === НОВЫЙ КОД НИЖЕ ===

	// Пишем SQL-запрос для MS SQL
	createTableQuery := `
	IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='books' and xtype='U')
	BEGIN
		CREATE TABLE books (
			id INT IDENTITY(1,1) PRIMARY KEY,
			title NVARCHAR(255) NOT NULL,
			author NVARCHAR(255) NOT NULL,
			year INT,
			is_available BIT DEFAULT 1
		)
	END;`

	// Выполняем запрос через функцию DB.Exec()
	// Она нужна для запросов, которые не возвращают данные (CREATE, INSERT, UPDATE, DELETE)
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Ошибка создания таблицы books: ", err)
	}

	fmt.Println("[БД] Таблица books готова к работе!")
}
