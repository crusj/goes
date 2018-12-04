package graph

// 广度优先算法
//解决1.从A到B是否存在可行路径?
//解决2.从A到B最短步数路径是什么?

type (
	//节点
	Node struct {
		Name string
		Next []*Node
		Pre  *Node
	}
	Queue []*Node
)

//is exists
func BreadFirstSearchIsExists(queue Queue, find string) bool {
	for _, item := range queue {
		if item.Name == find {
			return true
		} else if item.Next == nil {
			continue
		} else {
			//入队
			queue := append(queue, item.Next...)
			//当前出队
			return BreadFirstSearchIsExists(queue[1:], find)
		}
	}
	return false
}

//is exists 2
func BreadFirstSearchIsExists2(queue Queue, find string) bool {
	if len(queue) <= 0 {
		return false
	}
	readNode := make(map[string]string, 0)//已经入队
	for {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		if current.Name == find {
			return true
		} else {
			for _, item := range current.Next {
				//不存在则入队
				if _, ok := readNode[item.Name]; !ok {
					queue = append(queue, item)
				}
			}
		}
		//出队
		queue = queue[1:]
	}
	return false
}

//least step
func BreadFirstSearch(queue Queue, find string) *Node{
	if len(queue) <= 0 {
		return nil
	}
	readNode := make(map[string]string, 0)//已经入队

	for {
		if len(queue) <= 0 {
			break
		}
		current := queue[0]
		if current.Name == find {
			return current
		} else {
			for _, item := range current.Next {
				//不存在则入队
				if _, ok := readNode[item.Name]; !ok {
					readNode[item.Name] = "" //非常重要，如果不对是否如过对行判断，如果某个节点被不止一个节点指向的时候，会被改变pre指针指向
					item.Pre = current
					queue = append(queue, item)
				}
			}
		}
		//出队
		queue = queue[1:]
	}
	return nil

}
