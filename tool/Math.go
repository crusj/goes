package tool

//求和
func SliceSum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

//递归求和
func SliceSumR(arr []int) (sum int) {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + SliceSumR(arr[1:])
}
