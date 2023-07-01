/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "select" que se usa para seleccionar y visualizar registros de la tabla "Book" de una base de datos SQL.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// selectCmd representa el comando "select". Este comando se utiliza para seleccionar registros de la base de datos.
var selectCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "select", seguido de una flag opcional "-w" para "where" que limita los resultados seleccionados.
	Use: "select",
	// Una descripción corta del comando.
	Short: "Selecciona y muestra los libros de la base de datos",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Selecciona registros de la tabla "Book" en la base de datos y los muestra. Se puede proporcionar un argumento opcional:

	-w Condición WHERE para limitar los resultados (por ejemplo, "author = 'Autor'")

Por ejemplo, puedes usar el comando de esta manera:

	select -w "author = 'Autor del libro'"

Este comando se conectará a la base de datos, seleccionará los registros que cumplan con la condición proporcionada (o todos los registros si no se proporciona ninguna condición) y mostrará los resultados.`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		db := DBconn()
		defer db.Close()

		whereStr, _ := cmd.Flags().GetString("where")

		var query string
		if whereStr != "" {
			query = "SELECT * FROM Book WHERE " + whereStr
		} else {
			query = "SELECT * FROM Book"
		}

		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if !rows.Next() {
			fmt.Println("---No rows---")
		}

		for rows.Next() {
			var id int
			var title string
			var author string
			err := rows.Scan(&id, &title, &author)
			if err != nil {
				panic(err)
			}
			fmt.Printf("ID: %d, Título: %s, Autor: %s\n", id, title, author)
		}
	},
}

func init() {
	// Define los flags y sus descripciones para el comando "select"
	selectCmd.Flags().StringP("where", "w", "", "Condición WHERE para limitar los resultados seleccionados")
	rootCmd.AddCommand(selectCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// selectCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
