//
// Package heap sort.go
// 堆排序实现
//
package heap

// Sort 堆排序
func Sort(nums []int) {
	if len(nums) < 2 {
		return
	}

	buildHeap(nums)

	for i := len(nums) - 1; i > 0; i-- {
		// 顶堆最大的元素和最后元素交换
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i-1, 0)
	}
}

// 原地建堆, 大顶堆
func buildHeap(nums []int) {
	n := len(nums) - 1
	for i := n / 2; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

// 堆化, 从上向下
func heapify(nums []int, n, i int) {
	for {
		maxPos := i
		// 和左子节点比较
		if i*2+1 <= n && nums[maxPos] < nums[i*2+1] {
			maxPos = i*2 + 1
		}
		// 和右子节点比较
		if i*2+2 <= n && nums[maxPos] < nums[i*2+2] {
			maxPos = i*2 + 2
		}
		if i == maxPos {
			break
		}

		// 交换数据
		nums[i], nums[maxPos] = nums[maxPos], nums[i]

		i = maxPos
	}
}
