package cmd

import (
	"fmt"

	"github.com/hotttao/goalgo/algo"
	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "run sort algo",
	Long:  `run sort algo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sort called")
		fmt.Println("merge sort called")
		data := []int{1, 2, 1, 6, 9, 10, 3}
		data = algo.MergeSort(data)
		fmt.Println(data)

		data = []int{1, 2, 1, 6, 9, 10, 3}
		fmt.Println("heap sort")
		data = algo.HeapSort(data)
		fmt.Println(data)
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
