package graph_test

import (
	"github.com/stretchr/testify/assert"
	. "goes/graph"
	"testing"
)

var (
	sichuan  = new(Node)
	henan    = new(Node)
	hebei    = new(Node)
	tianjing = new(Node)
	shanghai = new(Node)
	beijing  = new(Node)
	queue    Queue
)

func TestBreadFirstSearchIsExists(t *testing.T) {
	sichuan.Name = "四川"
	sichuan.Next = nil
	henan.Name = "河南"
	henan.Next = []*Node{sichuan}
	hebei.Name = "河北"
	hebei.Next = []*Node{sichuan}

	tianjing.Name = "天津"
	tianjing.Next = []*Node{henan}

	shanghai.Name = "上海"
	shanghai.Next = []*Node{sichuan, henan, hebei, tianjing}

	beijing.Name = "北京"
	beijing.Next = []*Node{tianjing, shanghai}

	queue = append(queue, beijing)

	is := BreadFirstSearchIsExists(queue, "河北")
	assert.Equal(t, true, is, "河北存在")

	is = BreadFirstSearchIsExists(queue, "四川")
	//final []string{"北京","上海","四川"}
	assert.Equal(t, true, is)

	is = BreadFirstSearchIsExists(queue, "美国")
	//final []string{"北京","上海","四川"}
	assert.Equal(t, false, is)
}

func TestBreadFirstSearchIsExists2(t *testing.T) {
	sichuan.Name = "四川"
	sichuan.Next = nil
	henan.Name = "河南"
	henan.Next = []*Node{sichuan}
	hebei.Name = "河北"
	hebei.Next = []*Node{sichuan}

	tianjing.Name = "天津"
	tianjing.Next = []*Node{henan}

	shanghai.Name = "上海"
	shanghai.Next = []*Node{sichuan, henan, hebei, tianjing}

	beijing.Name = "北京"
	beijing.Next = []*Node{tianjing, shanghai}

	queue = append(queue, beijing)

	is := BreadFirstSearchIsExists2(queue, "河北")
	assert.Equal(t, true, is, "河北存在")

	is = BreadFirstSearchIsExists2(queue, "四川")
	//final []string{"北京","上海","四川"}
	assert.Equal(t, true, is)

	is = BreadFirstSearchIsExists2(queue, "美国")
	//final []string{"北京","上海","四川"}
	assert.Equal(t, false, is)

}

//最短路径
func TestBreadFirstSearch(t *testing.T) {
	sichuan.Name = "四川"
	sichuan.Next = nil
	henan.Name = "河南"
	henan.Next = []*Node{sichuan}
	hebei.Name = "河北"
	hebei.Next = []*Node{sichuan}

	tianjing.Name = "天津"
	tianjing.Next = []*Node{henan}

	shanghai.Name = "上海"
	shanghai.Next = []*Node{sichuan, henan, hebei, tianjing}

	beijing.Name = "北京"
	beijing.Next = []*Node{tianjing,shanghai}

	var queue Queue
	queue = append(queue, beijing)

	node := BreadFirstSearch(queue, "四川")

	var path []string
	for {
		if node == nil {
			break
		} else {
			path = append(path, node.Name)
		}
		node = node.Pre
	}
	assert.Equal(t, []string{"四川", "上海", "北京"}, path)

}
