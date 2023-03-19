package cmd_algo

import (
	"fmt"

	"github.com/hotttao/goalgo/pkg/algo"
	"github.com/spf13/cobra"
)

// quickCmd represents the quick command
var QuickCmd = &cobra.Command{
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
