package main

import (
	"fmt"
	"math/rand"
	"sort/sorts"
	"time"
)

func generateRandNums(max, number int) (re []int) {
	if max <= 0 {
		max = 100
	}

	for i := 0; i < number; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		t := r.Intn(max)
		re = append(re, t)
	}
	return

}

func main() {
	oldNums := generateRandNums(1000, 200000)
	oldNumsCopy := make([]int, len(oldNums))
	oldNumsCopy2 := make([]int, len(oldNums))

	copy(oldNumsCopy, oldNums)
	copy(oldNumsCopy2, oldNums)

	start := time.Now()
	sorts.Bubble(&oldNums)
	fmt.Printf("冒泡排序运行时间%v秒\n", time.Since(start))

	start = time.Now()
	sorts.QuickSort2(oldNumsCopy)
	fmt.Printf("快速排序运算运行时间%v秒\n", time.Since(start))

	start = time.Now()
	sorts.QuickSort(&oldNumsCopy2)
	fmt.Printf("快速排序2运算运行时间%v\n", time.Since(start))
}
