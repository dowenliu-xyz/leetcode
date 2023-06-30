package lc0200

func numIslands(grid [][]byte) int {
	//return solution1(grid)
	//return solution2(grid)
	return solution3(grid)
}

func solution1(grid [][]byte) int {
	// dfs
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '0' || vis[row][col] {
			return
		}
		vis[row][col] = true
		for _, dir := range directions {
			dfs(row+dir[0], col+dir[1])
		}
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' || vis[row][col] {
				continue
			}
			ans++
			dfs(row, col)
		}
	}
	return ans
}

func solution2(grid [][]byte) int {
	// bfs
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	var bfs func(row, col int)
	bfs = func(row, col int) {
		vis[row][col] = true
		q := [][]int{{row, col}}
		for len(q) > 0 {
			row, col := q[0][0], q[0][1]
			q = q[1:]
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow][newCol] == '0' || vis[newRow][newCol] {
					continue
				}
				vis[newRow][newCol] = true
				q = append(q, []int{newRow, newCol})
			}
		}
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' || vis[row][col] {
				continue
			}
			ans++
			bfs(row, col)
		}
	}
	return ans
}

func solution3(grid [][]byte) int {
	uf := NewUnionFind(grid)
	if uf == nil {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' {
				continue
			}
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow][newCol] == '0' {
					continue
				}
				uf.Union(row*cols+col, newRow*cols+newCol)
			}
		}
	}
	return uf.Count()
}

type UnionFind struct {
	roots []int
	count int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	rows := len(grid)
	if rows == 0 {
		return nil
	}
	cols := len(grid[0])
	if cols == 0 {
		return nil
	}
	roots := make([]int, rows*cols)
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' {
				continue
			}
			index := row*cols + col
			roots[index] = index
			count++
		}
	}
	return &UnionFind{
		roots: roots,
		count: count,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.roots[x] != x {
		uf.roots[x] = uf.Find(uf.roots[x])
	}
	return uf.roots[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.roots[rootX] = rootY
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}
