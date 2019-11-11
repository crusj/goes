package leetcode

import (
	"fmt"
	"math"
)

/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2

示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842...,
  由于返回类型是整数，小数部分将被舍去。
等于，和介于大小之间的

方案1：从1开始遍历递增求平方，如果数很大，则效率低下
方案2：1.能找到一个最大值根号值，即使这个值不预估值要大一些,2.近视二分法

1. 2 3 4 5 6 7 8 9 10   100   50
5 =》 25
==
2. 6 7 8 9 10
8 => 64
==

100
log2 100 = 10 10000
1000
log2 100 = 10 10000
10000
log2 1000 = 30 1000000
100000
Log2 1000 = 30 1000000

位数为偶数和奇数
如果是偶数位，则长度为偶数位/2 个0
如果是奇数位，长度为奇数位/2 + 1 个0
*/
type Question69 struct {
	Input  int
	Output int
}

func (it *Question69) Deal() {
	max := 0
	t := it.Input
	for t > 0 {
		max++
		t = t / 10
	}
	if max <= 0 {
		it.Output = 0
	}
	zero := 0
	if max%2 == 0 {
		zero = int(max / 2);
	} else {
		zero = int(max / 2);
		zero = zero + 1
	}
	max = int(math.Pow10(zero))
	it.Output = it.searchSqrt(1, max)
}
func (it *Question69) searchSqrt(start int, end int) int {
	if start == end {
		return start
	} else {
		half := int((start + end) / 2)
		if half*half == it.Input {
			return half
		} else if half*half < it.Input {
			if end >= half+1 {
				if (half+1)*(half+1) > it.Input {
					return half
				} else {
					return it.searchSqrt(half+1, end)
				}
			} else {
				return 0
			}
		} else {
			if start+1 <= half {
				if (half-1)*(half-1) < it.Input {
					return half - 1
				} else {
					return it.searchSqrt(start, half-1)
				}
			} else {
				return 0
			}
		}
	}
}
func (it *Question69) String() string {
	return fmt.Sprintf("%d的平方根是%d", it.Input, it.Output)
}
