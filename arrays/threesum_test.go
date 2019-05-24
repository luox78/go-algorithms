package arrays_test

import (
	"math"
	"testing"
)

func TestThreeSum(t *testing.T) {
	// 给定一个包含 n 个整数的数组 nums,判断 nums 中是否存在三个元素 a，b，c ,使得 a + b + c = 0 找出所有满足条件且不重复的三元组。
	// [-1, 0, 1, 2, -1, -4],
	// 满足要求的三元组集合为:
	// [
	//   [-1, 0, 1],
	//   [-1, -1, 2]
	// ]

	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	// [[-1 -1 2] [-1 0 1]]
	t.Log(res)

	nums = []int{0, 0, 0, 0}
	res = threeSum(nums)
	// [[0 0 0]]
	t.Log(res)

	nums = []int{-1, 0, 1, 2, -1, -4}
	res = threeSum(nums)
	// [[-1 -1 2],[-1 0 1]]
	t.Log(res)
}

// 找出符合条件且不重复的三元组
func threeSum(nums []int) [][]int {
	var res [][]int

	if len(nums) == 0 {
		return res
	}

	// 先进行从小到大排序
	qsort(nums, 0, len(nums)-1)

	// 全为正数或负数,无解
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return res
	}

	// 记录上一次处理过的数
	prev := math.MaxInt32

	for i := 0; i < len(nums)-2; i++ {
		// 判断是否处理过
		if prev == nums[i] {
			continue
		}
		prev = nums[i]

		first := i + 1
		last := len(nums) - 1

		for first < last {
			sum := nums[i] + nums[first] + nums[last]

			if sum == 0 {
				// 符合条件, 不能直接退出, 需要进行后面的组合
				res = append(res, []int{nums[i], nums[first], nums[last]})
			}
			if sum <= 0 {
				// 太小, 向右移动
				tmp := nums[first]
				first++
				// 跳过重复的数
				for first < last && tmp == nums[first] {
					first++
				}
			} else {
				// 太大, 向左移动
				tmp := nums[last]
				last--
				// 跳过重复的数
				for first < last && tmp == nums[last] {
					last--
				}
			}

		}

	}

	return res
}

// 快速排序
func qsort(nums []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(nums, p, r)

	qsort(nums, p, q-1)
	qsort(nums, q+1, r)
}

// 分区函数
func partition(nums []int, p, r int) int {
	// 分区值
	pivot := nums[r]

	i := p

	for j := p; j < r; j++ {
		if nums[j] <= pivot {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			i++
		}
	}

	nums[r] = nums[i]
	nums[i] = pivot
	return i
}
