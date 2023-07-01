/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "delete" que se usa para eliminar registros en la tabla "Book" de una base de datos SQL.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd representa el comando "delete". Este comando se utiliza para eliminar un registro en la base de datos.
var deleteCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "delete", seguido del flag "-i" para el "id" del libro.
	Use: "delete",
	// Una descripción corta del comando.
	Short: "Elimina un libro de la base de datos",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Elimina un registro en la tabla "Book" en la base de datos. Requiere un argumento:

	-i ID del libro

Por ejemplo, puedes usar el comando de esta manera:

	delete -i 123

Este comando se conectará a la base de datos y eliminará el registro correspondiente al ID proporcionado.`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		db := DBconn()
		defer db.Close()

		id, _ := cmd.Flags().GetString("id")

		a1, err := db.Exec(`DELETE FROM Book WHERE ID = $1`, id)
		if err != nil {
			panic(err)
		}
		a1.RowsAffected()

		fmt.Println("Libro eliminado correctamente")
	},
}

func init() {
	// Define los flags y sus descripciones para el comando "delete"
	deleteCmd.Flags().StringP("id", "i", "", "ID del libro a eliminar")
	rootCmd.AddCommand(deleteCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// deleteCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
