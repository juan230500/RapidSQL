/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "insert" que se usa para insertar nuevos registros en la tabla "Book" de una base de datos SQL.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// insertCmd representa el comando "insert". Este comando se utiliza para insertar nuevos registros en la base de datos.
var insertCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "insert", seguido de dos flags "-t" y "-a" para "title" y "author" respectivamente.
	Use: "insert",
	// Una descripción corta del comando.
	Short: "Inserta un nuevo libro en la base de datos",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Inserta un nuevo registro en la tabla "Book" en la base de datos. Requiere dos argumentos:

	-t Título del libro
	-a Autor del libro

Por ejemplo, puedes usar el comando de esta manera:

	insert -t "Título del libro" -a "Autor del libro"

Este comando se conectará a la base de datos y creará un nuevo registro con los datos proporcionados.`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		db := DBconn()
		defer db.Close()

		title, _ := cmd.Flags().GetString("title")
		author, _ := cmd.Flags().GetString("author")

		_, err := db.Exec(`INSERT INTO Book (title, author) VALUES ($1, $2)`, title, author)
		if err != nil {
			panic(err)
		}

		fmt.Println("Libro insertado correctamente")
	},
}

func init() {
	// Define los flags y sus descripciones para el comando "insert"
	insertCmd.Flags().StringP("title", "t", "", "Título del libro a insertar")
	insertCmd.Flags().StringP("author", "a", "", "Autor del libro a insertar")
	rootCmd.AddCommand(insertCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// insertCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
