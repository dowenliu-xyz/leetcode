package lc0200_2

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
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	var dfs func(grid [][]byte, row, col int)
	dfs = func(grid [][]byte, row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '0' || vis[row][col] {
			return
		}
		vis[row][col] = true
		dfs(grid, row, col+1)
		dfs(grid, row+1, col)
		dfs(grid, row, col-1)
		dfs(grid, row-1, col)
	}
	ans := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' || vis[row][col] {
				continue
			}
			ans++
			dfs(grid, row, col)
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
	vis := make([][]bool, rows)
	for i := range vis {
		vis[i] = make([]bool, cols)
	}
	var bfs func(grid [][]byte, row, col int)
	bfs = func(grid [][]byte, row, col int) {
		directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		q := [][]int{{row, col}}
		vis[row][col] = true
		for len(q) > 0 {
			pos := q[0]
			q = q[1:]
			for _, dir := range directions {
				newRow, newCol := pos[0]+dir[0], pos[1]+dir[1]
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
			bfs(grid, row, col)
		}
	}
	return ans
}

func solution3(grid [][]byte) int {
	// 并查集
	uf := NewUnionFind(grid)
	if uf == nil {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])
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
	uf := &UnionFind{
		roots: make([]int, rows*cols),
		count: 0,
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '0' {
				continue
			}
			index := row*cols + col
			uf.roots[index] = index
			uf.count++
		}
	}
	return uf
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
