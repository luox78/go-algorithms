package unionfind_test

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	// 给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

	// 示例 1:
	// 输入:
	// 11110
	// 11010
	// 11000
	// 00000
	// 输出: 1
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	numIslands(grid)

	// 示例 2:
	// 输入:
	// 11000
	// 11000
	// 00100
	// 00011
	// 输出: 3
	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	numIslands(grid)
}

func numIslands(grid [][]byte) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}

	// 初始化并查集
	uf := NewUnionFind(grid)

	m, n := len(grid), len(grid[0])
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 跳过
			if grid[i][j] == '0' {
				continue
			}

			// 判断四周
			for k := 0; k < len(dx); k++ {
				nr, nc := i+dx[k], j+dy[k]

				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == '1' {
					// 合并
					uf.Union(i*n+j, nr*n+nc)
				}
			}
		}
	}

	return uf.Count()
}

// UnionFind 并查集结构
type UnionFind struct {
	parents []int
	rank    []int
	count   int
}

// NewUnionFind 初始化并查集
func NewUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])

	uf := new(UnionFind)
	uf.parents = make([]int, m*n)
	uf.rank = make([]int, m*n)

	for i := 0; i < m*n; i++ {
		uf.parents[i] = -1
		uf.rank[i] = 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.parents[i*n+j] = i*n + j
				uf.count++
			}
		}
	}
	return uf
}

// Find 查找
func (uf *UnionFind) Find(i int) int {
	if uf.parents[i] != i {
		uf.parents[i] = uf.Find(uf.parents[i])
	}
	return uf.parents[i]
}

// Union 合并
func (uf *UnionFind) Union(x, y int) {
	rootx := uf.Find(x)
	rooty := uf.Find(y)

	if rootx != rooty {
		if uf.rank[rootx] > uf.rank[rooty] {
			uf.parents[rooty] = rootx
		} else if uf.rank[rootx] < uf.rank[rooty] {
			uf.parents[rootx] = rooty
		} else {
			uf.parents[rooty] = rootx
			uf.rank[rootx]++
		}
		uf.count--
	}
}

// Count 返回集合数量
func (uf *UnionFind) Count() int {
	return uf.count
}
