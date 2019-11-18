package leetcode

import (
	"fmt"
	"github.com/crusj/goes/tool"
)

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
[2],
[3,4],
[6,5,7],
[4,1,8,3]
]

自顶向下的最小路径和为 11 （即， 2 + 3 + 5 + 1 = 11）。

说明：

如果你可以只使用 O ( n ) 的额外空间（ n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/
type Question120 struct {
	Input [][]int
	Sum   int
	Path  []int
}

func (it *Question120) Deal() {
	step := make([]int, 1)
	step[0] = it.Input[0][0]

	//默认路径,和值
	for _, v := range it.Input {
		it.Path = append(it.Path, v[0])
		it.Sum += v[0]
	}

	it.Step(step, 0)

}
func (it *Question120) Step(steps []int, index int) {
	sLen := len(steps)
	//比较
	if (len(it.Input)) == sLen {
		sum := tool.SliceSum(steps)
		if it.Sum > sum {
			it.Sum = sum
			it.Path = steps
		}
	} else {
		//左边一位
		c := make([]int, sLen+1)
		copy(c, steps)
		c[sLen] = it.Input[sLen][index]
		it.Step(c, index)

		//右边一位
		c2 := make([]int, sLen+1)
		copy(c2, steps)
		c2[sLen] = it.Input[sLen][index+1]
		it.Step(c2, index+1)
	}
}
func (it *Question120) String() string {
	return fmt.Sprintf("最短路径为%v,最短路径和为%d", it.Path, it.Sum)
}

/**
1.当时首先感觉像BFS，但BFS是寻找是否存在的问题，于是想对BFS进行变通，失败了
2.其次想考虑回溯法，但因为是回溯针对的是如果有一步不满足条件进行回溯，而此问题是要找到一条连接所有层的路径进行判断，所有不幸
3.最终考虑维护一个直到最后一层的一个路径数组，设置一个初始化路径，初始化总和，每完成一个路径则进行总和比较替换，因为只能关联相邻的位置，所以递归的
时候还需要带上上一个路径节点的索引位置

缺点：每次递归需要增加当前层数大小的一个空间,比较费空间
优化：
*/

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i,_ := range triangle {
		dp[i] = make([]int,len(triangle[i]))
	}
	triangleIndex := make([]int,len(triangle))
	triangleIndexLen := len(triangle)
	for i,v := range triangle {
		triangleIndex[i] = len(v)
	}
	return Step(0, 0, triangle,triangleIndex,&triangleIndexLen, dp)

}
func Step(row int, col int, triangle [][]int,triangleIndex []int,triangleIndexLen *int, dp [][]int) int {

	if *triangleIndexLen == row || triangleIndex[row] == col {
		return 0
	}
	if  dp[row][col] != 0 {
		return dp[row][col]
	}
	a := Step(row+1, col, triangle,triangleIndex,triangleIndexLen,dp)
	b := Step(row+1, col+1, triangle,triangleIndex,triangleIndexLen, dp)
	if a > b {
		dp[row][col] = b + triangle[row][col]
	}else{
		dp[row][col] = a + triangle[row][col]
	}
	return dp[row][col]
}
