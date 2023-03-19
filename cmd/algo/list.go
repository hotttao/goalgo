package cmd_algo

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Test struct {
	A int
}

func (t *Test) String() string {
	s := fmt.Sprintf("Test{A:%d}", t.A)
	fmt.Println(s)
	return s
}

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "array and link",
	Long:  `array and link`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		// for i := 1; i < 10; i++ {
		// 	go func(i int) {
		// 		fmt.Printf("value: %d\n", i)
		// 	}(i)
		// }
		// time.Sleep(time.Second)
		var t Test
		f := t.String
		fmt.Printf("%T\n", f)
		fmt.Printf("%T", (*Test).String)
	},
}
