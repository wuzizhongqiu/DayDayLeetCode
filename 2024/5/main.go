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

// 2024_5_23 找出最长等值子数组（哈希，滑窗）
func longestEqualSubarray(nums []int, k int) int {
	pos := make([][]int, len(nums)+1)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	ans := 1
	for _, ps := range pos {
		l, r, n := 0, 0, len(ps)
		for r < n {
			// 下标距离 - 正常距离 = 非该值的数的个数（空隙的大小）
			for ps[r]-ps[l]-(r-l) > k {
				l++
			}
			ans = max(ans, r-l+1)
			r++
		}
	}
	return ans
}

// 2024_5_24 找出最具竞争力的子序列（栈，模拟，贪心）
func mostCompetitive(nums []int, k int) []int {
	st := []int{}
	for i, v := range nums {
		// 保证下一个子序列数 <= 当前子序列数（否则就有更优解，出栈）
		// 保证 nums 剩下的数足够 k 个
		for len(st) > 0 && st[len(st)-1] > v && len(nums)-i > k-len(st) {
			st = st[:len(st)-1]
		}
		if len(st) < k {
			st = append(st, v)
		}
	}
	return st
}
