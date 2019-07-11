package sorts

import (
	"fmt"
	"sync"
)

//交换快速排序
func QuickSort1(arr []int, wg *sync.WaitGroup) {
	len := len(arr)
	if len <= 1 {
	} else {
		base := arr[0]
		left := 0
		right := len - 1
		for {
			for arr[right] > base && left != right {
				right--
			}
			for arr[left] <= base && left != right {
				left++
			}
			if left != right {
				arr[left], arr[right] = arr[right], arr[left]
			} else {
				if left != 0 {
					arr[left], arr[0] = base, arr[left]
				}
				break
			}
		}
		wg.Add(2)
		go QuickSort1(arr[0:left], wg)
		go QuickSort1(arr[left+1:], wg)
	}
	wg.Done()
}
func QuickSort3(arr []int) {
	len := len(arr)
	if len <= 1 {
		return
	} else {
		base := arr[0]
		left := 0
		right := len - 1
		for {
			for arr[right] > base && left != right {
				right--
			}
			for arr[left] <= base && left != right {
				left++
			}
			if left != right {
				arr[left], arr[right] = arr[right], arr[left]
			} else {
				arr[left], arr[0] = base, arr[left]
				break
			}

		}
		QuickSort3(arr[0:left])
		QuickSort3(arr[left+1:])
	}
}

func QuickSort2(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	} else {
		less := make([]int, 0)
		more := make([]int, 0)
		base := arr[0]
		for i := 1; i < len; i++ {
			if arr[i] <= base {
				less = append(less, arr[i])
			} else {
				more = append(more, arr[i])
			}
		}
		return append(append(QuickSort2(less), base), QuickSort2(more)...)
	}
}
func PrintR(ch chan string, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)

	}
}
