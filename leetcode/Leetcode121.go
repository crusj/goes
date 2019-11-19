package leetcode

import "fmt"

/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
type Question121 struct {
	Input  []int
	Output int
}

func (it *Question121) Deal() {
	v := it.max(it.Input, true, 0)
	if v > 0 {
		it.Output = v
	}else{
		it.Output = 0
	}
}
func (it *Question121) max(input []int, buy bool, v int) int {
	if buy == true && len(input) == 2 {
		return input[len(input)-1] - input[len(input)-2]
	}
	if buy == false && len(input) == 1 {

		return input[len(input)-1] - v
	}
	//买入
	if buy == true {
		//买入
		a := it.max(input[1:], false, input[0])
		//不买
		b := it.max(input[1:], true, 0)
		if a > b {
			return a
		} else {
			return b
		}
	} else {
		a := input[0] - v //卖出
		b := it.max(input[1:], false, v)
		if b > a {
			return b
		} else {
			return a
		}
	}
}

func (it *Question121) String() string {
	return fmt.Sprintf("最大利润为:%d", it.Output)
}
