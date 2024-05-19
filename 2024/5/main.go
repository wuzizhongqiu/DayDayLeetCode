package main

import "sort"

// 2024_5_18 找出可整除性得分最大的整数（排序、暴力）
func maxDivScore(nums []int, divisors []int) int {
	sort.Ints(divisors)
	minV, cntMax := divisors[0], 0
	for _, dv := range divisors {
		cnt := 0
		for _, dn := range nums {
			if dn%dv == 0 {
				cnt++
			}
		}
		if cnt > cntMax {
			cntMax = cnt
			minV = dv
		}
	}
	return minV
}

// 2024_5_19 找出数组游戏的赢家（模拟）
func getWinner(arr []int, k int) int {
	mx := arr[0]
	win := 0
	for i := 1; i < len(arr) && win < k; i++ {
		if mx < arr[i] {
			mx = arr[i]
			win = 0
		}
		win++
	}
	return mx
}
