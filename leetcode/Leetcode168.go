package leetcode

/**
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

1 -> A
2 -> B
3 -> C
...
26 -> Z
27 -> AA
28 -> AB
...
示例 1:

输入: 1
输出: "A"

示例 2:

输入: 28
输出: "AB"

示例 3:

输入: 701
输出: "ZY"
*/
type Question168 struct {
	Input  int
	Output string
}

func (it *Question168) Deal() {
	name := make([]byte, 0)
	str := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y',
	}
	t := it.Input
	for t > 0 {
		mod := t % 26
		t = t / 26
		if mod == 0 {
			name = append(name, 'Z')
			t = t - 1
		} else {
			name = append(name, str[mod-1])
		}
	}
	for i := len(name) - 1; i >= 0; i-- {
		it.Output += string(name[i])
	}
}
