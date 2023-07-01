/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "update" que se utiliza para actualizar registros existentes en la tabla "Book" de una base de datos SQL.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd representa el comando "update". Este comando se utiliza para actualizar registros existentes en la base de datos.
var updateCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "update", seguido de tres flags "-i" para "id", "-t" para "title" y "-a" para "author".
	Use: "update",
	// Una descripción corta del comando.
	Short: "Actualiza un libro existente en la base de datos",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Actualiza un registro existente en la tabla "Book" en la base de datos. Requiere tres argumentos:

	-i ID del libro
	-t Título del libro
	-a Autor del libro

Por ejemplo, puedes usar el comando de esta manera:

	update -i 1 -t "Nuevo título del libro" -a "Nuevo autor del libro"

Este comando se conectará a la base de datos y actualizará el registro con los datos proporcionados.`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		db := DBconn()
		defer db.Close()

		id, _ := cmd.Flags().GetString("id")
		title, _ := cmd.Flags().GetString("title")
		author, _ := cmd.Flags().GetString("author")

		_, err := db.Exec(`UPDATE Book SET title = $1, author = $2  WHERE ID = $3;`, title, author, id)
		if err != nil {
			panic(err)
		}

		fmt.Println("Libro actualizado correctamente")
	},
}

func init() {
	// Define los flags y sus descripciones para el comando "update"
	updateCmd.Flags().StringP("id", "i", "", "ID del libro a actualizar")
	updateCmd.Flags().StringP("title", "t", "", "Nuevo título del libro")
	updateCmd.Flags().StringP("author", "a", "", "Nuevo autor del libro")
	rootCmd.AddCommand(updateCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// updateCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
