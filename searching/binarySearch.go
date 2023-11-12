package searching

import (
	"sort"
)

func BinarySearch(arr []int, low int, high int, x int) int {
	sort.Ints(arr)
	if low <= high {
		mid := (high + low) / 2
		if arr[mid] == x {
			return mid
		} else if x > arr[mid] {
			return BinarySearch(arr, mid+1, high, x)
		} else if x < arr[mid] {
			return BinarySearch(arr, low, mid-1, x)
		}
	}
	return -1
}
