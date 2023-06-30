package lc0200

func numIslands(grid [][]byte) int {
	//return solution1(grid)
	//return solution2(grid)
	return solution3(grid)
}

func solution1(grid [][]byte) int {
	// dfs + 染色法。会修改 grid
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		for _, dir := range directions {
			dfs(row+dir[0], col+dir[1])
		}
	}
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

func solution2(grid [][]byte) int {
	// dfs + 记忆法。当不允许改动 grid 时使用此方法
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	mem := make([][]bool, rows)
	for row := range mem {
		mem[row] = make([]bool, cols)
		for col := range mem[row] {
			if grid[row][col] == '1' {
				mem[row][col] = true
			}
		}
	}
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		if !mem[row][col] {
			return
		}
		mem[row][col] = false
		for _, dir := range directions {
			dfs(row+dir[0], col+dir[1])
		}
	}
	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' && mem[row][col] {
				count++
				dfs(row, col)
			}
		}
	}
	return count
}

func solution3(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	uf := NewUnionFind(grid, rows, cols)
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' {
				continue
			}
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
					continue
				}
				if grid[newRow][newCol] == '0' {
					continue
				}
				uf.Union(row*cols+col, newRow*cols+newCol)
			}
		}
	}
	return uf.Count()
}

type UnionFind struct {
	root  []int
	count int
}

func NewUnionFind(grid [][]byte, rows, cols int) *UnionFind {
	uf := &UnionFind{
		root:  make([]int, rows*cols),
		count: 0,
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				idx := row*cols + col
				uf.root[idx] = idx
				uf.count++
			}
		}
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.root[rootX] = rootY
	uf.count--
}

func (uf *UnionFind) Count() int {
	return uf.count
}
