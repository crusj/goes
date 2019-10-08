package leetcode

import "fmt"

//任意罗马数字字符转换为整数

type RomainToInteger struct {
	Romain        string
	arabic        int
	RomainCombine map[string]int
}

func (it *RomainToInteger) Deal() {
	if it.RomainCombine == nil {
		it.romainHash()
	}
	romainLen := len(it.Romain)
	it.arabic = 0
	for i := 0; i < romainLen; {
		//如果可能先取两位
		if i+1 < romainLen { //此时两位必然存在
			h := ""
			if i+2 < romainLen {
				h = it.Romain[i : i+2]
			} else {
				h = it.Romain[i:]
			}
			if v, ok := it.RomainCombine[h]; ok {
				it.arabic += v
				i = i + 2
				continue
			} else {
				if v, ok := it.RomainCombine[it.Romain[i:i+1]]; ok {
					it.arabic += v
					i += 1
					continue
				} else {
					it.arabic = 0
					break
				}
			}
		}
		//不可取两位或两位组成的罗马数字不存在组合
		if v, ok := it.RomainCombine[it.Romain[i:]]; ok {
			it.arabic += v
			i += 1
		} else {
			it.arabic = 0
			break
		}
	}
	fmt.Println(it)
}

func (it *RomainToInteger) romainHash() {
	it.RomainCombine = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
}
func (it *RomainToInteger) String() string {
	return fmt.Sprintf("输入的罗马数字为%s,对应的整整数位%d\n", it.Romain, it.arabic)
}
