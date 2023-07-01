package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

func readData() *Data {
	var readData Data

	fileData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return nil
	}

	err = json.Unmarshal(fileData, &readData)
	if err != nil {
		fmt.Println("Error al convertir el JSON:", err)
		return nil
	}

	return &readData
}

func DBconn() *sql.DB {
	readData := readData()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		readData.Host, readData.Port, readData.User, readData.Password, readData.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return db
}
