package cmd

import (
	"fmt"

	"github.com/hotttao/goalgo/algo"
	"github.com/spf13/cobra"
)

// binaryCmd represents the binary command
var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "binary search",
	Long:  `binary`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("binary called")
		fmt.Println("equal search")
		data := []int{1, 5, 6, 13, 21, 58, 69}
		v := algo.BinarySearch(data, 13)
		fmt.Printf("search %d at index: %d\n", 13, v)
		fmt.Printf("search %d at index: %d\n", 14, algo.BinarySearch(data, 14))
	},
}

func init() {
	rootCmd.AddCommand(binaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// binaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// binaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
