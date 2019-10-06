package tool

import "strconv"

type ArrayTransform struct {
}

/**
 * []int 转 []string
 */
func (it *ArrayTransform) ToString(arr []int) []string {
	strArr := make([]string, 0)
	for _,item := range arr {
		strArr = append(strArr, strconv.Itoa(item))
	}
	return strArr
}
