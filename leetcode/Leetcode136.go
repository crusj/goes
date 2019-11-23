package leetcode

/**
给定一个 非空 整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1

示例 2:

输入: [4,1,2,1,2]
输出: 4
*/
type Question136 struct {
	Input  []int
	Output int
}
type Nil struct{}

//列表法时间复杂度O(n²),空间复杂度O(n)
func (it *Question136) Way1() {
	list := make([]int, 0)
	for _, v := range it.Input {
		find := false
		for i, v2 := range list {
			if v2 == v {
				list = append(list[0:i], list[i+1:]...)
				find = true
				break
			}
		}
		if find == false {
			list = append(list, v)
		}
	}
	if len(list) == 1 {
		it.Output = list[0]
	}
}

//哈希表法,时间复杂度O(n)空间复杂度O(n)
func (it *Question136) Way2() {
	hash := make(map[int]Nil)
	for _, v := range it.Input {
		if _, ok := hash[v]; ok {
			delete(hash, v)
		} else {
			hash[v] = Nil{}
		}
	}
	for k, _ := range hash {
		it.Output = k
		return
	}

}

//数学集合法,时间复杂度o(n)空间复杂度O(n)
func (it *Question136) Way3() {
	set := make(map[int]Nil)
	totalSet := 0
	totalInput := 0
	for _, v := range it.Input {
		totalInput += v
		if _, ok := set[v]; !ok {
			set[v] = Nil{}
			totalSet += 2 * v
		}
	}
	it.Output = totalSet - totalInput

}

//位预算xor 按位异或  1 xor 0 = 1,1 xor 1 = 0,时间复杂度o(n) 空间复杂度o(1)
func (it *Question136) Way4() {
	a := 0
	for _, v := range it.Input {
		a = a ^ v
	}
	it.Output = a
}
