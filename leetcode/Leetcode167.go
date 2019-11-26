package leetcode

/**
给定一个已按照 升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2 。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
*/
type Question167 struct {
	Input []int
	Sum   int
	Index [2]int
}

//有三种方法，
// 暴力法，时间复杂度O(n²)空间复杂度O(1)
// Hash表法,时间复杂度O(n),空间复杂度O(n)
// 双指针法,时间复杂度O(n),空间复杂度O(1)
func (it *Question167) Deal() {
	inputLen := len(it.Input)
	if inputLen <= 1 {
		panic("数组长度至少为2")
	}
	left, right := 0, inputLen-1
	for left < right {
		sum := it.Input[left] + it.Input[right]
		if sum == it.Sum {
			it.Index[0] = left + 1
			it.Index[1] = right + 1
			return
		} else if sum < it.Sum {
			left++
		} else {
			right--
		}
	}
	panic("不存在")
}
