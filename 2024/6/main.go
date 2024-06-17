package main

import (
	"slices"
	"sort"
)

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

// 2024_6_8 相同分数的最大操作数目 II
func maxOperations2(nums []int) int {
	n := len(nums)
	// 选择 nums 中最前面两个元素并且删除它们
	res1, finish := add(nums[2:], nums[0]+nums[1])
	if finish == true {
		return n / 2
	}
	// 选择 nums 中最后两个元素并且删除它们
	res2, finish := add(nums[:n-2], nums[n-1]+nums[n-2])
	if finish == true {
		return n / 2
	}
	// 选择 nums 中第一个和最后一个元素并且删除它们
	res3, finish := add(nums[1:n-1], nums[0]+nums[n-1])
	if finish == true {
		return n / 2
	}
	return max(res1, res2, res3) + 1
}

func add(a []int, target int) (res int, finish bool) {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if finish == true {
			return res
		}
		if i >= j {
			finish = true
			return res
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if a[i]+a[i+1] == target {
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j]+a[j-1] == target {
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target {
			res = max(res, dfs(i+1, j-1)+1)
		}
		memo[i][j] = res
		return res
	}
	res = dfs(0, n-1)
	return res, finish
}

// 2024_6_9 戳气球（记忆化搜索）
func maxCoins(nums []int) int {
	n := len(nums)
	val, memo := map[int]int{}, make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo {
			memo[i][j] = -1
		}
	}
	for i, v := range nums {
		val[i] = v
	}
	val[-1], val[n] = 1, 1
	var dfs func(int, int) int
	dfs = func(i, j int) (ans int) {
		if i > j {
			return ans
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		for k := i; k <= j; k++ {
			ans = max(ans, nums[k]*val[i-1]*val[j+1]+dfs(i, k-1)+dfs(k+1, j))
		}
		memo[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}

// 2024_6_10 救生艇（贪心/二分）
func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)
	for i, v := range people {
		empty := limit - v
		if empty >= v {
			// 试试二分
			l, r := i+1, len(people)-1
			for l < r {
				mid := (l + r + 1) >> 1
				if people[mid] > empty {
					r = mid - 1
				} else {
					l = mid
				}
			}
			if people[l] <= empty {
				people[l] = limit + 1
			}
		}
		if v > limit {
			continue
		}
		ans++
	}
	return ans
}

// 2024_6_11 甲板上的战舰（模拟）
func countBattleships(board [][]byte) (ans int) {
	for i := range board {
		for j := range board[i] {
			if (board[i][j] == 'X') &&
				(i == 0 || board[i-1][j] != 'X') &&
				(j == 0 || board[i][j-1] != 'X') {
				ans++
			}
		}
	}
	return ans
}

// 2024_6_13 子序列最大优雅度（排序、贪心）
func findMaximumElegance(items [][]int, k int) int64 {
	slices.SortFunc(items, func(a, b []int) int {
		return b[0] - a[0]
	})
	ans, totalProfit := 0, 0
	vis, duplicate := map[int]bool{}, []int{} // 重复类的判断与利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if vis[category] == false {
				vis[category] = true
			} else { // 这个是重复的类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && vis[category] == false { // 不是重复类别
			vis[category] = true
			totalProfit += profit - duplicate[len(duplicate)-1] // 替换利润最低的重复类别
			duplicate = duplicate[:len(duplicate)-1]
		}
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}

// 2024_6_14 访问数组中的位置使分数最大（记忆化搜索/DP）
func maxScore(nums []int, x int) int64 {
	n := len(nums)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == n {
			return res
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() {
			*p = res
		}()
		// 下一个选择奇偶性相同的序列数
		if nums[i]%2 != j {
			return dfs(i+1, j)
		}
		// 下一个选择奇偶性不同的序列数（不选可能结果更大）
		return max(dfs(i+1, j), dfs(i+1, j^1)-x) + nums[i]
	}
	return int64(dfs(0, nums[0]%2))
}

// 2024_6_15 数组的最大美丽值（排序、滑动窗口）
func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums)
	l := 0
	for r, v := range nums {
		for v-nums[l] > k*2 {
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

// 2024_6_17 最长特殊序列 II（排序，暴力匹配）
func findLUSlength(strs []string) int {
	// 降序排序
	slices.SortFunc(strs, func(s1, s2 string) int {
		return len(s2) - len(s1)
	})
next:
	for i, v1 := range strs {
		for j, v2 := range strs {
			if i != j && isSub(v1, v2) {
				continue next
			}
		}
		return len(v1)
	}
	return -1
}

// 判断 v1 是否是 v2 的子序列
func isSub(v1, v2 string) bool {
	i := 0
	for _, v := range v2 {
		if v1[i] == byte(v) {
			i++
		}
		if i == len(v1) {
			return true
		}
	}
	return false
}
