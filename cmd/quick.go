package cmd

import (
	"fmt"

	"github.com/hotttao/goalgo/pkg/algo"
	"github.com/spf13/cobra"
)

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "rum quick sort",
	Long:  `run quick sort`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quick called")
		data := []int{1, 2, 1, 6, 9, 10, 3}
		algo.QuickSort(data, 0, len(data)-1)
		fmt.Println(data)
	},
}

func init() {
	sortCmd.AddCommand(quickCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
