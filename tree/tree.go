package tree

import "fmt"

type BinaryTree struct {
}

//前序遍历
func (it *BinaryTree) PreOrder(p *Node) {
	if p != nil {
		fmt.Printf("%d ", p.Value)
		it.PreOrder(p.LChild)
		it.PreOrder(p.RChild)
	}
}

//中序遍历
func (it *BinaryTree) InOrder(p *Node) {
	if p != nil {
		it.PreOrder(p.LChild)
		fmt.Printf("%d ", p.Value)
		it.PreOrder(p.RChild)
	}
}

//后序遍历
func (it *BinaryTree) PostOrder(p *Node) {
	if p != nil {
		it.PreOrder(p.LChild)
		it.PreOrder(p.RChild)
		fmt.Printf("%d ", p.Value)
	}
}

func (it *BinaryTree) String() string {
	return ""
}
