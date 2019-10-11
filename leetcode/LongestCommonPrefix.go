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

//二分法
func (it *LongestCommonPrefix) Answer4(input []string, lcp string) string {
	inputLen := len(input)
	if inputLen == 0 {
		return lcp
	}
	//找出长度最短的字符串并和第一位交换未知
	for i := 1; i < inputLen; i++ {
		if len(input[i]) < len(input[0]) {
			input[i], input[0] = input[0], input[i]
		}
	}
	pivotLen := len(input[0])

	if pivotLen == 0 {
		return lcp
	}
	if pivotLen == 1 {
		maybeLcp := lcp + input[0]
		if it.isLSP(input[1:], maybeLcp) {
			return maybeLcp
		} else {
			return lcp
		}
	}
	//中间
	middle := pivotLen / 2
	lcpWithMiddleLeft := lcp + input[0][0:middle]
	if it.isLSP(input[1:], lcpWithMiddleLeft) {
		input[0] = input[0][middle:]
		return it.Answer4(input, lcpWithMiddleLeft)
	} else {
		input[0] = input[0][0:middle]
		return it.Answer4(input, lcp)
	}

}

/**
 * 比较字符串是否为数组的最长公共前缀
 */
func (it *LongestCommonPrefix) isLSP(input []string, lsp string) bool {
	lspLen := len(lsp)
	for i := 0; i < len(input); i++ {
		if lspLen > len(input[i]) {
			return false
		} else if lspLen == len(input[i]) {
			if lsp != input[i] {
				return false
			} else {
				continue
			}
		} else {
			if lsp != input[i][0:lspLen] {
				return false
			} else {
				continue
			}
		}
	}
	return true
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
				rsl = a[0:i]
				return
			}
		}
	}
	return a
}
func (it *LongestCommonPrefix) String() string {
	return fmt.Sprintf("输入的字符串数组为%v,最长的公共前缀为:%s\n", it.Input, it.Output)
}
