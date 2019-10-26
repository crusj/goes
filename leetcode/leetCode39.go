/*
 *给定一个 无重复元素 的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target ）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
[7],
[2,2,3]
]

示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
package leetcode

import "fmt"

type Leetcode39 struct {
	Input  []int
	Target int
	Output [][]int
}

func (it *Leetcode39) Deal() {
	it.al(0, 0, make([]int, 0))
}
func (it *Leetcode39) al(index int, sum int, candidate []int) {
	inputLen := len(it.Input)
	if index == inputLen {
		return
	}
	if sum == it.Target {
		it.Output = append(it.Output, candidate)
		return
	}
	if sum > it.Target {
		return
	}
	it.al(index, sum+it.Input[index], append(candidate, it.Input[index]))
	candidateCopy := make([]int, len(candidate))
	copy(candidateCopy, candidate)
	it.al(index+1, sum, candidateCopy)
}
func (it *Leetcode39) String() string {
	return fmt.Sprintf("Input为%v,target为%d,符合要求的为:%v\n", it.Input, it.Target, it.Output)
}
