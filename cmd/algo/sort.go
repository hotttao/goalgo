package cmd_algo

import (
	"fmt"

	"github.com/hotttao/goalgo/pkg/algo"
	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var SortCmd = &cobra.Command{
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

		data = []int{1, 2, 1, 6, 9, 10, 3}
		fmt.Println("dubble sort")
		data = algo.DoubbleSort(data)
		fmt.Println(data)

		data = []int{1, 2, 1, 6, 9, 10, 3}
		fmt.Println("insert sort")
		algo.InsertSort(data)
		fmt.Println(data)

		data = []int{1, 2, 1, 6, 9, 10, 3}
		fmt.Println("select sort")
		algo.InsertSort(data)
		fmt.Println(data)

		data = []int{1, 2, 1, 6, 9, 10, 3}
		fmt.Println("count sort")
		data = algo.CountSort(data)
		fmt.Println(data)

	},
}
