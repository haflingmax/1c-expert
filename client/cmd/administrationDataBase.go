/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// administrationCmd represents the administration command
var databaseConnectionTestCmd = &cobra.Command{
	Use:   "database-connection-test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		if host == "" {
			host = "localhost"
		}

		if port == 0 {
			port = 1433
		}

		if user == "" {
			user = "sa"
		}

		if password == "" {
			password = "1"
		}

		switch command {
		case "database-connection-test":
			databaseConnectionTest(host, port, user, password)
		}
	},
}

func init() {
	administrationCmd.AddCommand(administrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// administrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	administrationCmd.Flags().String("host", "localhost", "Сервер СУБД")
	administrationCmd.Flags().Int("port", 1433, "Порт сервера СУБД")
	administrationCmd.Flags().String("user", "sa", "Учетная запись для подключения к СУБД")
	administrationCmd.Flags().String("password", "1", "Пароль для подключения к СУБД")
}
