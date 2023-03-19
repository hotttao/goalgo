package cmd_algo

import (
	"fmt"

	"github.com/hotttao/goalgo/pkg/algo"
	"github.com/spf13/cobra"
)

// binaryCmd represents the binary command
var BinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "binary search",
	Long:  `binary`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("binary called")
		fmt.Println("equal search")
		data := []int{1, 5, 6, 13, 13, 13, 13, 13, 14, 21, 58, 69}
		v := algo.BinarySearch(data, 14)
		fmt.Printf("search %d at index: %d\n", 14, v)

		fmt.Println("equal first")
		v = algo.BinaryEqualFirst(data, 13)
		fmt.Printf("search first equal %d at index: %d\n", 13, v)

		fmt.Println("equal last")
		v = algo.BinaryEqualEnd(data, 13)
		fmt.Printf("search first equal %d at index: %d\n", 13, v)

		fmt.Println("gte first")
		v = algo.BinaryGteFirst(data, 15)
		fmt.Printf("search gte %d at index: %d\n", 15, v)

		fmt.Println("lte en")
		v = algo.BinaryLteEnd(data, 15)
		fmt.Printf("search lte %d at index: %d\n", 15, v)
	},
}
