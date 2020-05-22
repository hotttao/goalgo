package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Combine 结构组合接口
type Combine interface {
	Name()
}

// Output 结构嵌套接口
type Output struct {
	Combine
}

// trialCmd represents the trial command
var trialCmd = &cobra.Command{
	Use:   "trial",
	Short: "trail some package",
	Long:  `trail some package`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trial called")
		// timer := time.NewTimer(3 * time.Second)
		// go func() {
		// 	select {
		// 	case <-timer.C:
		// 		fmt.Println("timer is running")
		// 	}
		// }()
		// time.Sleep(1 * time.Second)
		// fmt.Println(timer.Stop())
		// fmt.Println("timer stop")
		// time.Sleep(5 * time.Second)
		v := Output{}
		fmt.Println(v)
		v.Name()
	},
}

func init() {
	rootCmd.AddCommand(trialCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trialCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trialCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
