package leetcode

/**
wine := [1,4,2,3]  切片位置的值为红酒价格，从1到N开始每增加一年，价格乘以N倍，每年只能卖一瓶，最左边一瓶或最右边一瓶
求价格最大化
*/
type Wine struct {
	PriceYears [][]int
	Order      map[int][][]int
}

//初始化
func NewWine(wine []int) *Wine {
	wineLen := len(wine)
	instance := new(Wine)
	instance.PriceYears = make([][]int, wineLen)
	for i, _ := range instance.PriceYears {
		instance.PriceYears[i] = make([]int, wineLen)
	}
	return instance
}

//卖的方式
func (it *Wine) Way(wine []int, year int) int {
	wineLen := len(wine)
	if wineLen > 1 {
		a := wine[0]*year + it.Way(wine[1:], year+1)
		b := wine[wineLen-1]*year + it.Way(wine[0:wineLen-1], year+1)
		if a > b {
			return a
		} else {
			return b
		}
	} else {
		return wine[0] * year
	}
}

//动态规划
func (it *Wine) WayDp(wine []int, begin, end int) int {
	year := len(wine) - (end - begin + 1) + 1
	if begin == end {
		return wine[begin] * year
	} else if begin > end {
		return 0
	} else {
		if it.PriceYears[begin][end] != 0 {
			return it.PriceYears[begin][end]
		} else {
			a := it.WayDp(wine, begin+1, end)
			b := it.WayDp(wine, begin, end - 1)
			if a < b {
				it.PriceYears[begin][end] = b + year*wine[end]
			} else {
				it.PriceYears[begin][end] = a + year*wine[begin]
			}
			return it.PriceYears[begin][end]
		}

	}
}
