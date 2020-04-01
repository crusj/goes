package dp

import "math"

//先看下题目：给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
var coinTable = make(map[int]int)

func MinCoinAmount(list []int, coin int) int {
	res := math.MaxFloat64
	if coin == 0 {
		return 0
	}
	if coin < 0 {
		return -1
	}
	for _, v := range list {
		sub := -1
		if value, exist := coinTable[coin-v]; exist {
			sub = value
		} else {
			sub = MinCoinAmount(list, coin-v) //子问题
			coinTable[coin-v] = sub
		}
		if sub == -1 {
			continue
		}
		res = math.Min(res, 1+float64(sub))
	}
	return int(res)
}
