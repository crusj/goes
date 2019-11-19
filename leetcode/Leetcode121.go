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
	Buy    int
	Sale   int
}

func (it *Question121) Deal() {
	v1, v2, v3 := it.max(it.Input, true, 0, 0)
	if v3 > 0 {
		it.Output = v3
		it.Buy = v1
		it.Sale = v2
	} else {
		it.Output = 0
	}

}
func (it *Question121) max(input []int, buy bool, v int, day int) (int, int, int) {
	if buy == true && len(input) == 2 {
		return len(input) - 1, len(input), input[len(input)-1] - input[len(input)-2]
	}
	if buy == false && len(input) == 1 {

		return day, len(input), input[len(input)-1] - v
	}
	//买入
	if buy == true {
		//买入
		a1, a2, a3 := it.max(input[1:], false, input[0], len(it.Input)-len(input)+1)
		//不买
		b1, b2, b3 := it.max(input[1:], true, 0, 0)
		if a3 > b3 {
			return a1, a2, a3
		} else {
			return b1, b2, b3
		}
	} else {
		a := input[0] - v //卖出
		b1, b2, b3 := it.max(input[1:], false, v, day)
		if b3 > a {
			return b1, b2, b3
		} else {
			return day, len(it.Input) - len(input) + 1, a
		}
	}
}
func (it *Question121) Extra(arr []int, num int) int {
	if len(arr) == num  && num == 0{
		return 1
	}
	a := it.Extra(arr[1:], num-1)
	b := it.Extra(arr[1:], num)
	return a + b
}

func (it *Question121) String() string {
	return fmt.Sprintf("最大利润为:%d,买入时间第%d天,卖出时间第%d天", it.Output, it.Buy, it.Sale)
}
