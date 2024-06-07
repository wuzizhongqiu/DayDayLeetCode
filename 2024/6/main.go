package main

import "slices"

// 2024_6_1 给小朋友们分糖果 I
func distributeCandies(n int, limit int) (ans int) {
	// 一刻都没有为被蓝桥杯爆杀而哀悼
	// 我将立刻返回我忠实的 LeetCode 每日一题
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i+j > n {
				break
			}
			if n-i-j <= limit {
				ans++
			}
		}
	}
	return ans
}

// 2024_6_2 分糖果（模拟、计数）
func distributeCandies2(candyType []int) int {
	mp := map[int]int{}
	for _, v := range candyType {
		mp[v]++
	}
	return min(len(candyType)/2, len(mp))
}

// 2024_6_3 分糖果 II（模拟）
func distributeCandies3(candies int, num_people int) []int {
	ans := make([]int, num_people)
	num, idex := 1, 0
	for candies > 0 {
		i := idex % (num_people)
		if candies > num {
			ans[i] += num
			candies -= num
			num++
		} else {
			ans[i] += candies
			break
		}
		idex++
	}
	return ans
}

// 2024_6_4 在带权树网络中统计可连接服务器对数目（图、搜索、乘法原理（数学））
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}
	ans := make([]int, n)
	for i, gi := range g {
		var cnt int
		var dfs func(int, int, int)
		dfs = func(x, fa, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, e := range g[x] {
				if e.to != fa {
					dfs(e.to, x, sum+e.wt)
				}
			}
			return
		}
		sum := 0
		for _, e := range gi {
			cnt = 0
			dfs(e.to, i, e.wt)
			ans[i] += cnt * sum
			sum += cnt
		}
	}
	return ans
}

// 2024_6_4 将元素分配到两个数组中 II（二分、离散化、树状数组）

// 树状数组
type fenwick []int

// 维护 [1, i] 的元素个数
func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

// 获取 [1, i] 的元素个数和
func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res
}

func resultArray(nums []int) []int {
	// 排序去重 -> 离散化
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)
	a, b := []int{nums[0]}, []int{nums[1]}
	// 维护树状数组
	t1, t2 := make(fenwick, m+1), make(fenwick, m+1)
	for i, v := range sorted {
		if v == nums[0] {
			t1.add(i + 1)
		}
		if v == nums[1] {
			t2.add(i + 1)
		}
	}
	for _, x := range nums[2:] {
		// 二分查找离散化数组的下标位置
		l, r := 0, len(sorted)
		for l < r {
			mid := (l + r) >> 1
			if sorted[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		v := l + 1
		// greaterCount: 用数组所有元素 - 小于等于 val 元素的数量 = 大于 val 元素的数量
		gc1 := len(a) - t1.pre(v)
		gc2 := len(b) - t2.pre(v)
		if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
			a = append(a, x)
			t1.add(v)
		} else {
			b = append(b, x)
			t2.add(v)
		}
	}
	return append(a, b...)
}

// 2024_6_6 区分黑球与白球（模拟）
func minimumSteps(s string) (ans int64) {
	cnt := int64(0)
	for _, v := range s {
		if v == '1' {
			cnt++
		} else {
			ans += cnt
		}
	}
	return ans
}

// 2024_6_7 相同分数的最大操作数目 I（模拟）
func maxOperations(nums []int) int {
	ans := 1
	sum := nums[0] + nums[1]
	for i := 3; i < len(nums) && nums[i-1]+nums[i] == sum; i += 2 {
		ans++
	}
	return ans
}