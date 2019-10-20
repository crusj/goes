package leetcode

import (
	"fmt"
)

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1 。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/
type Substr struct {
	Input  string
	Search string
	index  int
}

func (it *Substr) Deal() {
	inputLen := len(it.Input)
	searchLen := len(it.Search)
	if searchLen > inputLen {
		it.index = -1
	} else {
		i := 0
		j := 0
		t := i
		for j < searchLen && i < inputLen {
			if it.Input[i] == it.Search[j] {
				j++
				i++
			} else {
				j = 0
				t++
				i = t
			}
		}
		if j == searchLen {
			it.index = i - searchLen
		} else {
			it.index = -1
		}
	}
}

//KMP算法
func (it *Substr) Kmp() {
	//不匹配位对应的最长公共前后缀长度
	table := make(map[int]int)
	table[0] = -1 //下标为0则最长公共前后缀长度为-1
	for i := 1; i <= len(it.Search)-1; i++ {
		table[i] = it.maxCommonPrefixSuffix(it.Search[0:i])
	}
	i := 0 //主串下标位置
	j := 0 //模式串下标位置
	inputLen := len(it.Input)
	searchLen := len(it.Search)
	init := 0
	for j < searchLen && (inputLen-i >= searchLen-j) {
		if it.Input[i] == it.Search[j] {
			i++
			j++
		} else {
			i = j - table[j] + init
			init = i
			if table[j] >= 0 {//初始化
				j = table[j]
			}else{
				j = 0
			}
		}
	}
	if j == searchLen { //匹配成功
		it.index = i - j
	} else {
		it.index = -1
	}
}

//求字符串的最长公共前后缀
func (it *Substr) maxCommonPrefixSuffix(str string) int {
	strLen := len(str)
	//计算前缀
	prefix := make(map[string]int)
	for i := 0; i < strLen-1; i++ {
		prefix[str[0:i+1]] = i + 1
	}
	//计算后缀
	suffix := make(map[string]int)
	for i := 1; i <= strLen-1; i++ {
		suffix[str[i:]] = strLen - 1
	}
	//找出最长公共前后缀
	longest := 0
	for k, v := range prefix {
		if _, ok := suffix[k]; ok {
			if v > longest {
				longest = v
			}
		}
	}
	return longest
}
func (it *Substr) String() string {
	return fmt.Sprintf("hackstack为:%s,needle为:%s,index为:%d", it.Input, it.Search, it.index)
}
