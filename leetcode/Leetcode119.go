package leetcode

import "strconv"

/**
给定一个非负索引 k ，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]

进阶：

你可以优化你的算法到 O ( k ) 空间复杂度吗？
*/
type Question119 struct {
	Input  int
	Output []int
}

func (it *Question119) Deal() {
	if it.Input == 1 {
		it.Output = []int{1}
	} else if it.Input == 2 {
		it.Output = []int{1, 1}
	} else {
		it.Output = append(it.Output, 1) //第一位为1
		for i := 1; i <= it.Input-2; i++ {
			it.Output = append(it.Output, it.V(i, it.Input))
		}
		it.Output = append(it.Output, 1) //最后一位为1
	}
}
func (it *Question119) V(index int, floor int) int {
	if index == 0 || index == floor-1 {
		return 1
	} else {
		return it.V(index, floor-1) + it.V(index-1, floor-1)
	}
}
func (it *Question119) String() string {
	a := ""
	for _, v := range it.Output {
		a += strconv.Itoa(v) + ","
	}
	return a
}
