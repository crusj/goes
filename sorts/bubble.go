package sorts

//冒泡排序
func Bubble(container []int) {
	len := len(container)
	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			if container[j] < container[j-1] {
				container[j-1], container[j] = container[j], container[j-1]
			}
		}
	}
}
