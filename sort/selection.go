package sort

// SelectionSort 选择排序
func SelectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums)-1; i++ {
		min := i
		j := i + 1

		// 选择未排序中最小的进行数据交换
		for ; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		// 数据交换
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}
}
