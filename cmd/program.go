/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	cmd_program "github.com/hotttao/goalgo/cmd/program"
	"github.com/spf13/cobra"
)

// programCmd represents the program command
var programCmd = &cobra.Command{
	Use:   "program",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("program called")
	},
}

func init() {
	rootCmd.AddCommand(programCmd)
	programCmd.AddCommand(cmd_program.TrialCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// programCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// programCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
