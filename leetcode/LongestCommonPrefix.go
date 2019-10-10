package leetcode

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 "" 。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"

示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:

所有输入只包含小写字母 a-z 。
*/
type LongestCommonPrefix struct {
	Input  []string
	Output string
}

func (it *LongestCommonPrefix) Deal() {

}

//解法1水平扫描法
func (it *LongestCommonPrefix) Answer1() {
	inputLen := len(it.Input)
	if inputLen == 0 {
		it.Output = ""
		return
	} else {
		it.Output = it.Input[0] //假设数组第一个元素就是最长公共前缀
		for i := 1; i < inputLen; i++ {
			it.Output = it.CommonPrefix(it.Output, it.Input[i])
			if it.Output != "" {
				continue
			} else {
				break
			}
		}
	}
	fmt.Println(it)

}

//解法2垂直扫描法
func (it *LongestCommonPrefix) Answer2() {
	inputLen := len(it.Input)
	lenArr := make([]int, inputLen)
	if inputLen < 1 {
		it.Output = ""
		return
	}
	//记录每个字符串的长度
	for i := 0; i < inputLen; i++ {
		lenArr[i] = len(it.Input[i])
	}
	for i := 0; ; i++ {
		//首先去第一个字符串第一个字符作为最长公共前缀
		tmp := ""
		if lenArr[0] < i+2 { //防止溢出
			tmp = it.Input[0][i:]
		} else {
			tmp = it.Input[0][i : i+1]
		}
		for j := 1; j < len(lenArr)-1; j++ {
			if lenArr[j] > i {
				current := ""
				if lenArr[j] > i+1 {
					current = it.Input[j][i : i+1]
				} else {
					current = it.Input[j][i:]
				}
				if current == tmp {
					continue
				} else {
					return
				}
			} else {
				return
			}
		}
		it.Output += tmp
	}

}

//分治法
func (it *LongestCommonPrefix) Answer3(input []string) (rsl string) {
	inputLen := len(input)
	if inputLen == 0 {
		rsl = ""
	} else if inputLen == 1 {
		rsl = input[0]
	} else if inputLen == 2 {
		rsl = it.CommonPrefix(input[0], input[1])
	} else {
		halfInputLen := inputLen / 2
		rsl = it.CommonPrefix(it.Answer3(input[0:halfInputLen+1]), it.Answer3(input[halfInputLen+1:]))
	}
	return
}

//求两个字符串的公共前缀
func (it *LongestCommonPrefix) CommonPrefix(a string, b string) (rsl string) {
	aLen := len(a)
	bLen := len(b)
	for i := 0; i < aLen; i++ {
		if i >= bLen {
			rsl = b[0:]
			return
		} else {
			if a[i] == b[i] {
				continue
			} else {
				rsl = a[0 : i]
				return
			}
		}
	}
	return a
}
func (it *LongestCommonPrefix) String() string {
	return fmt.Sprintf("输入的字符串数组为%v,最长的公共前缀为:%s\n", it.Input, it.Output)
}
