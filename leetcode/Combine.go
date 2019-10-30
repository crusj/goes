package leetcode

import (
	"fmt"
)

/**
实现组合
*/
type Combine struct {
	Input  []int
	OutPut [][]int
}

func (it *Combine) Deal() {
	it.do(0, make([]int, 0))
}
func (it *Combine) do(next int, arr []int) {
	inputLen := len(it.Input)
	arrLen := len(arr)
	if next > inputLen-1 {
		return
	}
	newArr := append(arr, it.Input[next])
	it.OutPut = append(it.OutPut, newArr)
	arrLen++
	for next+1 <= inputLen-1 {
		tmp := make([]int, arrLen-1)
		copy(tmp, newArr[0:arrLen-1])
		x := append(tmp, it.Input[next+1])
		it.OutPut = append(it.OutPut, x)
		it.do(next+1, newArr)
		next++
	}

}

func (it *Combine) String() string {
	return fmt.Sprintf("输入的数组为:%v,组合为:%v\n", it.Input, it.OutPut)
}
