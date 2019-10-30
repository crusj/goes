package leetcode

import (
	"fmt"
	"github.com/crusj/goes/tool"
	"strconv"
	"strings"
)

/**
 给定一个整数数组和一个目标值，找出数组中和为目标值的 两个 数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
type TwoSum struct {
	Input  []int
	Output [][]int
	Search int
}

func (it *TwoSum) Deal(way int) {
	if len(it.Input) <= 0 {
		return
	} else {
		switch way {
		case 1: //暴力法
			it.violence()
		case 2: //两遍hash
			it.hashTwice()
		case 3: //一遍哈希法
			it.hashOnce()
		default:
			fmt.Sprintln("方法错误")
			return
		}
		fmt.Println(it)
	}

}
func (it *TwoSum) Empty() {
	it.Output = [][]int{}
}

//暴力查找
func (it *TwoSum) violence() {
	for i := 0; i <= len(it.Input)-1; i++ {
		current := it.Input[i]
		for j := i + 1; j <= len(it.Input)-1; j++ {
			if current+it.Input[j] == it.Search {
				it.Output = append(it.Output, []int{current, it.Input[j]})
			}
		}
	}
}

//两边哈希法,只能针对input为集合,因为非集合存在相同元素后者会覆盖前者,查找元素不能是自己
func (it *TwoSum) hashTwice() {
	inputHash := make(map[int]int)
	//存入hash表
	for i, value := range it.Input {
		inputHash[value] = i
	}
	//查找
	for _, item := range it.Input {
		if _, ok := inputHash[it.Search-item]; ok && it.Search-item != item {
			it.Output = append(it.Output, []int{item, it.Search - item})
			return
		}
	}
}

//一遍hash法,可以针对非集合
func (it *TwoSum) hashOnce() {
	inputHash := make(map[int]int)
	for i, value := range it.Input {
		if _, ok := inputHash[it.Search-value]; ok {
			it.Output = append(it.Output, []int{it.Search - value, value})
			return
		} else {
			inputHash[value] = i
		}
	}
}

//显示结果
func (it *TwoSum) String() string {
	strSlice := make([]string, 0)
	for _,item := range it.Input {
		strSlice = append(strSlice, strconv.Itoa(item))
	}
	outStrSlice := make([]string, 0)

	t := &tool.ArrayTransform{};
	for _, item := range it.Output {
		outStrSlice = append(outStrSlice, "["+strings.Join(t.ToString(item), ",")+"]")
	}
	return fmt.Sprintf("输入数组为[%s],search为%d,符合要求的数据为%s\n", strings.Join(strSlice, ","), it.Search, outStrSlice)
}
