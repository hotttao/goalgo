/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leecode

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	queue := []*TreeNode{root}
	collect := []string{}
	for len(queue) != 0 {
		if queue[0] != nil {
			collect = append(collect, strconv.Itoa(queue[0].Val))
			queue = append(queue, queue[0].Left, queue[0].Right)
		} else {
			collect = append(collect, "null")
		}
		queue = queue[1:]
	}

	r := "[" + strings.Join(collect, ",") + "]"
	fmt.Println(r)
	return r
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		var p *TreeNode
		return p
	}
	collect := strings.Split(data, ",")
	length := len(collect)
	fmt.Println(length)
	collect[0] = strings.Replace(collect[0], "[", "", 1)
	collect[length-1] = strings.Replace(collect[length-1], "]", "", 1)
	nodes := make([]*TreeNode, length+1, length+1)
	for p := 1; p <= length; p++ {
		val := collect[p-1]
		if val != "null" {
			v, _ := strconv.Atoi(val)
			nodes[p] = &TreeNode{Val: v}
		}
	}
	for i := 1; i <= length/2; i++ {
		left := i * 2
		right := i*2 + 1
		if nodes[i] != nil {
			nodes[i].Left = nodes[left]
			if right <= length {
				nodes[i].Right = nodes[right]
			}
		}
	}

	return nodes[1]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
