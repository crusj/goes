package leetcode

import (
	"fmt"
	"strconv"
)

/**
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0 。

示例 1:

输入: a = "11", b = "1"
输出: "100"

示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
type Question67 struct {
	Input   []string
	OutputD int
	OutputB string
}

func (it *Question67) Deal() {
	d := it.Sum(it.Input)
	it.OutputD = d
	b := it.D2B(d)
	if len(b) <= 0 {
		it.OutputB = ""
		return
	} else {
		for i := b[0]; i >= 0; i-- {
			if len(b) > 0 && i == b[0] {
				it.OutputB += "1"
				if len(b) > 1 {
					b = b[1:]
				} else {
					b = make([]int, 0)
				}
			} else {
				it.OutputB += "0"
			}
		}
	}
}

//1.二进制字符串转10进制
func (it *Question67) B2D(str string) int {
	strLen := len(str)
	if strLen <= 0 {
		return 0
	}
	if first, err := strconv.Atoi(string(str[0])); err == nil {
		firstValue := first * (1 << (strLen - 1))
		if strLen == 1 {
			return firstValue
		} else {
			return firstValue + it.B2D(string(str[1:]))
		}
	} else {
		return 0;
	}
}

//2.求和
func (it *Question67) Sum(strGroup []string) int {
	if len(strGroup) == 0 {
		return 0
	} else {
		return it.B2D(strGroup[0]) + it.Sum(strGroup[1:])
	}
}

//3.10进制转二进制,位便宜
func (it *Question67) D2B(d int) []int {
	ret := make([]int, 0)
	if d == 0 {
		return ret
	} else {
		maxOffset := it.maxOffset(d, 0)
		ret = append(ret, maxOffset)
		return append(ret, it.D2B(d-1<<maxOffset)...)
	}
}

//4.max offset,前提d大于0
func (it *Question67) maxOffset(d int, offset int) int {
	if 2<<offset > d {
		return offset
	} else {
		return it.maxOffset(d, offset+1)
	}
}

func (it *Question67) String() string {
	return fmt.Sprintf("输入的二进制字符串为:%v,二进制求和为:%d,用二进制表示为:%s", it.Input, it.OutputD, it.OutputB)
}
