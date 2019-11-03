package leetcode

import "fmt"

/**
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。

如果不存在最后一个单词，请返回 0 。

说明： 一个单词是指由字母组成，但不包含任何空格的字符串。

示例:

输入: "Hello World"
输出: 5
*/
type Question58 struct {
	Input  string
	Output int
}

func (it *Question58) Deal() {
	//去掉空格
	j := len(it.Input) - 1
	for ; j >= 0; j-- {
		if it.Input[j] != 32 {
			break
		}
	}
	newInput := it.Input
	if j != len(it.Input)-1 {
		newInput = it.Input[0 : j+1]
	}

	//剩余长度
	if len(newInput) == 0 {
		return
	}
	for i := len(newInput) - 1; i >= 0; i-- {
		if newInput[i] == 32 {
			it.Output = len(newInput) - i - 1
			return
		}
	}
	it.Output = len(newInput)
}
func (it *Question58) String() string {
	return fmt.Sprintf("输入字符串为:%s,最后一个单词的长度为%d", it.Input, it.Output)
}
