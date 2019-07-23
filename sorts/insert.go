package sorts

func InsertSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}
	sortArr := make([]int, length)
	sortArr[0] = arr[0]
	for i := 1; i <= length-1; i++ {
		all := 0
		for j := i - 1; j >= 0; j-- {
			if arr[i] > sortArr[j] {
				all = 1
				sortArr[j+1] = arr[i]
				break
			} else {
				sortArr[j+1] = sortArr[j]
			}
		}

		//最小
		if 0 == all {
			sortArr[0] = arr[i]
		}


	}
	return sortArr
}
