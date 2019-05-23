package arrays_test

import "testing"

func TestMerge(t *testing.T) {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := []int{2, 4}

	// [1 2 3 4 5 7 9]
	t.Log(merge(a1, a2))

	a1 = []int{}
	a2 = []int{3, 4, 5}

	// [3 4 5]
	t.Log(merge(a1, a2))

}

// 两个有序数组合并为一个有序数组
func merge(a1, a2 []int) []int {
	if len(a1) == 0 {
		return a2
	}
	if len(a2) == 0 {
		return a1
	}

	// 创建新数组
	l1, l2 := len(a1), len(a2)
	arr := make([]int, l1+l2)

	i, j, k := 0, 0, 0

	// 两个有序数组中元素进行比较并放入新数组
	for i < l1 && j < l2 {
		if a1[i] < a2[j] {
			arr[k] = a1[i]
			i++
		} else {
			arr[k] = a2[j]
			j++
		}
		k++
	}

	// 处理两个有序数组中剩余部门
	for ; i < l1; i++ {
		arr[k] = a1[i]
		k++
	}
	for ; j < l2; j++ {
		arr[k] = a2[j]
		k++
	}

	return arr
}
