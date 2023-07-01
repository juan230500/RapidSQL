/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

// Este paquete define el comando raíz "RapidSQL", la pieza central de nuestra emocionante aplicación de línea de comandos para interactuar con bases de datos PostgreSQL.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd representa el comando base cuando se llama sin ningún subcomando.
// RapidSQL es un revolucionario punto de entrada para nuestra aplicación y contendrá subcomandos para operaciones específicas de la base de datos PostgreSQL.
var rootCmd = &cobra.Command{
	// El comando se llama con "RapidSQL".
	Use: "RapidSQL",
	// Una descripción vibrante y concisa de nuestra aplicación.
	Short: "RapidSQL: tu centro de comando definitivo para la gestión de bases de datos PostgreSQL",
	// Una descripción detallada y cautivadora de nuestra aplicación y cómo usarla.
	Long: `RapidSQL es una potente herramienta de línea de comandos para administrar tus bases de datos PostgreSQL. Es tu varita mágica para realizar operaciones CRUD, conectarte con diferentes bases de datos, y ejecutar consultas SQL directas, todo desde la comodidad de tu terminal.

Imagina que tienes que insertar un nuevo libro en tu base de datos. Con RapidSQL, es tan fácil como esto:

	RapidSQL insert -t "Guerra y Paz" -a "León Tolstoi"

Este comando se conectará a tu base de datos PostgreSQL y creará un nuevo registro con estos datos. ¡Increíble, verdad?`,
}

// Execute agrega todos los subcomandos al comando raíz y establece los flags apropiadamente.
// Esto es llamado por main.main(). Solo necesita suceder una vez en rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Aquí es donde puedes definir tus flags y configuraciones.
	// Cobra admite flags persistentes, que, si se definen aquí, serán globales para tu aplicación.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "archivo de configuración (por defecto es $HOME/.RapidSQL.yaml)")

	// Cobra también admite flags locales, que solo se ejecutarán cuando esta acción se llame directamente.
	rootCmd.Flags().BoolP("toggle", "t", false, "Mensaje de ayuda para toggle")
}
