package tree

type Node struct {
	Value  interface{}
	LChild *Node
	RChild *Node
}

//左子节点
func (it *Node) SetLChild(left *Node) {
	it.LChild = left
}
//右子节点
func (it *Node) SetRChild(right *Node) {
	it.RChild = right
}
