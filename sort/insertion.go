package sort

// InsertionSort 插入排序
func InsertionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 1; i < len(nums); i++ {
		val := nums[i]

		// 查找插入的位置
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				// 数据移动, 空出位置
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		// 插入数据
		nums[j+1] = val
	}
}
