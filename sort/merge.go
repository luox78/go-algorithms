package sort

// MergeSort 归并排序
func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mergeSort(nums, 0, len(nums)-1)
}

// 递归分治
func mergeSort(nums []int, p, r int) {
	if p >= r {
		return
	}

	// 求中间索引
	q := p + (r-p)/2

	mergeSort(nums, p, q)
	mergeSort(nums, q+1, r)

	merge(nums, p, q, r)
}

// 合并函数
func merge(nums []int, p, q, r int) {
	// 临时存储
	tmp := make([]int, r-p+1)

	i, j, k := p, q+1, 0

	// 比较并加入临时存储
	for i <= q && j <= r {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	// 处理剩余的部分
	start, end := i, q
	if j <= r {
		start = j
		end = r
	}

	for ; start <= end; start++ {
		tmp[k] = nums[start]
		k++
	}

	// 数据搬移至原空间中
	for i := 0; i <= r-p; i++ {
		nums[p+i] = tmp[i]
	}
}
