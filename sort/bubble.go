package sort

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		// 提前退出冒泡的标记位
		var flag bool

		for j := 0; j < len(nums)-1-i; j++ {
			// 比较并交换数据
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp

				flag = true
			}
		}

		// 没有数据交换, 提前退出
		if !flag {
			break
		}
	}
}
