package leetcode

import "fmt"

/**
假设你正在爬楼梯。需要 n 步你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意： 给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1. 1 步 + 1 步
2. 2 步

示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1. 1 步 + 1 步 + 1 步
2. 1 步 + 2 步
3. 2 步 + 1 步

1.第一印象会用到递归
2.关键词 共n步 一步 二步
3.分析

最小长度为N/2 最大长度为N
*/
type Question70 struct {
	Input  int
	Output int
}

func (it *Question70) Deal() {
	it.Output = it.fbly(it.Input)
}
func (it *Question70) fbly(i int) int {
	if i == 1 {
		return 1
	}
	if i == 2 {
		return 2
	}
	return it.fbly(i-1) + it.fbly(i-2)
}
func (it *Question70) String() string {
	return fmt.Sprintf("楼梯为%d层,方案有%d种", it.Input, it.Output)
}
