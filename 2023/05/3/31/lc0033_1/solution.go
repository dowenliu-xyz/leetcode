package lc0033_1

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 { // 空数组的特殊情况
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target { // 刚好找到
			return mid
		}
		// 这里有个特别容易忽略的点，mid 可能等于 left，此时如果条件写成 nums[left] < nums[mid] 会导致这种情况错误判断
		if nums[left] <= nums[mid] { // nums[left:mid+1] 是升序的。
			if nums[left] <= target && target < nums[mid] { // target 在该升序区间，继续在该区间二分查找
				right = mid - 1
			} else { // 否则向右探测
				left = mid + 1
			}
		} else { // nums[mid+1:right+1] 是升序的
			if nums[mid] < target && target <= nums[right] { // target 在该升序区间，继续在该区间二分查找
				left = mid + 1
			} else { // 否则向左探测
				right = mid - 1
			}
		}
	}
	return -1
}
