/*
Copyright © 2024 Josh Guarino <jguarino722@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// machinesCmd represents the machines command
var machinesCmd = &cobra.Command{
	Use:   "machines",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("machines called")
	},
}

func init() {
	rootCmd.AddCommand(machinesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// machinesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// machinesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
