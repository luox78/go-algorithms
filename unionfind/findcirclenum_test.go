package unionfind_test

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	// 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
	// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

	// 示例 1:
	// 输入:
	// [[1,1,0],
	//  [1,1,0],
	//  [0,0,1]]
	// 输出: 2
	// 说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
	// 第2个学生自己在一个朋友圈。所以返回2。
	M := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	t.Log(findCircleNum(M))

	// 示例 2:
	// 输入:
	// [[1,1,0],
	//  [1,1,1],
	//  [0,1,1]]
	// 输出: 1
	// 说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
	M = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 0, 1},
	}
	t.Log(findCircleNum(M))
}

func findCircleNum(M [][]int) int {
	if len(M) <= 0 || len(M[0]) <= 0 {
		return 0
	}

	uf := NewUF(M)
	m, n := len(M), len(M[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count()
}

// UF 并查集结构
type UF struct {
	roots []int
	count int
}

// NewUF 新建并查集
func NewUF(n [][]int) *UF {
	uf := new(UF)

	l := len(n)
	uf.roots = make([]int, l)
	for i := 0; i < l; i++ {
		uf.roots[i] = i
	}
	uf.count = l
	return uf
}

// Find 查找 root
func (uf *UF) Find(i int) int {
	root := i
	for uf.roots[root] != root {
		root = uf.roots[root]
	}

	// 路径压缩
	for i != root {
		uf.roots[i], i = root, uf.roots[i]
	}
	return root
}

// Union 合并
func (uf *UF) Union(p, q int) {
	rootp, rootq := uf.Find(p), uf.Find(q)
	if rootp != rootq {
		uf.roots[rootp] = uf.roots[rootq]
		uf.count--
	}
}

// Count 返回集合数量
func (uf *UF) Count() int { return uf.count }
