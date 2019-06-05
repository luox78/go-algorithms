package arrays_test

import "testing"

func TestMerge1(t *testing.T) {
	// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

	// 说明:
	// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

	// 示例:
	// 输入:
	// nums1 = [1,2,3,0,0,0], m = 3
	// nums2 = [2,5,6],       n = 3

	// 输出: [1,2,2,3,5,6]

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge1(nums1, 3, nums2, 3)
	t.Log(nums1)
}

// 合并两个有序数组, 将 nums2 合并到 nums1 中
func merge1(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1

	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for ; p1 >= 0; p1-- {
		nums1[p] = nums1[p1]
		p--
	}
	for ; p2 >= 0; p2-- {
		nums1[p] = nums2[p2]
		p--
	}
}

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
