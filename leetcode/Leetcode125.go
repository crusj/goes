package leetcode

import (
	"fmt"
	"strings"
	"unicode"
)

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明： 本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:

输入: "race a car"
输出: false
*/
type Question125 struct {
	Input  string
	Output bool
}

func (it *Question125) Deal() {
	inputLen := len(it.Input)
	it.Input = strings.ToLower(it.Input)
	if inputLen == 0 {
		it.Output = true
		return
	}
	left := 0
	right := inputLen - 1
	for {
		for left != right {
			if unicode.IsLetter(rune(it.Input[left])) || unicode.IsNumber(rune(it.Input[left])) {
				break
			}
			left++
		}
		for left != right {
			if unicode.IsLetter(rune(it.Input[right])) || unicode.IsNumber(rune(it.Input[right])) {
				break
			}
			right--

		}
		if left == right {
			it.Output = true
			return
		} else {
			if it.Input[left] == it.Input[right] {
				left++
				right--
				if left > right {
					it.Output = true
					return
				}
			} else {
				it.Output = false
				return
			}
		}
	}
}
func (it *Question125) String() string {
	return fmt.Sprintf("%s是回文串?%b", string(it.Input), it.Output)
}
