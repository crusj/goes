package tree

import "fmt"

type BinaryTree struct {
}

//前序遍历
func (it *BinaryTree) PreOrder(p *Node) {
	if p != nil {
		fmt.Printf("%v ", p.Value)
		it.PreOrder(p.LChild)
		it.PreOrder(p.RChild)
	}
}

//中序遍历
func (it *BinaryTree) InOrder(p *Node) {
	if p != nil {
		it.InOrder(p.LChild)
		fmt.Printf("%v ", p.Value)
		it.InOrder(p.RChild)
	}
}

//后序遍历
func (it *BinaryTree) PostOrder(p *Node) {
	if p != nil {
		it.PostOrder(p.LChild)
		it.PostOrder(p.RChild)
		fmt.Printf("%v ", p.Value)
	}
}

//前序遍历创建二叉树
func (it *BinaryTree) CreateByArr(arr []string, index *int) *Node {
	if len(arr) == *index { //完成
		return nil
	}
	if arr[*index] == "#" {
		*index++
		return nil
	}
	node := &Node{Value: arr[*index]}
	*index++
	node.LChild = it.CreateByArr(arr, index)
	node.RChild = it.CreateByArr(arr, index)
	return node
}
