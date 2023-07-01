package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear tabla
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()

	// Insertar datos
	statement, _ = db.Prepare("INSERT INTO users (firstname, lastname) VALUES (?, ?)")
	statement.Exec("John", "Doe")

	// Leer datos
	rows, _ := db.Query("SELECT id, firstname, lastname FROM users")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(id, firstname, lastname)
	}

	// Actualizar datos
	statement, _ = db.Prepare("UPDATE users SET firstname = ? WHERE id = ?")
	statement.Exec("Jane", 1)

	// Verificar la actualización
	rows, _ = db.Query("SELECT id, firstname, lastname FROM users")
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(id, firstname, lastname)
	}

	// Eliminar datos
	statement, _ = db.Prepare("DELETE FROM users WHERE id = ?")
	statement.Exec(1)

	// Verificar la eliminación
	rows, _ = db.Query("SELECT id, firstname, lastname FROM users")
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(id, firstname, lastname)
	}
}
