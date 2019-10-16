package leetcode

import "fmt"

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
//todo KMP算法
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
func (it *Substr) String() string {
	return fmt.Sprintf("hackstack为:%s,needle为:%s,index为:%d", it.Input, it.Search, it.index)
}
