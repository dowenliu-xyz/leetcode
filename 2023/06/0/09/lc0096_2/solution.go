package lc0096

func numTrees(n int) int {
	// dp 数组。G[i] 表示 i 个节点组成不同二叉搜索树的个数
	G := make([]int, n+1)
	G[0], G[1] = 1, 1 // 由 0 个、 1 个节点构成的情况都只有一种
	for i := 2; i <= n; i++ {
		// G[i] = G[0]*G[i-1] + G[1]*G[i-2] + .. + G[i-1]*G[0]
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
