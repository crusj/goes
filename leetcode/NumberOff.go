package leetcode

import (
	"fmt"
	"strconv"
)

/**
报数序列是指一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1. 1
2. 11
3. 21
4. 1211
5. 111221

1 被读作 "one 1" ( "一个一" ) , 即 11 。
11 被读作 "two 1s" ( "两个一" ）, 即 21 。
21 被读作 "one 2" ,  " one 1" （ "一个二" , "一个一" ) , 即 1211 。

给定一个正整数 n ，输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:

输入: 1
输出: "1"

示例 2:

输入: 4
输出: "1211"
*/
type NumberOff struct {
	Input  int
	outPut int
}
type Dot struct {
	Value int
	Count int
}

func (it *NumberOff) Deal(i int) int {
	if i <= 0 {
		return -1
	} else if i == 1 {
		return 1
	} else if i == 2 {
		return 11
	} else {
		return it.Say(it.Deal(i - 1))
	}
}
func (it *NumberOff) Say(js int) int {
	h := make([]*Dot, 0)
	for js > 0 {
		len := len(h)
		t := js % 10
		if len == 0 || h[len-1].Value != t {
			h = append(h, &Dot{t, 1})
		} else {
			h[len-1].Count++
		}
		js = js / 10
	}
	retStr := ""
	for i := len(h) - 1; i >= 0; i-- {
		retStr += strconv.Itoa(h[i].Count) + strconv.Itoa(h[i].Value)
	}
	ret, _ := strconv.Atoi(retStr)
	return ret
}
func (it *NumberOff) String() string {
	return fmt.Sprintf("输入的值为%d,输出的值为%d", it.Input, it.outPut)
}
