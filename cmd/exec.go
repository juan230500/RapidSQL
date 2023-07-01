/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "exec" que se usa para ejecutar consultas SQL personalizadas en la base de datos.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// execCmd representa el comando "exec". Este comando se utiliza para ejecutar consultas SQL personalizadas en la base de datos.
var execCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "exec", seguido del flag "-q" para "query".
	Use: "exec",
	// Una descripción corta del comando.
	Short: "Ejecuta una consulta SQL personalizada en la base de datos",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Ejecuta una consulta SQL personalizada en la base de datos. Requiere un argumento:

	-q Consulta SQL personalizada

Por ejemplo, puedes usar el comando de esta manera:

	exec -q "SELECT * FROM Book"

Este comando se conectará a la base de datos y ejecutará la consulta SQL proporcionada.`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		db := DBconn()
		defer db.Close()

		query, _ := cmd.Flags().GetString("query")

		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		if !rows.Next() {
			fmt.Println("---No rows---")
			return
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
	// Define los flags y sus descripciones para el comando "exec"
	execCmd.Flags().StringP("query", "q", "", "Consulta SQL personalizada a ejecutar")
	rootCmd.AddCommand(execCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// execCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// execCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
