package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Combine 结构组合接口
type C struct {
	Name string
}

func (c C) GetName() string {
	return c.Name
}

type A struct {
	Owner string
}

func (a A) GetName() string {
	return a.Owner
}

// Output 结构嵌套接口
type Output struct {
	C
	A
}

func (o Output) GetName() string {
	return o.C.Name
}

// trialCmd represents the trial command
var trialCmd = &cobra.Command{
	Use:   "trial",
	Short: "trail some package",
	Long:  `trail some package`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("trial called")
		o := Output{
			A: A{Owner: "aa"},
			C: C{Name: "cc"},
		}
		fmt.Printf("%v\n", o.GetName())
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
