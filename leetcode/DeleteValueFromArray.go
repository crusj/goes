package leetcode

import "fmt"

/**
给定一个数组 nums 和一个值 val ，你需要 原地 移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:

给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

你不需要考虑数组中超出新长度后面的元素。
*/
type DeleteValueFromArray struct {
	Input  []int
	Delete int
	Index  int
}

//双指针
func (it *DeleteValueFromArray) Way1() {
	inputLen := len(it.Input)
	if inputLen == 0 {
		return
	}
	i := 0
	j := 0

	for j < inputLen {
		if it.Input[j] != it.Delete {
			it.Input[i] = it.Input[j]
			i += 1
		}
		j++
	}
	it.Index = i
}

//删除末尾元素
func (it *DeleteValueFromArray) Way2() {
	inputLen := len(it.Input)
	if inputLen < 0 {
		return
	}
	i := 0
	j := inputLen
	for i < j {
		if it.Input[i] == it.Delete {
			it.Input[i] = it.Input[j-1]
			j -= 1
		} else {
			i++
		}
	}
	it.Index = i

}
func (it *DeleteValueFromArray) String() string {
	if it.Index == len(it.Input) {
		return fmt.Sprintf("删除元素%d后的数组是%v", it.Delete, it.Input)
	} else {
		return fmt.Sprintf("删除元素%d后的数组是%v", it.Delete, it.Input[0:it.Index])
	}
}
