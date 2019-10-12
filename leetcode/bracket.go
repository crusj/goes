package leetcode

/*
给定一个只包括 '(' ， ')' ， '{' ， '}' ， '[' ， ']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true

示例 2:

输入: "()[]{}"
输出: true

示例 3:

输入: "(]"
输出: false

示例 4:

输入: "([)]"
输出: false

示例 5:

输入: "{[]}"
输出: true
*/
type Bracket struct {
	Input   string
	IsValid bool
}

func (it *Bracket) Deal() {
	inputLen := len(it.Input)
	if inputLen == 0 {
		it.IsValid = true
		return
	}
	if inputLen == 1 {
		it.IsValid = false
		return
	}
	stack := make([]uint8, 0)
	rightBracket := []uint8{')', ']', '}'}
	bracketMap := map[uint8]uint8{')': '(', ']': '[', '}': '{'}

	for i := 0; i < inputLen; i++ {
		//最后一位
		match := false
		for _, v := range rightBracket {
			if v == it.Input[i] { //代表是括号的有半部分,需要匹配栈的最后一位
				if len(stack) == 0 || stack[len(stack)-1] != bracketMap[v] { //不等于最后一位
					it.IsValid = false
					return
				} else { //是最后一位
					//出栈
					stack = stack[0 : len(stack)-1]
					match = true
					break
				}
			}
		}
		//未匹配，入栈
		if match != true {
			stack = append(stack, it.Input[i])
		}
	}
	it.IsValid = len(stack) == 0
}
