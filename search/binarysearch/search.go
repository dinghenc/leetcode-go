package binarysearch

func Locate(arr []int, target int) int {
	left, right := 0, len(arr)-1
	// 查询范围 [left, right]
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
