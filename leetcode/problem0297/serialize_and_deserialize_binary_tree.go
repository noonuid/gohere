package problem0297

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/noonuid/go/leetcode/structure"
)

// 297. 二叉树的序列化与反序列化

// 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

// 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structure.TreeNode

type Codec struct {
	null string
}

func Constructor() Codec {
	return Codec{null: "null"}
}

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d", root.Val))
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			sb.WriteString("," + fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			sb.WriteString("," + codec.null)
		}
	}
	return sb.String()
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	values := strings.Split(data, ",")
	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	i, valuesLen := 1, len(values)
	for len(queue) > 0 && i < valuesLen {
		node := queue[0]
		queue = queue[1:]
		if values[i] != codec.null {
			val, _ = strconv.Atoi(values[i])
			left := &TreeNode{Val: val}
			node.Left = left
			queue = append(queue, left)
		}
		i++
		if i < valuesLen && values[i] != codec.null {
			val, _ = strconv.Atoi(values[i])
			right := &TreeNode{Val: val}
			node.Right = right
			queue = append(queue, right)
		}
		i++
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
