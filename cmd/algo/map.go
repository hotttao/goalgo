package cmd_algo

import (
	"fmt"

	"github.com/hotttao/goalgo/pkg/algo"
	"github.com/spf13/cobra"
)

// mapCmd represents the map command
var MapCmd = &cobra.Command{
	Use:   "map",
	Short: "map structure",
	Long:  `map structure`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test point equal")
		algo.PointEqual()
		fmt.Println("map call")
		m := algo.NewProbeHashMap(3)
		m.Set(1, "abc")
		v, ok := m.Get(1)
		fmt.Println(v, ok)
		fmt.Println(m)

		m.Set(1, "abc")
		v, ok = m.Get(1)
		fmt.Println(v, ok)
		fmt.Println(m)

		m.Set(2, "nnnn")
		v, ok = m.Get(2)
		fmt.Println(v, ok)
		fmt.Println(m)

		m.Set(4, "4444")
		v, ok = m.Get(4)
		fmt.Println(v, ok)
		fmt.Println(m)

		m.Del(1)
		fmt.Println(m)

		m.Set(2, "222")
		v, ok = m.Get(2)
		fmt.Println(v, ok)
		fmt.Println(m)
	},
}
