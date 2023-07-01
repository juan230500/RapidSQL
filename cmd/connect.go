/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando "connect" que se utiliza para establecer una conexión con una base de datos SQL y crear una tabla "Book" si no existe.
package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	// Uso del comando. El comando se llama con "connect", seguido de cinco flags para especificar los detalles de la conexión a la base de datos.
	Use: "connect",
	// Una descripción corta del comando.
	Short: "Conecta a la base de datos y crea la tabla 'Book'",
	// Descripción detallada del comando y cómo usarlo.
	Long: `Conecta a la base de datos PostgreSQL proporcionada y crea la tabla 'Book' si no existe. Se deben proporcionar los detalles de la conexión a la base de datos:

	host: el host de la base de datos
	port: el puerto de la base de datos
	user: el nombre de usuario de la base de datos
	password: la contraseña de la base de datos
	dbname: el nombre de la base de datos

Por ejemplo, puedes usar el comando de esta manera:

	connect --host localhost --port 5432 --user postgres --password secreto --dbname libreria`,
	// Función que se ejecuta cuando se llama al comando.
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		dbname, _ := cmd.Flags().GetString("dbname")
		dataIn := Data{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Dbname:   dbname,
		}

		jsonData, err := json.Marshal(dataIn)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			return
		}

		err = os.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error al escribir el archivo JSON:", err)
			return
		}

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		_, err = db.Exec(`DROP TABLE IF EXISTS Book`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Book (id SERIAL PRIMARY KEY, title VARCHAR(50), author VARCHAR(50))`)
		if err != nil {
			panic(err)
		}

		fmt.Println("Conectado exitosamente y creada la tabla 'Book'!")
	},
}

func init() {
	connectCmd.Flags().StringP("host", "", "", "Host de la base de datos a conectar")
	connectCmd.Flags().IntP("port", "", 0, "Puerto de la base de datos a conectar")
	connectCmd.Flags().StringP("user", "", "", "Nombre de usuario de la base de datos a conectar")
	connectCmd.Flags().StringP("password", "", "", "Contraseña de la base de datos a conectar")
	connectCmd.Flags().StringP("dbname", "", "", "Nombre de la base de datos a conectar")
	rootCmd.AddCommand(connectCmd)

	// Aquí puedes definir tus flags y configuraciones.

	// Cobra admite Flags Persistentes que funcionarán para este comando y todos los subcomandos, por ejemplo:
	// connectCmd.PersistentFlags().String("foo", "", "Una ayuda para foo")

	// Cobra admite flags locales que solo se ejecutarán cuando se llame directamente a este comando, por ejemplo:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
