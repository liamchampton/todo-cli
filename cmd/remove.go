/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			for i := 0; i < len(todos.Todos); i++ {
				if todos.Todos[i].Name == todoName {
					todos.Todos = append(todos.Todos[0:i], todos.Todos[i+1:]...)
					break
				}
			}
			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
