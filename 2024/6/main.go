package main

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
