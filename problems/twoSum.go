package problems

func TwoSum(arr []int, target int) (int, int) {
	p1 := 0
	p2 := len(arr) - 1

	for p1 < p2 {
		sum := arr[p1] + arr[p2]
		if sum == target {
			return p1, p2
		} else if sum < target {
			p1++
		} else {
			p2--
		}
	}
	return -1, -1
}
