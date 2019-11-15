package leetcode

import "strconv"

/**
给定一个非负整数 numRows， 生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
[1],
[1,1],
[1,2,1],
[1,3,3,1],
[1,4,6,4,1]
]
*/
type Question118 struct {
	Input  int
	Output [][]int
}

func (it *Question118) Deal() {
	if it.Input == 1 {
		it.Output = [][]int{
			[]int{1},
		}
	} else if it.Input == 2 {
		it.Output = [][]int{
			[]int{1},
			[]int{1, 1},
		}
	} else {
		it.Output = [][]int{
			[]int{1},
			[]int{1, 1},
		}
		for i := 2; i < it.Input; i++ {
			t := make([]int, 0)
			t = append(t, 1)
			for j := 1; j <= i-1; j++ {
				t = append(t, it.Output[i-1][j]+it.Output[i-1][j-1])
			}
			t = append(t, 1)
			it.Output = append(it.Output,t)
		}
	}
}
func (it *Question118) String() string {
	str := ""
	for _, floor := range it.Output {
		for _, v := range floor {
			str += strconv.Itoa(v)+","
		}
		str += "\n"
	}
	return str
}
