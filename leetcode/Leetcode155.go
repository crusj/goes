package leetcode

import "errors"

/**
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); --> 返回 -3.
minStack.pop();
minStack.top(); --> 返回 0.
minStack.getMin(); --> 返回 -2.
*/
type Question155 struct {
	Stack    []int
	MinStack []int
}

func (it *Question155) Push(x int) {
	it.Stack = append(it.Stack, x)
	minStackLen := len(it.MinStack)
	if minStackLen == 0 || it.MinStack[minStackLen-1] >= x {
		it.MinStack = append(it.MinStack, x)
	}
}
func (it *Question155) Pop() (int, error) {
	stackLen := len(it.Stack)
	if stackLen > 0 {
		ret := it.Stack[stackLen-1]
		it.Stack = it.Stack[0 : stackLen-1]
		minStackLen := len(it.MinStack)
		if ret == it.MinStack[minStackLen-1] {
			it.MinStack = it.MinStack[0 : minStackLen-1]
		}
		return ret, nil
	} else {
		return 0, errors.New("nil stack")
	}

}
func (it *Question155) Top() (int, error) {
	stackLen := len(it.Stack)
	if stackLen == 0 {
		return 0, errors.New("nil stack")
	} else {
		return it.Stack[stackLen-1], nil
	}

}
func (it *Question155) GetMin() (int, error) {
	minStackLen := len(it.MinStack)
	if minStackLen == 0 {
		return 0, errors.New("nil stack")
	}
	return it.MinStack[minStackLen-1], nil
}
