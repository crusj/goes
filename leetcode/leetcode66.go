package leetcode

import "fmt"

/**
给定一个 非负整数 组成的 非空 数组，在该数的基础上加一，返回一个新的数组。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
type Question66 struct {
	Input  []int
	Output []int
}

func (it *Question66) Deal() {
	index := len(it.Input) - 1
	it.Output = make([]int, len(it.Input))
	copy(it.Output, it.Input)
	for index >= 0 && it.Input[index]+1 == 10 {
		it.Output[index] = 0
		index--
	}
	if index == -1 {
		it.Output[0] = 1
		for i := 1; i <= len(it.Output)-1; i++ {
			it.Output[i] = 0
		}
		it.Output = append(it.Output, 0)
	} else {
		it.Output[index] = it.Output[index] + 1
	}
}
func (it *Question66) String() string {
	return fmt.Sprintf("输入数组为%v，输出为:%v", it.Input, it.Output)
}
