package sort

// QuickSort 快速排序
func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	quickSort(nums, 0, len(nums)-1)
}

// 递归函数
func quickSort(nums []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(nums, p, r)

	quickSort(nums, p, q-1)
	quickSort(nums, q+1, r)
}

// 分区函数
func partition(nums []int, p, r int) int {
	pivot := nums[r]

	j := p
	for i := p; i < r; i++ {
		if nums[i] < pivot {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}

	nums[r] = nums[j]
	nums[j] = pivot

	return j
}
