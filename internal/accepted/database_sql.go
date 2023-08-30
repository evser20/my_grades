package accepted

import (
	"database/sql"
	"fmt"
)

// database/sql
func DatabaseSQL() {
	// Устанавливаем соединение с базой данных PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5432 user=your_username password=your_password dbname=your_database sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Подготавливаем SQL-запрос
	query := "SELECT name, age FROM users"
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Выполняем запрос и получаем результаты
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Обрабатываем результаты
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	// Проверяем наличие ошибок после обхода результатов
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
