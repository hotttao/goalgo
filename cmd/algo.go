/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	cmd_algo "github.com/hotttao/goalgo/cmd/algo"
	"github.com/spf13/cobra"
)

// algoCmd represents the algo command
var algoCmd = &cobra.Command{
	Use:   "algo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("algo called")
	},
}

func init() {
	rootCmd.AddCommand(algoCmd)
	algoCmd.AddCommand(cmd_algo.BinaryCmd)
	algoCmd.AddCommand(cmd_algo.ListCmd)
	algoCmd.AddCommand(cmd_algo.MapCmd)
	algoCmd.AddCommand(cmd_algo.QuickCmd)
	algoCmd.AddCommand(cmd_algo.SortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// algoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// algoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
