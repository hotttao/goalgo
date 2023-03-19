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
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	de []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + this.serialize(root.Left) + this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	this.de = strings.Split(data, ",")
	return this.dfs()
}

func (this *Codec) dfs() *TreeNode {
	if len(this.de) == 0 {
		return nil
	}
	node := this.de[0]
	this.de = this.de[1:]
	if node == "null" {
		return nil
	}
	v, _ := strconv.Atoi(node)

	parent := &TreeNode{
		Val:   v,
		Left:  this.dfs(),
		Right: this.dfs(),
	}
	return parent
}

// // Serializes a tree to a single string.
// func (this *Codec) serialize(root *TreeNode) string {
// 	if root == nil {
// 		return ""
// 	}
// 	queue := []*TreeNode{root}
// 	collect := []string{}
// 	for len(queue) != 0 {
// 		if queue[0] != nil {
// 			collect = append(collect, strconv.Itoa(queue[0].Val))
// 			queue = append(queue, queue[0].Left, queue[0].Right)
// 		} else {
// 			collect = append(collect, "null")
// 		}
// 		queue = queue[1:]
// 	}

// 	r := strings.Join(collect, ",")
// 	fmt.Println(r)
// 	return r
// }

// // Deserializes your encoded data to tree.
// func (this *Codec) deserialize(data string) *TreeNode {
// 	if data == "" {
// 		return nil
// 	}
// 	collect := strings.Split(data, ",")
// 	v, _ := strconv.Atoi(collect[0])
// 	root := &TreeNode{Val: v}
// 	queue := []*TreeNode{root}
// 	collect = collect[1:]
// 	for len(queue) > 0 {
// 		parent := queue[0]
// 		leftV, rightV := collect[0], collect[1]
// 		if leftV != "null" {
// 			v, _ := strconv.Atoi(leftV)
// 			left := &TreeNode{Val: v}
// 			queue = append(queue, left)
// 			parent.Left = left
// 		}
// 		if rightV != "null" {
// 			v, _ := strconv.Atoi(rightV)
// 			right := &TreeNode{Val: v}
// 			queue = append(queue, right)
// 			parent.Right = right
// 		}
// 		queue = queue[1:]
// 		collect = collect[2:]
// 	}
// 	return root
// }

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
