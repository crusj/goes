package leetcode

/**
wine := [1,4,2,3]  切片位置的值为红酒价格，从1到N开始每增加一年，价格乘以N倍，每年只能卖一瓶，最左边一瓶或最右边一瓶
求价格最大化
*/
type Wine struct {
	PriceYears map[int][]int
	Order      map[int][][]int
}

//初始化
func NewWine(wine []int) *Wine {
	wineLen := len(wine)
	instance := new(Wine)
	instance.PriceYears = make(map[int][]int, 0)
	instance.Order = make(map[int][][]int, 0)
	for _, v := range wine {
		instance.PriceYears[v] = make([]int, wineLen)
		instance.Order[v] = make([][]int, wineLen)
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
func (it *Wine) WayDp(wine []int, year int) int {
	wineLen := len(wine)
	if wineLen > 1 {
		a := it.PriceYears[wine[0]][year-1]
		b := it.PriceYears[wine[wineLen-1]][year-1 ]
		if a == 0 {
			a = wine[0]*year + it.WayDp(wine[1:], year+1)
			it.PriceYears[wine[0]][year-1] = a
		}
		if b == 0 {
			b = wine[wineLen-1]*year + it.WayDp(wine[0:wineLen-1], year+1)
			it.PriceYears[wine[wineLen-1]][year-1] = b
		}
		if a > b {
			return a
		} else {
			return b
		}
	} else {
		return wine[0] * year
	}
}
