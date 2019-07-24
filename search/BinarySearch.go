package search

func BinarySearch(arr []int, value int) bool {
	length := len(arr)
	if length == 0 {
		return false
	}
	middleIndex := length / 2
	if arr[middleIndex] == value {
		return true
	} else {

		return BinarySearch(arr[0:middleIndex], value) || BinarySearch(arr[middleIndex+1:], value)
	}

}

// 1 2 3 4 5  2
