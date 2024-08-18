package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	fmt.Println("Start db")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//установка данных
	// insertInDb(db, "Tom", 32)
	// insertInDb(db, "Ben", 22)
	// insertInDb(db, "Max", 12)

	//получение данных
	getAllUsers(db)
	//Добавить всем +10 к возрасту
	// increaseAllAge(db, 10)

	updateName(db, "new Name", 8)

	getAllUsers(db)
	//Удаление данных
	// deleteUsers(db)

	//Обновление данных
	// insertInDb(db, user.Name+"1", user.Age+10)
	fmt.Println("Подключено к mySql")
}

// Функции
func updateName(db *sql.DB, name string, id uint16) {
	update, err := db.Query("UPDATE `users` SET `name` = ? WHERE `id` = ?", name, id)

	if err != nil {
		panic(err)
	}
	defer update.Close()
}
func insertInDb(db *sql.DB, name string, age uint16) {
	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES (?, ?)", name, age)
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}
func increaseAllAge(db *sql.DB, value uint16) {
	update, err := db.Query("UPDATE `users` SET `age` = `age` + ?", value)
	if err != nil {
		panic(err)
	}
	defer update.Close()
	fmt.Println("Возраст увеличен на", value)
}
func deleteUsers(db *sql.DB) {
	delete, err := db.Query("DELETE FROM `users`")
	if err != nil {
		panic(err)
	}
	defer delete.Close()
}
func getAllUsers(db *sql.DB) {
	res, err := db.Query("SELECT name, age FROM `users`")

	if err != nil {
		panic(err)
	}

	defer res.Close()

	for res.Next() {
		var user User
		err := res.Scan(&user.Name, &user.Age)

		if err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s, age: %d \n", user.Name, user.Age)

	}
}
