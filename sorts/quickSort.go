package sorts

//快速排序
func QuickSort(arr *[]int) {
	length := len(*arr)
	if length <= 1 {
		return
	} else {
		left, right := 0, length-1
		base := (*arr)[0]
	Loop:
		for {
			for {
				if right == left {
					(*arr)[left] = base
					break Loop
				}
				if (*arr)[right] > base {
					right--
				} else {
					(*arr)[left] = (*arr)[right]
					left++
					break
				}
			}
			for {
				if left == right {
					(*arr)[right] = base
					break Loop
				}
				if (*arr)[left] < base {
					left++
				} else {
					(*arr)[right] = (*arr)[left]
					right--
					break
				}
			}
		}
		leftSlice := (*arr)[0:left]
		rightSlice := (*arr)[right+1:]
		QuickSort(&leftSlice)
		QuickSort(&rightSlice)
	}
}

//快速排序2
func QuickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		base := arr[0]
		less := make([]int, 0)
		more := make([]int, 0)
		for i := 1; i < len(arr); i++ {
			if base < arr[i] {
				less = append(less, arr[i])
			} else {
				more = append(more, arr[i])
			}
		}
		return append(append(QuickSort2(less),base),QuickSort2(more)...)
	}

}
