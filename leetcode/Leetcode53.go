package leetcode

import (
	"fmt"
)

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
type Question53 struct {
	Input  []int
	Output int
}

func (it *Question53) Deal() {
	if len(it.Input) <= 0 {
		panic("输入数组不能为空")
	}
	it.Output = it.Input[0]
	sum := 0
	for _, item := range it.Input {
		if sum > 0 {
			sum += item
		} else {
			sum = item
		}
		if it.Output < sum {
			it.Output = sum
		}
	}
}
func (it *Question53) String() string {
	return fmt.Sprintf(" 输入数组为:%v,最大连续子序和为:%d", it.Input, it.Output)
}
