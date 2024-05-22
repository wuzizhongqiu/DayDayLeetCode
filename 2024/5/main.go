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

// 2024_5_21 找出最大的可达成数字（阅读理解）
func theMaximumAchievableX(num int, t int) int {
	return num + t*2
}

// 2024_5_22 找出输掉零场或一场比赛的玩家（哈希，排序）
func findWinners(matches [][]int) [][]int {
	mpWin := map[int]int{}
	for _, v := range matches {
		if mpWin[v[0]] == 0 {
			mpWin[v[0]] = 0
		}
		mpWin[v[1]]++
	}
	ans := make([][]int, 2)
	for k, v := range mpWin {
		if v == 0 {
			ans[0] = append(ans[0], k)
		}
		if v == 1 {
			ans[1] = append(ans[1], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
