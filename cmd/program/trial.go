package cmd_program

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
var TrialCmd = &cobra.Command{
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
